package server

import (
	"bytes"
	"testing"

	"github.com/dyweb/gommon/config"
	"github.com/dyweb/gommon/util/testutil"

	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"github.com/benchhub/benchhub/pkg/common/spec"
)

func TestJobController_AcquireNodes(t *testing.T) {
	testutil.SkipIf(t, testutil.IsTravis())

	j, err := NewJobController(nil)
	if err != nil {
		t.Fatal(err)
	}

	twoNodes := []pbc.Node{
		{
			Uid:  "a",
			Role: pbc.Role_ANY,
		},
		{
			Uid:  "b",
			Role: pbc.Role_ANY,
		},
	}
	t.Run("two agent two nodes", func(t *testing.T) {
		// FIXME: hardcoded value
		job := readJob(t, "/home/at15/workspace/src/github.com/benchhub/benchhub/pkg/common/spec/pingpong.yml")
		res, err := j.AcquireNodes(twoNodes, job.Nodes)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	})
}

func readJob(t *testing.T, path string) spec.Job {
	b := testutil.ReadFixture(t, path)
	var job spec.Job
	if err := config.LoadYAMLDirectFrom(bytes.NewReader(b), &job); err != nil {
		t.Fatal(err)
	}
	if err := job.Validate(); err != nil {
		t.Fatal(err)
	}
	return job
}