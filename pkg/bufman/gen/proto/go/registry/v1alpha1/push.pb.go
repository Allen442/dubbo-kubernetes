// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: registry/v1alpha1/push.proto

package registryv1alpha1

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

import (
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/module/v1alpha1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PushRequest specifies the module to push to the BSR.
type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// Deprecated: Do not use.
	Branch string           `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Module *v1alpha1.Module `protobuf:"bytes,4,opt,name=module,proto3" json:"module,omitempty"`
	// Optional; if provided, the provided tags
	// are created for the pushed commit.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional; if provided, the pushed commit
	// will be appended to these tracks. If the
	// tracks do not exist, they will be created.
	//
	// Deprecated: Do not use.
	Tracks []string `protobuf:"bytes,6,rep,name=tracks,proto3" json:"tracks,omitempty"`
	// If non-empty, the push creates a draft commit with this name.
	DraftName string `protobuf:"bytes,7,opt,name=draft_name,json=draftName,proto3" json:"draft_name,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_push_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PushRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

// Deprecated: Do not use.
func (x *PushRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *PushRequest) GetModule() *v1alpha1.Module {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *PushRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Deprecated: Do not use.
func (x *PushRequest) GetTracks() []string {
	if x != nil {
		return x.Tracks
	}
	return nil
}

func (x *PushRequest) GetDraftName() string {
	if x != nil {
		return x.DraftName
	}
	return ""
}

// PushResponse is the pushed module pin, local to the used remote.
type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalModulePin *LocalModulePin `protobuf:"bytes,5,opt,name=local_module_pin,json=localModulePin,proto3" json:"local_module_pin,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_push_proto_rawDescGZIP(), []int{1}
}

func (x *PushResponse) GetLocalModulePin() *LocalModulePin {
	if x != nil {
		return x.LocalModulePin
	}
	return nil
}

// PushManifestAndBlobsRequest holds the module to push in the manifest+blobs
// encoding format.
type PushManifestAndBlobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// Manifest with all the module files being pushed.
	Manifest *v1alpha1.Blob `protobuf:"bytes,3,opt,name=manifest,proto3" json:"manifest,omitempty"`
	// Referenced blobs in the manifest. Keep in mind there is not necessarily one
	// blob per file, but one blob per digest, so for files with exactly the same
	// content, you can send just one blob.
	Blobs []*v1alpha1.Blob `protobuf:"bytes,4,rep,name=blobs,proto3" json:"blobs,omitempty"`
	// Optional; if provided, the provided tags
	// are created for the pushed commit.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// If non-empty, the push creates a draft commit with this name.
	DraftName string `protobuf:"bytes,6,opt,name=draft_name,json=draftName,proto3" json:"draft_name,omitempty"`
}

func (x *PushManifestAndBlobsRequest) Reset() {
	*x = PushManifestAndBlobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushManifestAndBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushManifestAndBlobsRequest) ProtoMessage() {}

func (x *PushManifestAndBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushManifestAndBlobsRequest.ProtoReflect.Descriptor instead.
func (*PushManifestAndBlobsRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_push_proto_rawDescGZIP(), []int{2}
}

func (x *PushManifestAndBlobsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PushManifestAndBlobsRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *PushManifestAndBlobsRequest) GetManifest() *v1alpha1.Blob {
	if x != nil {
		return x.Manifest
	}
	return nil
}

func (x *PushManifestAndBlobsRequest) GetBlobs() []*v1alpha1.Blob {
	if x != nil {
		return x.Blobs
	}
	return nil
}

func (x *PushManifestAndBlobsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PushManifestAndBlobsRequest) GetDraftName() string {
	if x != nil {
		return x.DraftName
	}
	return ""
}

// PushManifestAndBlobsResponse is the pushed module pin, local to the used
// remote.
type PushManifestAndBlobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalModulePin *LocalModulePin `protobuf:"bytes,1,opt,name=local_module_pin,json=localModulePin,proto3" json:"local_module_pin,omitempty"`
}

func (x *PushManifestAndBlobsResponse) Reset() {
	*x = PushManifestAndBlobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_push_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushManifestAndBlobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushManifestAndBlobsResponse) ProtoMessage() {}

func (x *PushManifestAndBlobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_push_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushManifestAndBlobsResponse.ProtoReflect.Descriptor instead.
func (*PushManifestAndBlobsResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_push_proto_rawDescGZIP(), []int{3}
}

func (x *PushManifestAndBlobsResponse) GetLocalModulePin() *LocalModulePin {
	if x != nil {
		return x.LocalModulePin
	}
	return nil
}

var File_registry_v1alpha1_push_proto protoreflect.FileDescriptor

var file_registry_v1alpha1_push_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29,
	0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x47, 0x0a, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x75, 0x66, 0x6d,
	0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x72, 0x61, 0x66, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x73, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x62, 0x75,
	0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x22, 0x96, 0x02, 0x0a, 0x1b, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x08,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x08, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e,
	0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x72, 0x61, 0x66, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x83, 0x01, 0x0a, 0x1c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x41, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x70, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x62, 0x75, 0x66,
	0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x50, 0x69, 0x6e, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x50, 0x69, 0x6e, 0x32, 0xba, 0x02, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x36, 0x2e,
	0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x90, 0x02, 0x02, 0x12, 0xac, 0x01, 0x0a, 0x14, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x46, 0x2e, 0x62,
	0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x41, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x41, 0x6e, 0x64,
	0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90,
	0x02, 0x02, 0x42, 0xe4, 0x02, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x05, 0x42, 0x44, 0x41, 0x4f, 0x52, 0xaa, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x2e, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x4f,
	0x72, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75,
	0x62, 0x62, 0x6f, 0x5c, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x35, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x5c,
	0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x42, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x3a, 0x3a, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_registry_v1alpha1_push_proto_rawDescOnce sync.Once
	file_registry_v1alpha1_push_proto_rawDescData = file_registry_v1alpha1_push_proto_rawDesc
)

func file_registry_v1alpha1_push_proto_rawDescGZIP() []byte {
	file_registry_v1alpha1_push_proto_rawDescOnce.Do(func() {
		file_registry_v1alpha1_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_v1alpha1_push_proto_rawDescData)
	})
	return file_registry_v1alpha1_push_proto_rawDescData
}

var file_registry_v1alpha1_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_registry_v1alpha1_push_proto_goTypes = []interface{}{
	(*PushRequest)(nil),                  // 0: bufman.dubbo.apache.org.registry.v1alpha1.PushRequest
	(*PushResponse)(nil),                 // 1: bufman.dubbo.apache.org.registry.v1alpha1.PushResponse
	(*PushManifestAndBlobsRequest)(nil),  // 2: bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsRequest
	(*PushManifestAndBlobsResponse)(nil), // 3: bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsResponse
	(*v1alpha1.Module)(nil),              // 4: bufman.dubbo.apache.org.module.v1alpha1.Module
	(*LocalModulePin)(nil),               // 5: bufman.dubbo.apache.org.registry.v1alpha1.LocalModulePin
	(*v1alpha1.Blob)(nil),                // 6: bufman.dubbo.apache.org.module.v1alpha1.Blob
}
var file_registry_v1alpha1_push_proto_depIdxs = []int32{
	4, // 0: bufman.dubbo.apache.org.registry.v1alpha1.PushRequest.module:type_name -> bufman.dubbo.apache.org.module.v1alpha1.Module
	5, // 1: bufman.dubbo.apache.org.registry.v1alpha1.PushResponse.local_module_pin:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.LocalModulePin
	6, // 2: bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsRequest.manifest:type_name -> bufman.dubbo.apache.org.module.v1alpha1.Blob
	6, // 3: bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsRequest.blobs:type_name -> bufman.dubbo.apache.org.module.v1alpha1.Blob
	5, // 4: bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsResponse.local_module_pin:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.LocalModulePin
	0, // 5: bufman.dubbo.apache.org.registry.v1alpha1.PushService.Push:input_type -> bufman.dubbo.apache.org.registry.v1alpha1.PushRequest
	2, // 6: bufman.dubbo.apache.org.registry.v1alpha1.PushService.PushManifestAndBlobs:input_type -> bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsRequest
	1, // 7: bufman.dubbo.apache.org.registry.v1alpha1.PushService.Push:output_type -> bufman.dubbo.apache.org.registry.v1alpha1.PushResponse
	3, // 8: bufman.dubbo.apache.org.registry.v1alpha1.PushService.PushManifestAndBlobs:output_type -> bufman.dubbo.apache.org.registry.v1alpha1.PushManifestAndBlobsResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_registry_v1alpha1_push_proto_init() }
func file_registry_v1alpha1_push_proto_init() {
	if File_registry_v1alpha1_push_proto != nil {
		return
	}
	file_registry_v1alpha1_module_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_registry_v1alpha1_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_registry_v1alpha1_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_registry_v1alpha1_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushManifestAndBlobsRequest); i {
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
		file_registry_v1alpha1_push_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushManifestAndBlobsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_registry_v1alpha1_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_v1alpha1_push_proto_goTypes,
		DependencyIndexes: file_registry_v1alpha1_push_proto_depIdxs,
		MessageInfos:      file_registry_v1alpha1_push_proto_msgTypes,
	}.Build()
	File_registry_v1alpha1_push_proto = out.File
	file_registry_v1alpha1_push_proto_rawDesc = nil
	file_registry_v1alpha1_push_proto_goTypes = nil
	file_registry_v1alpha1_push_proto_depIdxs = nil
}
