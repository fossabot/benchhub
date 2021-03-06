package nodeutil

import (
	"github.com/dyweb/gommon/errors"
	"os"

	"github.com/benchhub/benchhub/lib/monitor/host"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/common/config"
)

// return node info that is needed when register agent and heartbeat

// GetNode returns node id, capacity, start & boot time
// TODO: addr https://github.com/benchhub/benchhub/issues/18
func GetNodeInfo(cfg config.NodeConfig) (*pb.NodeInfo, error) {
	m := host.NewMachine()
	if err := m.Update(); err != nil {
		return nil, errors.Wrap(err, "can't get node info")
	}
	node := &pb.NodeInfo{
		Id:        UID(),
		Role:      cfg.Role,
		Host:      hostname(),
		BootTime:  int64(m.BootTime), // unix ts in second
		StartTime: startTime.Unix(),  // unix ts in second
		Capacity: pb.NodeCapacity{
			Cores:       int32(m.NumCores),
			MemoryFree:  int32(m.MemFree / 1024),              // KB -> MB
			MemoryTotal: int32(m.MemTotal / 1024),             // KB -> MB
			DiskFree:    int32(m.DiskSpaceFree / 1024 / 1024), // B -> KB -> MB
			DiskTotal:   int32(m.DiskSpaceTotal / 1024 / 1024),
		},
		Provider: cfg.Provider,
	}
	return node, nil
}

func hostname() string {
	if name, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return name
	}
}
