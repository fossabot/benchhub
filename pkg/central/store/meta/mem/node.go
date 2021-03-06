package mem

import (
	"github.com/dyweb/gommon/errors"

	"fmt"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

var emptyNode = pb.Node{}

// -- start of read --

func (s *MetaStore) NumNodes() (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.nodes), nil
}

func (s *MetaStore) FindNodeById(id string) (pb.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if n, ok := s.nodes[id]; ok {
		return n, nil
	} else {
		return emptyNode, &pb.Error{
			Code:    pb.ErrorCode_NOT_FOUND,
			Message: fmt.Sprintf("node %s not found", id),
		}
	}
}

func (s *MetaStore) ListNodes() ([]pb.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	nodes := make([]pb.Node, 0, len(s.nodes))
	for id := range s.nodes {
		nodes = append(nodes, s.nodes[id])
	}
	return nodes, nil
}

func (s *MetaStore) ListNodesStatus() ([]pb.NodeStatus, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	status := make([]pb.NodeStatus, 0, len(s.status))
	for id := range s.status {
		status = append(status, s.status[id])
	}
	return status, nil
}

// -- end of read--

// -- start of write --

func (s *MetaStore) AddNode(id string, node pb.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; ok {
		return &pb.Error{
			Code:    pb.ErrorCode_ALREADY_EXISTS,
			Message: fmt.Sprintf("node %s already exists", id),
		}
	}
	s.nodes[id] = node
	return nil
}

func (s *MetaStore) UpdateNode(id string, node pb.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; !ok {
		return &pb.Error{
			Code:    pb.ErrorCode_NOT_FOUND,
			Message: fmt.Sprintf("node %s not found", id),
		}
	}
	s.nodes[id] = node
	return nil
}

func (s *MetaStore) UpdateNodeStatus(id string, status pb.NodeStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status[id] = status
	return nil
}

// -- end of write --

// -- start of delete --

func (s *MetaStore) RemoveNode(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; !ok {
		return errors.Errorf("node %s does not exists", id)
	}
	delete(s.nodes, id)
	return nil
}

// -- end of delete --
