// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: timer.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// EXECUTOR CONFIGS
type Method int32

const (
	Method_GET  Method = 0
	Method_POST Method = 1
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "GET",
		1: "POST",
	}
	Method_value = map[string]int32{
		"GET":  0,
		"POST": 1,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_timer_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_timer_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{0}
}

// TIMER
type Timer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Account        string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	ExecutionCount uint32 `protobuf:"varint,3,opt,name=execution_count,json=executionCount,proto3" json:"execution_count,omitempty"`
	Meta           *Meta  `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	// Types that are assignable to ExecutorConfig:
	//	*Timer_HttpConfig
	ExecutorConfig isTimer_ExecutorConfig `protobuf_oneof:"executor_config"`
	// Types that are assignable to SchedulerConfig:
	//	*Timer_CronConfig
	//	*Timer_IntervalConfig
	SchedulerConfig isTimer_SchedulerConfig `protobuf_oneof:"scheduler_config"`
}

func (x *Timer) Reset() {
	*x = Timer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timer) ProtoMessage() {}

func (x *Timer) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timer.ProtoReflect.Descriptor instead.
func (*Timer) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{0}
}

func (x *Timer) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Timer) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Timer) GetExecutionCount() uint32 {
	if x != nil {
		return x.ExecutionCount
	}
	return 0
}

func (x *Timer) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (m *Timer) GetExecutorConfig() isTimer_ExecutorConfig {
	if m != nil {
		return m.ExecutorConfig
	}
	return nil
}

func (x *Timer) GetHttpConfig() *HTTPConfig {
	if x, ok := x.GetExecutorConfig().(*Timer_HttpConfig); ok {
		return x.HttpConfig
	}
	return nil
}

func (m *Timer) GetSchedulerConfig() isTimer_SchedulerConfig {
	if m != nil {
		return m.SchedulerConfig
	}
	return nil
}

func (x *Timer) GetCronConfig() *CronConfig {
	if x, ok := x.GetSchedulerConfig().(*Timer_CronConfig); ok {
		return x.CronConfig
	}
	return nil
}

func (x *Timer) GetIntervalConfig() *IntervalConfig {
	if x, ok := x.GetSchedulerConfig().(*Timer_IntervalConfig); ok {
		return x.IntervalConfig
	}
	return nil
}

type isTimer_ExecutorConfig interface {
	isTimer_ExecutorConfig()
}

type Timer_HttpConfig struct {
	HttpConfig *HTTPConfig `protobuf:"bytes,5,opt,name=http_config,json=httpConfig,proto3,oneof"`
}

func (*Timer_HttpConfig) isTimer_ExecutorConfig() {}

type isTimer_SchedulerConfig interface {
	isTimer_SchedulerConfig()
}

type Timer_CronConfig struct {
	CronConfig *CronConfig `protobuf:"bytes,6,opt,name=cron_config,json=cronConfig,proto3,oneof"`
}

type Timer_IntervalConfig struct {
	IntervalConfig *IntervalConfig `protobuf:"bytes,7,opt,name=interval_config,json=intervalConfig,proto3,oneof"`
}

func (*Timer_CronConfig) isTimer_SchedulerConfig() {}

func (*Timer_IntervalConfig) isTimer_SchedulerConfig() {}

// META
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type HTTPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method  Method   `protobuf:"varint,2,opt,name=method,proto3,enum=executor.Method" json:"method,omitempty"`
	Body    string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Headers []string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *HTTPConfig) Reset() {
	*x = HTTPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPConfig) ProtoMessage() {}

func (x *HTTPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPConfig.ProtoReflect.Descriptor instead.
func (*HTTPConfig) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HTTPConfig) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_GET
}

func (x *HTTPConfig) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *HTTPConfig) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

// SCHEDULER CONFIGS
type IntervalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Interval   int32                `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Executions int32                `protobuf:"zigzag32,3,opt,name=executions,proto3" json:"executions,omitempty"`
}

func (x *IntervalConfig) Reset() {
	*x = IntervalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntervalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntervalConfig) ProtoMessage() {}

func (x *IntervalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntervalConfig.ProtoReflect.Descriptor instead.
func (*IntervalConfig) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{3}
}

func (x *IntervalConfig) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *IntervalConfig) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *IntervalConfig) GetExecutions() int32 {
	if x != nil {
		return x.Executions
	}
	return 0
}

type CronConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Cron       string               `protobuf:"bytes,2,opt,name=cron,proto3" json:"cron,omitempty"`
	Executions int32                `protobuf:"zigzag32,3,opt,name=executions,proto3" json:"executions,omitempty"`
}

func (x *CronConfig) Reset() {
	*x = CronConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronConfig) ProtoMessage() {}

func (x *CronConfig) ProtoReflect() protoreflect.Message {
	mi := &file_timer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronConfig.ProtoReflect.Descriptor instead.
func (*CronConfig) Descriptor() ([]byte, []int) {
	return file_timer_proto_rawDescGZIP(), []int{4}
}

func (x *CronConfig) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CronConfig) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *CronConfig) GetExecutions() int32 {
	if x != nil {
		return x.Executions
	}
	return 0
}

var File_timer_proto protoreflect.FileDescriptor

var file_timer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x05, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0b, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x01, 0x52,
	0x0a, 0x63, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x01,
	0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x11, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x12, 0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x43, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x0a,
	0x48, 0x54, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7b,
	0x0a, 0x0a, 0x43, 0x72, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x1b, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timer_proto_rawDescOnce sync.Once
	file_timer_proto_rawDescData = file_timer_proto_rawDesc
)

func file_timer_proto_rawDescGZIP() []byte {
	file_timer_proto_rawDescOnce.Do(func() {
		file_timer_proto_rawDescData = protoimpl.X.CompressGZIP(file_timer_proto_rawDescData)
	})
	return file_timer_proto_rawDescData
}

var file_timer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_timer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_timer_proto_goTypes = []interface{}{
	(Method)(0),                 // 0: executor.Method
	(*Timer)(nil),               // 1: executor.Timer
	(*Meta)(nil),                // 2: executor.Meta
	(*HTTPConfig)(nil),          // 3: executor.HTTPConfig
	(*IntervalConfig)(nil),      // 4: executor.IntervalConfig
	(*CronConfig)(nil),          // 5: executor.CronConfig
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_timer_proto_depIdxs = []int32{
	2, // 0: executor.Timer.meta:type_name -> executor.Meta
	3, // 1: executor.Timer.http_config:type_name -> executor.HTTPConfig
	5, // 2: executor.Timer.cron_config:type_name -> executor.CronConfig
	4, // 3: executor.Timer.interval_config:type_name -> executor.IntervalConfig
	6, // 4: executor.Meta.create_time:type_name -> google.protobuf.Timestamp
	0, // 5: executor.HTTPConfig.method:type_name -> executor.Method
	6, // 6: executor.IntervalConfig.start_time:type_name -> google.protobuf.Timestamp
	6, // 7: executor.CronConfig.start_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_timer_proto_init() }
func file_timer_proto_init() {
	if File_timer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timer); i {
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
		file_timer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_timer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPConfig); i {
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
		file_timer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntervalConfig); i {
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
		file_timer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronConfig); i {
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
	file_timer_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Timer_HttpConfig)(nil),
		(*Timer_CronConfig)(nil),
		(*Timer_IntervalConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_timer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timer_proto_goTypes,
		DependencyIndexes: file_timer_proto_depIdxs,
		EnumInfos:         file_timer_proto_enumTypes,
		MessageInfos:      file_timer_proto_msgTypes,
	}.Build()
	File_timer_proto = out.File
	file_timer_proto_rawDesc = nil
	file_timer_proto_goTypes = nil
	file_timer_proto_depIdxs = nil
}
