// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mflow/apps/task/pb/job_rpc.proto

package task

import (
	request "github.com/infraboard/mcube/http/request"
	resource "github.com/infraboard/mcube/pb/resource"
	job "github.com/infraboard/mflow/apps/job"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchJobTaskLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务Id
	// @gotags: json:"task_id"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	// 当一个任务有多个container时, 可以单独查询Container的日志
	// @gotags: json:"container_name"
	ContainerName string `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name"`
}

func (x *WatchJobTaskLogRequest) Reset() {
	*x = WatchJobTaskLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchJobTaskLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchJobTaskLogRequest) ProtoMessage() {}

func (x *WatchJobTaskLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchJobTaskLogRequest.ProtoReflect.Descriptor instead.
func (*WatchJobTaskLogRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *WatchJobTaskLogRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *WatchJobTaskLogRequest) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

type JobTaskStreamReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 日志数据, 二进制格式
	// @gotags: json:"data"
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
}

func (x *JobTaskStreamReponse) Reset() {
	*x = JobTaskStreamReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobTaskStreamReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskStreamReponse) ProtoMessage() {}

func (x *JobTaskStreamReponse) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskStreamReponse.ProtoReflect.Descriptor instead.
func (*JobTaskStreamReponse) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *JobTaskStreamReponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueryJobTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// 资源标签过滤
	// @gotags: json:"filters"
	Filters []*resource.LabelRequirement `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters"`
	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page"`
	// 任务Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,4,rep,name=ids,proto3" json:"ids"`
	// 与该pipeline task关联的任务
	// @gotags: json:"id"
	PipelineTaskId string `protobuf:"bytes,5,opt,name=pipeline_task_id,json=pipelineTaskId,proto3" json:"id"`
	// 任务当前状态
	// @gotags: json:"stage"
	Stage *STAGE `protobuf:"varint,6,opt,name=stage,proto3,enum=infraboard.mflow.task.STAGE,oneof" json:"stage"`
	// 任务标签
	// @gotags: json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryJobTaskRequest) Reset() {
	*x = QueryJobTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryJobTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJobTaskRequest) ProtoMessage() {}

func (x *QueryJobTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJobTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryJobTaskRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryJobTaskRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *QueryJobTaskRequest) GetFilters() []*resource.LabelRequirement {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *QueryJobTaskRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryJobTaskRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *QueryJobTaskRequest) GetPipelineTaskId() string {
	if x != nil {
		return x.PipelineTaskId
	}
	return ""
}

func (x *QueryJobTaskRequest) GetStage() STAGE {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return STAGE_PENDDING
}

func (x *QueryJobTaskRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type UpdateJobTaskStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 任务状态需要手动更新时的Token
	// @gotags: bson:"token" json:"token"
	UpdateToken string `protobuf:"bytes,2,opt,name=update_token,json=updateToken,proto3" json:"token" bson:"token"`
	// 强制更新任务状态, 默认已经完成的任务状态不能再修改
	// 用于任务重新运行
	// @gotags: bson:"force_update_status" json:"force_update_status"
	ForceUpdateStatus bool `protobuf:"varint,3,opt,name=force_update_status,json=forceUpdateStatus,proto3" json:"force_update_status" bson:"force_update_status"`
	// 强制触发流水线执行, 默认如果状态未变化不触发流水线执行
	// 用于任务重新运行
	// @gotags: bson:"force_trigger_pipeline" json:"force_trigger_pipeline"
	ForceTriggerPipeline bool `protobuf:"varint,7,opt,name=force_trigger_pipeline,json=forceTriggerPipeline,proto3" json:"force_trigger_pipeline" bson:"force_trigger_pipeline"`
	// 任务当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,4,opt,name=stage,proto3,enum=infraboard.mflow.task.STAGE" json:"stage" bson:"stage"`
	// 状态描述
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message" bson:"message"`
	// 任务状态详细描述, 用于Debug
	// @gotags: bson:"detail" json:"detail"
	Detail string `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail" bson:"detail"`
	// 状态扩展属性
	// @gotags: bson:"extension" json:"extension"
	Extension map[string]string `protobuf:"bytes,15,rep,name=extension,proto3" json:"extension" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extension"`
}

func (x *UpdateJobTaskStatusRequest) Reset() {
	*x = UpdateJobTaskStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobTaskStatusRequest) ProtoMessage() {}

func (x *UpdateJobTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateJobTaskStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobTaskStatusRequest) GetUpdateToken() string {
	if x != nil {
		return x.UpdateToken
	}
	return ""
}

func (x *UpdateJobTaskStatusRequest) GetForceUpdateStatus() bool {
	if x != nil {
		return x.ForceUpdateStatus
	}
	return false
}

func (x *UpdateJobTaskStatusRequest) GetForceTriggerPipeline() bool {
	if x != nil {
		return x.ForceTriggerPipeline
	}
	return false
}

func (x *UpdateJobTaskStatusRequest) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_PENDDING
}

func (x *UpdateJobTaskStatusRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateJobTaskStatusRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *UpdateJobTaskStatusRequest) GetExtension() map[string]string {
	if x != nil {
		return x.Extension
	}
	return nil
}

type UpdateJobTaskOutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 任务状态需要手动更新时的Token
	// @gotags: json:"token"
	UpdateToken string `protobuf:"bytes,2,opt,name=update_token,json=updateToken,proto3" json:"token"`
	// 强制更新, 默认已经完成的任务状态不能再修改
	// 用于任务重新运行
	// @gotags: bson:"force" json:"force"
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force" bson:"force"`
	// Job Task运行时环境变量, 大写开头的变量会更新到pipline中, 注入到后续执行的任务中
	// 小写开头的变量, 作为Task运行后的输出, task自己保留
	// @gotags: json:"runtime_envs"
	RuntimeEnvs []*job.RunParam `protobuf:"bytes,4,rep,name=runtime_envs,json=runtimeEnvs,proto3" json:"runtime_envs"`
	// 任务运行后的产物信息, 用于界面展示
	// @gotags: json:"markdown_output"
	MarkdownOutput string `protobuf:"bytes,5,opt,name=markdown_output,json=markdownOutput,proto3" json:"markdown_output"`
}

func (x *UpdateJobTaskOutputRequest) Reset() {
	*x = UpdateJobTaskOutputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobTaskOutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobTaskOutputRequest) ProtoMessage() {}

func (x *UpdateJobTaskOutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobTaskOutputRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobTaskOutputRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJobTaskOutputRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobTaskOutputRequest) GetUpdateToken() string {
	if x != nil {
		return x.UpdateToken
	}
	return ""
}

func (x *UpdateJobTaskOutputRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *UpdateJobTaskOutputRequest) GetRuntimeEnvs() []*job.RunParam {
	if x != nil {
		return x.RuntimeEnvs
	}
	return nil
}

func (x *UpdateJobTaskOutputRequest) GetMarkdownOutput() string {
	if x != nil {
		return x.MarkdownOutput
	}
	return ""
}

type DescribeJobTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeJobTaskRequest) Reset() {
	*x = DescribeJobTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeJobTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJobTaskRequest) ProtoMessage() {}

func (x *DescribeJobTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJobTaskRequest.ProtoReflect.Descriptor instead.
func (*DescribeJobTaskRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeJobTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteJobTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 强制删除, 当job有资源无法释放时, 比如 k8s里对应的job已经被删除了
	// @gotags: json:"force"
	Force bool `protobuf:"varint,2,opt,name=force,proto3" json:"force"`
}

func (x *DeleteJobTaskRequest) Reset() {
	*x = DeleteJobTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobTaskRequest) ProtoMessage() {}

func (x *DeleteJobTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mflow_apps_task_pb_job_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobTaskRequest) Descriptor() ([]byte, []int) {
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteJobTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteJobTaskRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

var File_mflow_apps_task_pb_job_rpc_proto protoreflect.FileDescriptor

var file_mflow_apps_task_pb_job_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x21, 0x6d, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f,
	0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x58, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x4a, 0x6f,
	0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd6, 0x03, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x48, 0x00,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22,
	0xb9, 0x03, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x34, 0x0a, 0x16, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x14, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53,
	0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x5e, 0x0a,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3c, 0x0a,
	0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd1, 0x01, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65,
	0x6e, 0x76, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x52, 0x75, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x6e, 0x76, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f,
	0x77, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x28, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x32, 0x8e, 0x04, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x52,
	0x50, 0x43, 0x12, 0x5d, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65,
	0x74, 0x12, 0x68, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x68, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f,
	0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x60, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x6f, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6d, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mflow_apps_task_pb_job_rpc_proto_rawDescOnce sync.Once
	file_mflow_apps_task_pb_job_rpc_proto_rawDescData = file_mflow_apps_task_pb_job_rpc_proto_rawDesc
)

func file_mflow_apps_task_pb_job_rpc_proto_rawDescGZIP() []byte {
	file_mflow_apps_task_pb_job_rpc_proto_rawDescOnce.Do(func() {
		file_mflow_apps_task_pb_job_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mflow_apps_task_pb_job_rpc_proto_rawDescData)
	})
	return file_mflow_apps_task_pb_job_rpc_proto_rawDescData
}

var file_mflow_apps_task_pb_job_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mflow_apps_task_pb_job_rpc_proto_goTypes = []interface{}{
	(*WatchJobTaskLogRequest)(nil),     // 0: infraboard.mflow.task.WatchJobTaskLogRequest
	(*JobTaskStreamReponse)(nil),       // 1: infraboard.mflow.task.JobTaskStreamReponse
	(*QueryJobTaskRequest)(nil),        // 2: infraboard.mflow.task.QueryJobTaskRequest
	(*UpdateJobTaskStatusRequest)(nil), // 3: infraboard.mflow.task.UpdateJobTaskStatusRequest
	(*UpdateJobTaskOutputRequest)(nil), // 4: infraboard.mflow.task.UpdateJobTaskOutputRequest
	(*DescribeJobTaskRequest)(nil),     // 5: infraboard.mflow.task.DescribeJobTaskRequest
	(*DeleteJobTaskRequest)(nil),       // 6: infraboard.mflow.task.DeleteJobTaskRequest
	nil,                                // 7: infraboard.mflow.task.QueryJobTaskRequest.LabelsEntry
	nil,                                // 8: infraboard.mflow.task.UpdateJobTaskStatusRequest.ExtensionEntry
	(*resource.Scope)(nil),             // 9: infraboard.mcube.resource.Scope
	(*resource.LabelRequirement)(nil),  // 10: infraboard.mcube.resource.LabelRequirement
	(*request.PageRequest)(nil),        // 11: infraboard.mcube.page.PageRequest
	(STAGE)(0),                         // 12: infraboard.mflow.task.STAGE
	(*job.RunParam)(nil),               // 13: infraboard.mflow.job.RunParam
	(*JobTaskSet)(nil),                 // 14: infraboard.mflow.task.JobTaskSet
	(*JobTask)(nil),                    // 15: infraboard.mflow.task.JobTask
}
var file_mflow_apps_task_pb_job_rpc_proto_depIdxs = []int32{
	9,  // 0: infraboard.mflow.task.QueryJobTaskRequest.scope:type_name -> infraboard.mcube.resource.Scope
	10, // 1: infraboard.mflow.task.QueryJobTaskRequest.filters:type_name -> infraboard.mcube.resource.LabelRequirement
	11, // 2: infraboard.mflow.task.QueryJobTaskRequest.page:type_name -> infraboard.mcube.page.PageRequest
	12, // 3: infraboard.mflow.task.QueryJobTaskRequest.stage:type_name -> infraboard.mflow.task.STAGE
	7,  // 4: infraboard.mflow.task.QueryJobTaskRequest.labels:type_name -> infraboard.mflow.task.QueryJobTaskRequest.LabelsEntry
	12, // 5: infraboard.mflow.task.UpdateJobTaskStatusRequest.stage:type_name -> infraboard.mflow.task.STAGE
	8,  // 6: infraboard.mflow.task.UpdateJobTaskStatusRequest.extension:type_name -> infraboard.mflow.task.UpdateJobTaskStatusRequest.ExtensionEntry
	13, // 7: infraboard.mflow.task.UpdateJobTaskOutputRequest.runtime_envs:type_name -> infraboard.mflow.job.RunParam
	2,  // 8: infraboard.mflow.task.JobRPC.QueryJobTask:input_type -> infraboard.mflow.task.QueryJobTaskRequest
	3,  // 9: infraboard.mflow.task.JobRPC.UpdateJobTaskStatus:input_type -> infraboard.mflow.task.UpdateJobTaskStatusRequest
	4,  // 10: infraboard.mflow.task.JobRPC.UpdateJobTaskOutput:input_type -> infraboard.mflow.task.UpdateJobTaskOutputRequest
	5,  // 11: infraboard.mflow.task.JobRPC.DescribeJobTask:input_type -> infraboard.mflow.task.DescribeJobTaskRequest
	0,  // 12: infraboard.mflow.task.JobRPC.WatchJobTaskLog:input_type -> infraboard.mflow.task.WatchJobTaskLogRequest
	14, // 13: infraboard.mflow.task.JobRPC.QueryJobTask:output_type -> infraboard.mflow.task.JobTaskSet
	15, // 14: infraboard.mflow.task.JobRPC.UpdateJobTaskStatus:output_type -> infraboard.mflow.task.JobTask
	15, // 15: infraboard.mflow.task.JobRPC.UpdateJobTaskOutput:output_type -> infraboard.mflow.task.JobTask
	15, // 16: infraboard.mflow.task.JobRPC.DescribeJobTask:output_type -> infraboard.mflow.task.JobTask
	1,  // 17: infraboard.mflow.task.JobRPC.WatchJobTaskLog:output_type -> infraboard.mflow.task.JobTaskStreamReponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mflow_apps_task_pb_job_rpc_proto_init() }
func file_mflow_apps_task_pb_job_rpc_proto_init() {
	if File_mflow_apps_task_pb_job_rpc_proto != nil {
		return
	}
	file_mflow_apps_task_pb_job_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchJobTaskLogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobTaskStreamReponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryJobTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobTaskStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobTaskOutputRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeJobTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mflow_apps_task_pb_job_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_mflow_apps_task_pb_job_rpc_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mflow_apps_task_pb_job_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mflow_apps_task_pb_job_rpc_proto_goTypes,
		DependencyIndexes: file_mflow_apps_task_pb_job_rpc_proto_depIdxs,
		MessageInfos:      file_mflow_apps_task_pb_job_rpc_proto_msgTypes,
	}.Build()
	File_mflow_apps_task_pb_job_rpc_proto = out.File
	file_mflow_apps_task_pb_job_rpc_proto_rawDesc = nil
	file_mflow_apps_task_pb_job_rpc_proto_goTypes = nil
	file_mflow_apps_task_pb_job_rpc_proto_depIdxs = nil
}