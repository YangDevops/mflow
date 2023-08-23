package rpc_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mflow/apps/task"
	"github.com/infraboard/mflow/clients/rpc"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestQueryJobTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	set, err := client.JobTask().QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestWatchJobTaskLog(t *testing.T) {
	req := task.NewWatchJobTaskLogRequest("xx")
	stream, err := client.JobTask().WatchJobTaskLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(resp.Data))
	}
}

func init() {
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	c, err := rpc.NewClientSetFromEnv()
	if err != nil {
		panic(err)
	}
	client = c
}
