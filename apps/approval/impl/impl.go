package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mflow/apps/approval"
	"github.com/infraboard/mflow/apps/pipeline"
	"github.com/infraboard/mflow/apps/task"
	"github.com/infraboard/mflow/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	approval.UnimplementedRPCServer
	ioc.IocObjectImpl

	pipeline pipeline.Service
	task     task.PipelineService
	mcenter  *rpc.ClientSet
}

func (s *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{
				{Key: "domain", Value: bsonx.Int32(-1)},
				{Key: "namespace", Value: bsonx.Int32(-1)},
				{Key: "version", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = s.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.pipeline = ioc.GetController(pipeline.AppName).(pipeline.Service)
	s.task = ioc.GetController(task.AppName).(task.Service)
	s.mcenter = rpc.C()
	return nil
}

func (s *impl) Name() string {
	return approval.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	approval.RegisterRPCServer(server, s)
}
