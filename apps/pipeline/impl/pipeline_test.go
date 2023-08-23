package impl_test

import (
	"testing"

	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mflow/apps/job"
	"github.com/infraboard/mflow/apps/pipeline"
	"github.com/infraboard/mflow/test/conf"
	"github.com/infraboard/mflow/test/tools"
)

func TestCreatemflowPipeline(t *testing.T) {
	req := pipeline.NewCreatePipelineRequest()
	tools.MustReadYamlFile("test/image_build_deploy.yml", req)
	ins, err := impl.CreatePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestQueryPipeline(t *testing.T) {
	req := pipeline.NewQueryPipelineRequest()
	set, err := impl.QueryPipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestDescribePipeline(t *testing.T) {
	req := pipeline.NewDescribePipelineRequest(conf.C.CICD_PIPELINE_ID)
	ins, err := impl.DescribePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateTestPipeline(t *testing.T) {
	req := pipeline.NewPutPipelineRequest(conf.C.CICD_PIPELINE_ID)
	tools.MustReadYamlFile("test/image_build_deploy.yml", req.Spec)
	req.Spec.VisiableMode = resource.VISIABLE_GLOBAL
	ins, err := impl.UpdatePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeletePipeline(t *testing.T) {
	req := pipeline.NewDeletePipelineRequest(conf.C.CICD_PIPELINE_ID)
	ins, err := impl.DeletePipeline(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestNewCreatePipelineRequestFromYAML(t *testing.T) {
	yml := tools.MustReadContentFile("test/test.yml")

	obj, err := pipeline.NewCreatePipelineRequestFromYAML(yml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(obj)
}

func TestToYaml(t *testing.T) {
	in := pipeline.NewCreatePipelineRequest()
	in.Name = "pipeline_example"
	in.Description = "example"
	in.AddStage(
		&pipeline.Stage{
			Name: "stage_01",
			With: []*job.RunParam{
				{Name: "param1", Value: "value1"},
			},
			Jobs: []*pipeline.RunJobRequest{
				{JobName: "job01", RunParams: &job.RunParamSet{
					Params: []*job.RunParam{
						{Name: "param1", Value: "value1"},
					},
				}},
			},
		},
		&pipeline.Stage{
			Name: "stage_02",
			With: []*job.RunParam{
				{Name: "param1", Value: "value1"},
			},
			Jobs: []*pipeline.RunJobRequest{
				{JobName: "job01", RunParams: &job.RunParamSet{
					Params: []*job.RunParam{
						{Name: "param1", Value: "value1"},
					},
				}},
			},
		},
	)
	t.Log(in.ToYAML())
}
