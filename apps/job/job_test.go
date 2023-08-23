package job_test

import (
	"testing"

	"github.com/infraboard/mflow/apps/job"
	"github.com/infraboard/mflow/test/tools"
)

func TestK8SJobRunnerParams(t *testing.T) {
	param := job.NewRunParamSet()
	param.Add(&job.RunParam{
		Name:     "cluster_id",
		Required: true,
		Value:    "k8s-test",
	}, &job.RunParam{
		Name:     "namespace",
		Required: true,
		Value:    "default",
	})
	t.Log(param.K8SJobRunnerParams())
}

func TestCheckDuplicate(t *testing.T) {
	param := job.NewRunParamSet()
	param.Add(&job.RunParam{
		Name:     "cluster_id",
		Required: true,
		Value:    "k8s-test",
	}, &job.RunParam{
		Name:     "cluster_id",
		Required: true,
		Value:    "default",
	})
	t.Log(param.CheckDuplicate())
}

func TestNewRunParamWithKVPaire(t *testing.T) {
	param := job.NewRunParamWithKVPaire("key1", "value1", "key2", "value2")
	t.Log(param)
}

func TestParseRuntimeEnvFromBytes(t *testing.T) {
	data := tools.MustReadContentFile("test/pipeline.env")
	envs, err := job.ParseRunParamFromBytes([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(envs)
}
