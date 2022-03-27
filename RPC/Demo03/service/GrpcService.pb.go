// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: GrpcService.proto

package service

import (
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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_GrpcService_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Rectangle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chang int32 `protobuf:"varint,1,opt,name=chang,proto3" json:"chang,omitempty"`
	Kuan  int32 `protobuf:"varint,2,opt,name=kuan,proto3" json:"kuan,omitempty"`
}

func (x *Rectangle) Reset() {
	*x = Rectangle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rectangle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rectangle) ProtoMessage() {}

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rectangle.ProtoReflect.Descriptor instead.
func (*Rectangle) Descriptor() ([]byte, []int) {
	return file_GrpcService_proto_rawDescGZIP(), []int{1}
}

func (x *Rectangle) GetChang() int32 {
	if x != nil {
		return x.Chang
	}
	return 0
}

func (x *Rectangle) GetKuan() int32 {
	if x != nil {
		return x.Kuan
	}
	return 0
}

type RouteNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note int32 `protobuf:"varint,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *RouteNote) Reset() {
	*x = RouteNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNote) ProtoMessage() {}

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNote.ProtoReflect.Descriptor instead.
func (*RouteNote) Descriptor() ([]byte, []int) {
	return file_GrpcService_proto_rawDescGZIP(), []int{2}
}

func (x *RouteNote) GetNote() int32 {
	if x != nil {
		return x.Note
	}
	return 0
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature string `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_GrpcService_proto_rawDescGZIP(), []int{3}
}

func (x *Feature) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

type RouteSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteSummary string `protobuf:"bytes,1,opt,name=route_summary,json=routeSummary,proto3" json:"route_summary,omitempty"`
}

func (x *RouteSummary) Reset() {
	*x = RouteSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GrpcService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteSummary) ProtoMessage() {}

func (x *RouteSummary) ProtoReflect() protoreflect.Message {
	mi := &file_GrpcService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteSummary.ProtoReflect.Descriptor instead.
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return file_GrpcService_proto_rawDescGZIP(), []int{4}
}

func (x *RouteSummary) GetRouteSummary() string {
	if x != nil {
		return x.RouteSummary
	}
	return ""
}

var File_GrpcService_proto protoreflect.FileDescriptor

var file_GrpcService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x22, 0x41, 0x0a, 0x05, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x35,
	0x0a, 0x09, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x75, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6b, 0x75, 0x61, 0x6e, 0x22, 0x1f, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x33, 0x0a, 0x0c, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x32, 0xe5, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12,
	0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0d, 0x2e,
	0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x44,
	0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12,
	0x11, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x1a, 0x0f, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x37, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x44,
	0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x1a,
	0x11, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x30, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GrpcService_proto_rawDescOnce sync.Once
	file_GrpcService_proto_rawDescData = file_GrpcService_proto_rawDesc
)

func file_GrpcService_proto_rawDescGZIP() []byte {
	file_GrpcService_proto_rawDescOnce.Do(func() {
		file_GrpcService_proto_rawDescData = protoimpl.X.CompressGZIP(file_GrpcService_proto_rawDescData)
	})
	return file_GrpcService_proto_rawDescData
}

var file_GrpcService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_GrpcService_proto_goTypes = []interface{}{
	(*Point)(nil),        // 0: Demo03.Point
	(*Rectangle)(nil),    // 1: Demo03.Rectangle
	(*RouteNote)(nil),    // 2: Demo03.RouteNote
	(*Feature)(nil),      // 3: Demo03.Feature
	(*RouteSummary)(nil), // 4: Demo03.RouteSummary
}
var file_GrpcService_proto_depIdxs = []int32{
	0, // 0: Demo03.RouteGuide.GetFeature:input_type -> Demo03.Point
	1, // 1: Demo03.RouteGuide.ListFeatures:input_type -> Demo03.Rectangle
	0, // 2: Demo03.RouteGuide.RecordRoute:input_type -> Demo03.Point
	2, // 3: Demo03.RouteGuide.RouteChat:input_type -> Demo03.RouteNote
	3, // 4: Demo03.RouteGuide.GetFeature:output_type -> Demo03.Feature
	3, // 5: Demo03.RouteGuide.ListFeatures:output_type -> Demo03.Feature
	4, // 6: Demo03.RouteGuide.RecordRoute:output_type -> Demo03.RouteSummary
	2, // 7: Demo03.RouteGuide.RouteChat:output_type -> Demo03.RouteNote
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GrpcService_proto_init() }
func file_GrpcService_proto_init() {
	if File_GrpcService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GrpcService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_GrpcService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rectangle); i {
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
		file_GrpcService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNote); i {
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
		file_GrpcService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_GrpcService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteSummary); i {
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
			RawDescriptor: file_GrpcService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_GrpcService_proto_goTypes,
		DependencyIndexes: file_GrpcService_proto_depIdxs,
		MessageInfos:      file_GrpcService_proto_msgTypes,
	}.Build()
	File_GrpcService_proto = out.File
	file_GrpcService_proto_rawDesc = nil
	file_GrpcService_proto_goTypes = nil
	file_GrpcService_proto_depIdxs = nil
}
