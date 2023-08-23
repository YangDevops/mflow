package impl_test

import (
	"testing"

	"github.com/infraboard/mflow/apps/trigger"
	"github.com/infraboard/mflow/test/tools"
)

func TestHandleEvent(t *testing.T) {
	raw := tools.MustReadContentFile("test/gitlab_push.json")
	req := trigger.NewGitlabEvent(raw)
	req.SkipRunPipeline = false

	ps, err := impl.HandleEvent(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ps))
}

func TestQueryRecord(t *testing.T) {
	req := trigger.NewQueryRecordRequest()
	set, err := impl.QueryRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
