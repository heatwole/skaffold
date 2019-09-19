// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skaffold.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StateResponse struct {
	State                *State   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{0}
}

func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type State struct {
	BuildState           *BuildState          `protobuf:"bytes,1,opt,name=buildState,proto3" json:"buildState,omitempty"`
	DeployState          *DeployState         `protobuf:"bytes,2,opt,name=deployState,proto3" json:"deployState,omitempty"`
	ForwardedPorts       map[int32]*PortEvent `protobuf:"bytes,4,rep,name=forwardedPorts,proto3" json:"forwardedPorts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{3}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetBuildState() *BuildState {
	if m != nil {
		return m.BuildState
	}
	return nil
}

func (m *State) GetDeployState() *DeployState {
	if m != nil {
		return m.DeployState
	}
	return nil
}

func (m *State) GetForwardedPorts() map[int32]*PortEvent {
	if m != nil {
		return m.ForwardedPorts
	}
	return nil
}

// BuildState contains a map of all skaffold artifacts to their current build
// states
type BuildState struct {
	Artifacts            map[string]string `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BuildState) Reset()         { *m = BuildState{} }
func (m *BuildState) String() string { return proto.CompactTextString(m) }
func (*BuildState) ProtoMessage()    {}
func (*BuildState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{4}
}

func (m *BuildState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildState.Unmarshal(m, b)
}
func (m *BuildState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildState.Marshal(b, m, deterministic)
}
func (m *BuildState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildState.Merge(m, src)
}
func (m *BuildState) XXX_Size() int {
	return xxx_messageInfo_BuildState.Size(m)
}
func (m *BuildState) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildState.DiscardUnknown(m)
}

var xxx_messageInfo_BuildState proto.InternalMessageInfo

func (m *BuildState) GetArtifacts() map[string]string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

// DeployState contains the status of the current deploy
type DeployState struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployState) Reset()         { *m = DeployState{} }
func (m *DeployState) String() string { return proto.CompactTextString(m) }
func (*DeployState) ProtoMessage()    {}
func (*DeployState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{5}
}

func (m *DeployState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployState.Unmarshal(m, b)
}
func (m *DeployState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployState.Marshal(b, m, deterministic)
}
func (m *DeployState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployState.Merge(m, src)
}
func (m *DeployState) XXX_Size() int {
	return xxx_messageInfo_DeployState.Size(m)
}
func (m *DeployState) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployState.DiscardUnknown(m)
}

var xxx_messageInfo_DeployState proto.InternalMessageInfo

func (m *DeployState) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Event struct {
	// Types that are valid to be assigned to EventType:
	//	*Event_MetaEvent
	//	*Event_BuildEvent
	//	*Event_DeployEvent
	//	*Event_PortEvent
	EventType            isEvent_EventType `protobuf_oneof:"event_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{6}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_EventType interface {
	isEvent_EventType()
}

type Event_MetaEvent struct {
	MetaEvent *MetaEvent `protobuf:"bytes,1,opt,name=metaEvent,proto3,oneof"`
}

type Event_BuildEvent struct {
	BuildEvent *BuildEvent `protobuf:"bytes,2,opt,name=buildEvent,proto3,oneof"`
}

type Event_DeployEvent struct {
	DeployEvent *DeployEvent `protobuf:"bytes,3,opt,name=deployEvent,proto3,oneof"`
}

type Event_PortEvent struct {
	PortEvent *PortEvent `protobuf:"bytes,4,opt,name=portEvent,proto3,oneof"`
}

func (*Event_MetaEvent) isEvent_EventType() {}

func (*Event_BuildEvent) isEvent_EventType() {}

func (*Event_DeployEvent) isEvent_EventType() {}

func (*Event_PortEvent) isEvent_EventType() {}

func (m *Event) GetEventType() isEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *Event) GetMetaEvent() *MetaEvent {
	if x, ok := m.GetEventType().(*Event_MetaEvent); ok {
		return x.MetaEvent
	}
	return nil
}

func (m *Event) GetBuildEvent() *BuildEvent {
	if x, ok := m.GetEventType().(*Event_BuildEvent); ok {
		return x.BuildEvent
	}
	return nil
}

func (m *Event) GetDeployEvent() *DeployEvent {
	if x, ok := m.GetEventType().(*Event_DeployEvent); ok {
		return x.DeployEvent
	}
	return nil
}

func (m *Event) GetPortEvent() *PortEvent {
	if x, ok := m.GetEventType().(*Event_PortEvent); ok {
		return x.PortEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_MetaEvent)(nil),
		(*Event_BuildEvent)(nil),
		(*Event_DeployEvent)(nil),
		(*Event_PortEvent)(nil),
	}
}

type MetaEvent struct {
	Entry                string   `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaEvent) Reset()         { *m = MetaEvent{} }
func (m *MetaEvent) String() string { return proto.CompactTextString(m) }
func (*MetaEvent) ProtoMessage()    {}
func (*MetaEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{7}
}

func (m *MetaEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaEvent.Unmarshal(m, b)
}
func (m *MetaEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaEvent.Marshal(b, m, deterministic)
}
func (m *MetaEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaEvent.Merge(m, src)
}
func (m *MetaEvent) XXX_Size() int {
	return xxx_messageInfo_MetaEvent.Size(m)
}
func (m *MetaEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MetaEvent proto.InternalMessageInfo

func (m *MetaEvent) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

type BuildEvent struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Err                  string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildEvent) Reset()         { *m = BuildEvent{} }
func (m *BuildEvent) String() string { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()    {}
func (*BuildEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{8}
}

func (m *BuildEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent.Unmarshal(m, b)
}
func (m *BuildEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent.Marshal(b, m, deterministic)
}
func (m *BuildEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent.Merge(m, src)
}
func (m *BuildEvent) XXX_Size() int {
	return xxx_messageInfo_BuildEvent.Size(m)
}
func (m *BuildEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent proto.InternalMessageInfo

func (m *BuildEvent) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *BuildEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *BuildEvent) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DeployEvent struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployEvent) Reset()         { *m = DeployEvent{} }
func (m *DeployEvent) String() string { return proto.CompactTextString(m) }
func (*DeployEvent) ProtoMessage()    {}
func (*DeployEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{9}
}

func (m *DeployEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployEvent.Unmarshal(m, b)
}
func (m *DeployEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployEvent.Marshal(b, m, deterministic)
}
func (m *DeployEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployEvent.Merge(m, src)
}
func (m *DeployEvent) XXX_Size() int {
	return xxx_messageInfo_DeployEvent.Size(m)
}
func (m *DeployEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeployEvent proto.InternalMessageInfo

func (m *DeployEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeployEvent) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type PortEvent struct {
	LocalPort            int32    `protobuf:"varint,1,opt,name=localPort,proto3" json:"localPort,omitempty"`
	RemotePort           int32    `protobuf:"varint,2,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
	PodName              string   `protobuf:"bytes,3,opt,name=podName,proto3" json:"podName,omitempty"`
	ContainerName        string   `protobuf:"bytes,4,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Namespace            string   `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PortName             string   `protobuf:"bytes,6,opt,name=portName,proto3" json:"portName,omitempty"`
	ResourceType         string   `protobuf:"bytes,7,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceName         string   `protobuf:"bytes,8,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortEvent) Reset()         { *m = PortEvent{} }
func (m *PortEvent) String() string { return proto.CompactTextString(m) }
func (*PortEvent) ProtoMessage()    {}
func (*PortEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{10}
}

func (m *PortEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortEvent.Unmarshal(m, b)
}
func (m *PortEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortEvent.Marshal(b, m, deterministic)
}
func (m *PortEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortEvent.Merge(m, src)
}
func (m *PortEvent) XXX_Size() int {
	return xxx_messageInfo_PortEvent.Size(m)
}
func (m *PortEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PortEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PortEvent proto.InternalMessageInfo

func (m *PortEvent) GetLocalPort() int32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *PortEvent) GetRemotePort() int32 {
	if m != nil {
		return m.RemotePort
	}
	return 0
}

func (m *PortEvent) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PortEvent) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *PortEvent) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PortEvent) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *PortEvent) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *PortEvent) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

type LogEntry struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Event                *Event               `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Entry                string               `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{11}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *LogEntry) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

type UserIntentRequest struct {
	Intent               *Intent  `protobuf:"bytes,1,opt,name=intent,proto3" json:"intent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIntentRequest) Reset()         { *m = UserIntentRequest{} }
func (m *UserIntentRequest) String() string { return proto.CompactTextString(m) }
func (*UserIntentRequest) ProtoMessage()    {}
func (*UserIntentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{12}
}

func (m *UserIntentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIntentRequest.Unmarshal(m, b)
}
func (m *UserIntentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIntentRequest.Marshal(b, m, deterministic)
}
func (m *UserIntentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIntentRequest.Merge(m, src)
}
func (m *UserIntentRequest) XXX_Size() int {
	return xxx_messageInfo_UserIntentRequest.Size(m)
}
func (m *UserIntentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIntentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIntentRequest proto.InternalMessageInfo

func (m *UserIntentRequest) GetIntent() *Intent {
	if m != nil {
		return m.Intent
	}
	return nil
}

type Intent struct {
	Build                bool     `protobuf:"varint,1,opt,name=build,proto3" json:"build,omitempty"`
	Sync                 bool     `protobuf:"varint,2,opt,name=sync,proto3" json:"sync,omitempty"`
	Deploy               bool     `protobuf:"varint,3,opt,name=deploy,proto3" json:"deploy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Intent) Reset()         { *m = Intent{} }
func (m *Intent) String() string { return proto.CompactTextString(m) }
func (*Intent) ProtoMessage()    {}
func (*Intent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2d38e344f9dbf5, []int{13}
}

func (m *Intent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Intent.Unmarshal(m, b)
}
func (m *Intent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Intent.Marshal(b, m, deterministic)
}
func (m *Intent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intent.Merge(m, src)
}
func (m *Intent) XXX_Size() int {
	return xxx_messageInfo_Intent.Size(m)
}
func (m *Intent) XXX_DiscardUnknown() {
	xxx_messageInfo_Intent.DiscardUnknown(m)
}

var xxx_messageInfo_Intent proto.InternalMessageInfo

func (m *Intent) GetBuild() bool {
	if m != nil {
		return m.Build
	}
	return false
}

func (m *Intent) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

func (m *Intent) GetDeploy() bool {
	if m != nil {
		return m.Deploy
	}
	return false
}

func init() {
	proto.RegisterType((*StateResponse)(nil), "proto.StateResponse")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*State)(nil), "proto.State")
	proto.RegisterMapType((map[int32]*PortEvent)(nil), "proto.State.ForwardedPortsEntry")
	proto.RegisterType((*BuildState)(nil), "proto.BuildState")
	proto.RegisterMapType((map[string]string)(nil), "proto.BuildState.ArtifactsEntry")
	proto.RegisterType((*DeployState)(nil), "proto.DeployState")
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterType((*MetaEvent)(nil), "proto.MetaEvent")
	proto.RegisterType((*BuildEvent)(nil), "proto.BuildEvent")
	proto.RegisterType((*DeployEvent)(nil), "proto.DeployEvent")
	proto.RegisterType((*PortEvent)(nil), "proto.PortEvent")
	proto.RegisterType((*LogEntry)(nil), "proto.LogEntry")
	proto.RegisterType((*UserIntentRequest)(nil), "proto.UserIntentRequest")
	proto.RegisterType((*Intent)(nil), "proto.Intent")
}

func init() { proto.RegisterFile("skaffold.proto", fileDescriptor_4f2d38e344f9dbf5) }

var fileDescriptor_4f2d38e344f9dbf5 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4b, 0x8f, 0xe3, 0x44,
	0x10, 0xc6, 0xce, 0x63, 0xec, 0xca, 0x3c, 0x1b, 0x76, 0x15, 0x79, 0x07, 0x18, 0x5a, 0x2c, 0x1a,
	0x71, 0x48, 0x76, 0x67, 0x10, 0xac, 0x46, 0x08, 0x89, 0x81, 0x61, 0x87, 0xd5, 0x80, 0x90, 0xb3,
	0x48, 0xdc, 0x90, 0x27, 0xa9, 0x84, 0x68, 0x1d, 0xb7, 0x71, 0x77, 0x02, 0xb9, 0x70, 0xe0, 0xc8,
	0x1e, 0x11, 0xbf, 0x8c, 0xbf, 0x80, 0xf8, 0x1d, 0xa8, 0xab, 0xbb, 0xe3, 0xf6, 0x3c, 0x24, 0x4e,
	0xee, 0xaa, 0xfa, 0xea, 0x73, 0x55, 0x7d, 0xfd, 0x80, 0x5d, 0xf9, 0x2a, 0x9b, 0x4e, 0x45, 0x3e,
	0x19, 0x94, 0x95, 0x50, 0x82, 0x75, 0xe8, 0x93, 0x1c, 0xce, 0x84, 0x98, 0xe5, 0x38, 0xcc, 0xca,
	0xf9, 0x30, 0x2b, 0x0a, 0xa1, 0x32, 0x35, 0x17, 0x85, 0x34, 0xa0, 0xe4, 0x5d, 0x1b, 0x25, 0xeb,
	0x7a, 0x39, 0x1d, 0xaa, 0xf9, 0x02, 0xa5, 0xca, 0x16, 0xa5, 0x05, 0x3c, 0xba, 0x09, 0xc0, 0x45,
	0xa9, 0xd6, 0x26, 0xc8, 0x4f, 0x61, 0x67, 0xa4, 0x32, 0x85, 0x29, 0xca, 0x52, 0x14, 0x12, 0x19,
	0x87, 0x8e, 0xd4, 0x8e, 0x7e, 0x70, 0x14, 0x1c, 0xf7, 0x4e, 0xb6, 0x0d, 0x6e, 0x60, 0x40, 0x26,
	0xc4, 0x0f, 0x21, 0xda, 0xe0, 0xf7, 0xa1, 0xb5, 0x90, 0x33, 0x42, 0xc7, 0xa9, 0x5e, 0xf2, 0xb7,
	0x61, 0x2b, 0xc5, 0x9f, 0x97, 0x28, 0x15, 0x63, 0xd0, 0x2e, 0xb2, 0x05, 0xda, 0x28, 0xad, 0xf9,
	0x5f, 0x21, 0x74, 0x88, 0x8d, 0x3d, 0x05, 0xb8, 0x5e, 0xce, 0xf3, 0xc9, 0xc8, 0xfb, 0xdf, 0x81,
	0xfd, 0xdf, 0xf9, 0x26, 0x90, 0x7a, 0x20, 0xf6, 0x11, 0xf4, 0x26, 0x58, 0xe6, 0x62, 0x6d, 0x72,
	0x42, 0xca, 0x61, 0x36, 0xe7, 0xcb, 0x3a, 0x92, 0xfa, 0x30, 0x76, 0x09, 0xbb, 0x53, 0x51, 0xfd,
	0x92, 0x55, 0x13, 0x9c, 0x7c, 0x27, 0x2a, 0x25, 0xfb, 0xed, 0xa3, 0xd6, 0x71, 0xef, 0xe4, 0xc8,
	0x6f, 0x6e, 0xf0, 0x55, 0x03, 0x72, 0x51, 0xa8, 0x6a, 0x9d, 0xde, 0xc8, 0x4b, 0x46, 0xf0, 0xe6,
	0x1d, 0x30, 0x3d, 0x84, 0x57, 0xb8, 0xa6, 0x16, 0x3a, 0xa9, 0x5e, 0xb2, 0x0f, 0xa0, 0xb3, 0xca,
	0xf2, 0xa5, 0x2b, 0x71, 0xdf, 0xfe, 0x49, 0xe7, 0x5c, 0xac, 0xb0, 0x50, 0xa9, 0x09, 0x9f, 0x85,
	0xcf, 0x82, 0x17, 0xed, 0xa8, 0xb5, 0xdf, 0xe6, 0x7f, 0x04, 0x00, 0x75, 0xd7, 0xec, 0x33, 0x88,
	0xb3, 0x4a, 0xcd, 0xa7, 0xd9, 0x58, 0xc9, 0x7e, 0xd0, 0x28, 0xb7, 0x46, 0x0d, 0x3e, 0x77, 0x10,
	0x53, 0x6e, 0x9d, 0x92, 0x7c, 0x0a, 0xbb, 0xcd, 0xa0, 0x5f, 0x64, 0x6c, 0x8a, 0x7c, 0xcb, 0x2f,
	0x32, 0xf6, 0x4a, 0xe2, 0x8f, 0xa1, 0xe7, 0x4d, 0x93, 0x3d, 0x84, 0xae, 0x56, 0x7e, 0x29, 0x6d,
	0xb6, 0xb5, 0xf8, 0xbf, 0x01, 0x74, 0xa8, 0x1d, 0xf6, 0x04, 0xe2, 0x05, 0xaa, 0x8c, 0x0c, 0x2b,
	0xa5, 0xeb, 0xf9, 0x1b, 0xe7, 0xbf, 0x7c, 0x23, 0xad, 0x41, 0xec, 0xd4, 0xaa, 0x6f, 0x52, 0xc2,
	0xdb, 0xea, 0xbb, 0x1c, 0x0f, 0xc6, 0x3e, 0x76, 0xfa, 0x9b, 0xac, 0xd6, 0x1d, 0xfa, 0xbb, 0x34,
	0x1f, 0xa8, 0xcb, 0x2b, 0xdd, 0xe8, 0xfb, 0xed, 0xbb, 0x25, 0xd1, 0xe5, 0x6d, 0x40, 0xe7, 0xdb,
	0x00, 0xa8, 0x17, 0x3f, 0xaa, 0x75, 0x89, 0xfc, 0x3d, 0x88, 0x37, 0x6d, 0xe8, 0xb1, 0xa1, 0x9e,
	0xa8, 0x1d, 0x86, 0x31, 0x78, 0x6a, 0xe5, 0x33, 0x98, 0x04, 0x22, 0xa7, 0x85, 0x85, 0x6d, 0x6c,
	0x6f, 0x9a, 0xa1, 0x3f, 0x4d, 0x2d, 0x10, 0x56, 0x15, 0x35, 0x15, 0xa7, 0x7a, 0xc9, 0x3f, 0x71,
	0x32, 0x18, 0xd2, 0x7b, 0x64, 0x70, 0x89, 0x61, 0x9d, 0xf8, 0x3a, 0x84, 0x78, 0xd3, 0x18, 0x3b,
	0x84, 0x38, 0x17, 0xe3, 0x2c, 0xd7, 0x1e, 0xbb, 0x49, 0x6b, 0x07, 0x7b, 0x07, 0xa0, 0xc2, 0x85,
	0x50, 0x48, 0xe1, 0x90, 0xc2, 0x9e, 0x87, 0xf5, 0x61, 0xab, 0x14, 0x93, 0x6f, 0xf5, 0x39, 0x36,
	0xa5, 0x39, 0x93, 0xbd, 0x0f, 0x3b, 0x63, 0x51, 0xa8, 0x6c, 0x5e, 0x60, 0x45, 0xf1, 0x36, 0xc5,
	0x9b, 0x4e, 0xfd, 0x77, 0x7d, 0xf0, 0x65, 0x99, 0x8d, 0xb1, 0xdf, 0x21, 0x44, 0xed, 0xd0, 0x83,
	0xd2, 0x43, 0xa7, 0xf4, 0xae, 0x19, 0x94, 0xb3, 0x19, 0x87, 0xed, 0x0a, 0xa5, 0x58, 0x56, 0x63,
	0x7c, 0xb9, 0x2e, 0xb1, 0xbf, 0x45, 0xf1, 0x86, 0xcf, 0xc7, 0x10, 0x47, 0xd4, 0xc4, 0x68, 0x1f,
	0xff, 0x0d, 0xa2, 0x2b, 0x31, 0x33, 0xa7, 0xe0, 0x19, 0xc4, 0x9b, 0x0b, 0xd2, 0x6e, 0xd4, 0x64,
	0x60, 0x6e, 0xc8, 0x81, 0xbb, 0x21, 0x07, 0x2f, 0x1d, 0x22, 0xad, 0xc1, 0xfa, 0x66, 0x44, 0x6f,
	0xaf, 0xba, 0x9b, 0xd1, 0x1e, 0x67, 0x6c, 0x6e, 0x8d, 0x96, 0xbf, 0x35, 0xce, 0xe0, 0xe0, 0x7b,
	0x89, 0xd5, 0xd7, 0x85, 0xd2, 0x50, 0x7b, 0x37, 0x3e, 0x86, 0xee, 0x9c, 0x1c, 0xb6, 0x8a, 0x1d,
	0xcb, 0x67, 0x51, 0x36, 0xc8, 0x5f, 0x40, 0xd7, 0x78, 0x34, 0x37, 0x9d, 0x04, 0xc2, 0x47, 0xa9,
	0x31, 0xf4, 0x15, 0x2b, 0xd7, 0xc5, 0x98, 0x8a, 0x8a, 0x52, 0x5a, 0xeb, 0x7d, 0x62, 0x36, 0x3f,
	0x95, 0x11, 0xa5, 0xd6, 0x3a, 0x79, 0xdd, 0x82, 0xbd, 0x91, 0x7d, 0x62, 0x46, 0x58, 0xad, 0xe6,
	0x63, 0x64, 0x5f, 0x40, 0xf4, 0x1c, 0x95, 0x3d, 0xe6, 0xb7, 0x06, 0x71, 0xa1, 0x9f, 0x8a, 0xa4,
	0xf1, 0x08, 0xf0, 0x83, 0xdf, 0xff, 0xfe, 0xe7, 0xcf, 0xb0, 0xc7, 0xe2, 0xe1, 0xea, 0xe9, 0x90,
	0x1e, 0x04, 0xf6, 0x1c, 0x22, 0x1a, 0xc3, 0x95, 0x98, 0xb1, 0x3d, 0x0b, 0x76, 0x13, 0x4f, 0x6e,
	0x3a, 0xf8, 0x03, 0x22, 0xd8, 0x63, 0x3b, 0x9a, 0xc0, 0x1c, 0xb2, 0x5c, 0xcc, 0x8e, 0x83, 0x27,
	0x01, 0x3b, 0x87, 0x2e, 0x11, 0xc9, 0xff, 0x41, 0xc3, 0x88, 0x66, 0x9b, 0xc1, 0x86, 0x46, 0x12,
	0xc7, 0x15, 0x74, 0x2f, 0xb3, 0x62, 0x92, 0x23, 0x6b, 0x48, 0x94, 0xdc, 0xd3, 0x1d, 0x3f, 0x24,
	0x9e, 0x87, 0xfc, 0xa0, 0xe6, 0x19, 0xfe, 0x44, 0x04, 0x67, 0xc1, 0x87, 0xec, 0x07, 0xd8, 0xba,
	0xf8, 0x15, 0xc7, 0x4b, 0x85, 0xac, 0x6f, 0xe9, 0x6e, 0x69, 0x79, 0x2f, 0xf5, 0x23, 0xa2, 0x7e,
	0xc0, 0x7b, 0x44, 0x6d, 0x68, 0xce, 0xac, 0xb2, 0xd7, 0x5d, 0x02, 0x9f, 0xfe, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x27, 0x18, 0xf6, 0x7f, 0xf6, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SkaffoldServiceClient is the client API for SkaffoldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SkaffoldServiceClient interface {
	GetState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*State, error)
	EventLog(ctx context.Context, opts ...grpc.CallOption) (SkaffoldService_EventLogClient, error)
	Events(ctx context.Context, opts ...grpc.CallOption) (SkaffoldService_EventsClient, error)
	Handle(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
	Execute(ctx context.Context, in *UserIntentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type skaffoldServiceClient struct {
	cc *grpc.ClientConn
}

func NewSkaffoldServiceClient(cc *grpc.ClientConn) SkaffoldServiceClient {
	return &skaffoldServiceClient{cc}
}

func (c *skaffoldServiceClient) GetState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/proto.SkaffoldService/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skaffoldServiceClient) EventLog(ctx context.Context, opts ...grpc.CallOption) (SkaffoldService_EventLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SkaffoldService_serviceDesc.Streams[0], "/proto.SkaffoldService/EventLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &skaffoldServiceEventLogClient{stream}
	return x, nil
}

type SkaffoldService_EventLogClient interface {
	Send(*LogEntry) error
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type skaffoldServiceEventLogClient struct {
	grpc.ClientStream
}

func (x *skaffoldServiceEventLogClient) Send(m *LogEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *skaffoldServiceEventLogClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skaffoldServiceClient) Events(ctx context.Context, opts ...grpc.CallOption) (SkaffoldService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SkaffoldService_serviceDesc.Streams[1], "/proto.SkaffoldService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &skaffoldServiceEventsClient{stream}
	return x, nil
}

type SkaffoldService_EventsClient interface {
	Send(*LogEntry) error
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type skaffoldServiceEventsClient struct {
	grpc.ClientStream
}

func (x *skaffoldServiceEventsClient) Send(m *LogEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *skaffoldServiceEventsClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *skaffoldServiceClient) Handle(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.SkaffoldService/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skaffoldServiceClient) Execute(ctx context.Context, in *UserIntentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.SkaffoldService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkaffoldServiceServer is the server API for SkaffoldService service.
type SkaffoldServiceServer interface {
	GetState(context.Context, *empty.Empty) (*State, error)
	EventLog(SkaffoldService_EventLogServer) error
	Events(SkaffoldService_EventsServer) error
	Handle(context.Context, *Event) (*empty.Empty, error)
	Execute(context.Context, *UserIntentRequest) (*empty.Empty, error)
}

// UnimplementedSkaffoldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSkaffoldServiceServer struct {
}

func (*UnimplementedSkaffoldServiceServer) GetState(ctx context.Context, req *empty.Empty) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedSkaffoldServiceServer) EventLog(srv SkaffoldService_EventLogServer) error {
	return status.Errorf(codes.Unimplemented, "method EventLog not implemented")
}
func (*UnimplementedSkaffoldServiceServer) Events(srv SkaffoldService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (*UnimplementedSkaffoldServiceServer) Handle(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (*UnimplementedSkaffoldServiceServer) Execute(ctx context.Context, req *UserIntentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterSkaffoldServiceServer(s *grpc.Server, srv SkaffoldServiceServer) {
	s.RegisterService(&_SkaffoldService_serviceDesc, srv)
}

func _SkaffoldService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkaffoldServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SkaffoldService/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkaffoldServiceServer).GetState(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkaffoldService_EventLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SkaffoldServiceServer).EventLog(&skaffoldServiceEventLogServer{stream})
}

type SkaffoldService_EventLogServer interface {
	Send(*LogEntry) error
	Recv() (*LogEntry, error)
	grpc.ServerStream
}

type skaffoldServiceEventLogServer struct {
	grpc.ServerStream
}

func (x *skaffoldServiceEventLogServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *skaffoldServiceEventLogServer) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SkaffoldService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SkaffoldServiceServer).Events(&skaffoldServiceEventsServer{stream})
}

type SkaffoldService_EventsServer interface {
	Send(*LogEntry) error
	Recv() (*LogEntry, error)
	grpc.ServerStream
}

type skaffoldServiceEventsServer struct {
	grpc.ServerStream
}

func (x *skaffoldServiceEventsServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *skaffoldServiceEventsServer) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SkaffoldService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkaffoldServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SkaffoldService/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkaffoldServiceServer).Handle(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkaffoldService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIntentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkaffoldServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SkaffoldService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkaffoldServiceServer).Execute(ctx, req.(*UserIntentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SkaffoldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SkaffoldService",
	HandlerType: (*SkaffoldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _SkaffoldService_GetState_Handler,
		},
		{
			MethodName: "Handle",
			Handler:    _SkaffoldService_Handle_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _SkaffoldService_Execute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventLog",
			Handler:       _SkaffoldService_EventLog_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Events",
			Handler:       _SkaffoldService_Events_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "skaffold.proto",
}
