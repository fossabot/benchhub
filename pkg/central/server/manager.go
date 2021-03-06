package server

import (
	"context"
	"sync"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"

	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

type Manager struct {
	cfg config.ServerConfig

	registry *Registry

	meta          meta.Provider
	job           *JobPoller
	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server

	log *dlog.Logger
}

func NewManager(cfg config.ServerConfig) (*Manager, error) {
	log.Infof("creating benchhub central manager")
	metaStore, err := meta.GetProvider(cfg.Meta.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create meta store")
	}

	// registry
	r := NewRegistry(cfg)
	r.Meta = metaStore

	// job poller
	job, err := NewJobPoller(r, cfg.Job.PollInterval)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create job controller")
	}

	// grpc http
	grpcSrv, err := NewGrpcServer(metaStore, r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterBenchHubCentralServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc transport")
	}
	httpSrv, err := NewHttpServer(metaStore, r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http server")
	}
	httpTransport, err := ihttp.NewServer(cfg.Http, httpSrv.Handler(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http transport")
	}
	mgr := &Manager{
		cfg:           cfg,
		registry:      r,
		meta:          metaStore,
		job:           job,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
		httpSrv:       httpSrv,
		httpTransport: httpTransport,
	}
	dlog.NewStructLogger(log, mgr)
	return mgr, nil
}

func (mgr *Manager) Run() error {
	var (
		wg      sync.WaitGroup
		grpcErr error
		httpErr error
		merr    = errors.NewMultiErrSafe()
	)
	wg.Add(3) // grpc + http + job controller
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO: logic here are pretty duplicated
	// grpc server
	go func() {
		go func() {
			if err := mgr.grpcTransport.Run(); err != nil {
				grpcErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if grpcErr != nil {
				merr.Append(grpcErr)
				mgr.log.Errorf("can't run grpc server %v", grpcErr)
			} else {
				mgr.log.Warn("TODO: other's fault, need to shutdown grpc server")
			}
			wg.Done()
			return
		}
	}()
	// http server
	go func() {
		go func() {
			if err := mgr.httpTransport.Run(); err != nil {
				httpErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if httpErr != nil {
				merr.Append(httpErr)
				mgr.log.Errorf("can't run http server %v", httpErr)
			} else {
				// other service's fault
				mgr.log.Warn("TODO: other's fault, need to shutdown http server")
			}
			wg.Done()
			return
		}
	}()
	// job poller
	go func() {
		if err := mgr.job.RunWithContext(ctx); err != nil {
			merr.Append(err)
			mgr.log.Warnf("can't run job controller %v", err)
			cancel()
		}
		wg.Done()
	}()
	wg.Wait()
	return merr.ErrorOrNil()
}
