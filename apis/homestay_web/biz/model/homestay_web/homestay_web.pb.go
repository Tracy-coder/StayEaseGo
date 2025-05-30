// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: homestay_web.proto

package homestay_web

import (
	_ "StayEaseGo/apis/homestay_web/biz/model/api"
	base "StayEaseGo/apis/homestay_web/biz/model/base"
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

type Homestay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                     int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" form:"ID" query:"ID"`
	Title                  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`
	SubTitle               string `protobuf:"bytes,3,opt,name=subTitle,proto3" json:"subTitle,omitempty" form:"subTitle" query:"subTitle"`
	Banner                 string `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner,omitempty" form:"banner" query:"banner"`
	Info                   string `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty" form:"info" query:"info"`
	PeopleNum              int64  `protobuf:"varint,6,opt,name=peopleNum,proto3" json:"peopleNum,omitempty" form:"peopleNum" query:"peopleNum"`                                                     //容纳人的数量
	HomestayBusinessBossID int64  `protobuf:"varint,7,opt,name=homestayBusinessBossID,proto3" json:"homestayBusinessBossID,omitempty" form:"homestayBusinessBossID" query:"homestayBusinessBossID"` //房东id
	RowState               int64  `protobuf:"varint,8,opt,name=rowState,proto3" json:"rowState,omitempty" form:"rowState" query:"rowState"`                                                         //0:下架 1:上架
	RowType                int64  `protobuf:"varint,9,opt,name=rowType,proto3" json:"rowType,omitempty" form:"rowType" query:"rowType"`                                                             //售卖类型0：按房间出售 1:按人次出售
	FoodInfo               string `protobuf:"bytes,10,opt,name=foodInfo,proto3" json:"foodInfo,omitempty" form:"foodInfo" query:"foodInfo"`                                                         //餐食标准
	FoodPrice              int64  `protobuf:"varint,11,opt,name=foodPrice,proto3" json:"foodPrice,omitempty" form:"foodPrice" query:"foodPrice"`                                                    //餐食价格(分)
	HomestayPrice          int64  `protobuf:"varint,12,opt,name=homestayPrice,proto3" json:"homestayPrice,omitempty" form:"homestayPrice" query:"homestayPrice"`                                    //民宿价格(分)
	MarketHomestayPrice    int64  `protobuf:"varint,13,opt,name=marketHomestayPrice,proto3" json:"marketHomestayPrice,omitempty" form:"marketHomestayPrice" query:"marketHomestayPrice"`            //民宿市场价格(分)
}

func (x *Homestay) Reset() {
	*x = Homestay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Homestay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Homestay) ProtoMessage() {}

func (x *Homestay) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Homestay.ProtoReflect.Descriptor instead.
func (*Homestay) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{0}
}

func (x *Homestay) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Homestay) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Homestay) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *Homestay) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *Homestay) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Homestay) GetPeopleNum() int64 {
	if x != nil {
		return x.PeopleNum
	}
	return 0
}

func (x *Homestay) GetHomestayBusinessBossID() int64 {
	if x != nil {
		return x.HomestayBusinessBossID
	}
	return 0
}

func (x *Homestay) GetRowState() int64 {
	if x != nil {
		return x.RowState
	}
	return 0
}

func (x *Homestay) GetRowType() int64 {
	if x != nil {
		return x.RowType
	}
	return 0
}

func (x *Homestay) GetFoodInfo() string {
	if x != nil {
		return x.FoodInfo
	}
	return ""
}

func (x *Homestay) GetFoodPrice() int64 {
	if x != nil {
		return x.FoodPrice
	}
	return 0
}

func (x *Homestay) GetHomestayPrice() int64 {
	if x != nil {
		return x.HomestayPrice
	}
	return 0
}

func (x *Homestay) GetMarketHomestayPrice() int64 {
	if x != nil {
		return x.MarketHomestayPrice
	}
	return 0
}

type HomestayDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" form:"ID" query:"ID"`
}

func (x *HomestayDetailReq) Reset() {
	*x = HomestayDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayDetailReq) ProtoMessage() {}

func (x *HomestayDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayDetailReq.ProtoReflect.Descriptor instead.
func (*HomestayDetailReq) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{1}
}

func (x *HomestayDetailReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type HomestayDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homestay *Homestay `protobuf:"bytes,1,opt,name=homestay,proto3" json:"homestay,omitempty" form:"homestay" query:"homestay"`
}

func (x *HomestayDetailResp) Reset() {
	*x = HomestayDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayDetailResp) ProtoMessage() {}

func (x *HomestayDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayDetailResp.ProtoReflect.Descriptor instead.
func (*HomestayDetailResp) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{2}
}

func (x *HomestayDetailResp) GetHomestay() *Homestay {
	if x != nil {
		return x.Homestay
	}
	return nil
}

type HomestayListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinPrice               int64 `protobuf:"varint,1,opt,name=minPrice,proto3" json:"minPrice,omitempty" form:"minPrice" query:"minPrice"`
	MaxPrice               int64 `protobuf:"varint,2,opt,name=maxPrice,proto3" json:"maxPrice,omitempty" form:"maxPrice" query:"maxPrice"`
	MinPeopleNum           int64 `protobuf:"varint,3,opt,name=minPeopleNum,proto3" json:"minPeopleNum,omitempty" form:"minPeopleNum" query:"minPeopleNum"`
	MaxPeopleNum           int64 `protobuf:"varint,4,opt,name=maxPeopleNum,proto3" json:"maxPeopleNum,omitempty" form:"maxPeopleNum" query:"maxPeopleNum"`
	RowType                int64 `protobuf:"varint,5,opt,name=rowType,proto3" json:"rowType,omitempty" form:"rowType" query:"rowType"`
	Offset                 int64 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty" form:"offset" query:"offset"`
	Limit                  int64 `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty" form:"limit" query:"limit"`
	HomestayBusinessBossID int64 `protobuf:"varint,8,opt,name=homestayBusinessBossID,proto3" json:"homestayBusinessBossID,omitempty" form:"homestayBusinessBossID" query:"homestayBusinessBossID"`
}

func (x *HomestayListReq) Reset() {
	*x = HomestayListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayListReq) ProtoMessage() {}

func (x *HomestayListReq) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayListReq.ProtoReflect.Descriptor instead.
func (*HomestayListReq) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{3}
}

func (x *HomestayListReq) GetMinPrice() int64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *HomestayListReq) GetMaxPrice() int64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *HomestayListReq) GetMinPeopleNum() int64 {
	if x != nil {
		return x.MinPeopleNum
	}
	return 0
}

func (x *HomestayListReq) GetMaxPeopleNum() int64 {
	if x != nil {
		return x.MaxPeopleNum
	}
	return 0
}

func (x *HomestayListReq) GetRowType() int64 {
	if x != nil {
		return x.RowType
	}
	return 0
}

func (x *HomestayListReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *HomestayListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *HomestayListReq) GetHomestayBusinessBossID() int64 {
	if x != nil {
		return x.HomestayBusinessBossID
	}
	return 0
}

type HomestayListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homestays []*Homestay `protobuf:"bytes,1,rep,name=homestays,proto3" json:"homestays,omitempty" form:"homestays" query:"homestays"`
}

func (x *HomestayListResp) Reset() {
	*x = HomestayListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayListResp) ProtoMessage() {}

func (x *HomestayListResp) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayListResp.ProtoReflect.Descriptor instead.
func (*HomestayListResp) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{4}
}

func (x *HomestayListResp) GetHomestays() []*Homestay {
	if x != nil {
		return x.Homestays
	}
	return nil
}

type CreateHomestayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homestay *Homestay `protobuf:"bytes,1,opt,name=homestay,proto3" json:"homestay,omitempty" form:"homestay" query:"homestay"`
}

func (x *CreateHomestayReq) Reset() {
	*x = CreateHomestayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHomestayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHomestayReq) ProtoMessage() {}

func (x *CreateHomestayReq) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHomestayReq.ProtoReflect.Descriptor instead.
func (*CreateHomestayReq) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{5}
}

func (x *CreateHomestayReq) GetHomestay() *Homestay {
	if x != nil {
		return x.Homestay
	}
	return nil
}

type CreateHomestayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homestay *Homestay `protobuf:"bytes,1,opt,name=homestay,proto3" json:"homestay,omitempty" form:"homestay" query:"homestay"`
}

func (x *CreateHomestayResp) Reset() {
	*x = CreateHomestayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHomestayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHomestayResp) ProtoMessage() {}

func (x *CreateHomestayResp) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHomestayResp.ProtoReflect.Descriptor instead.
func (*CreateHomestayResp) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{6}
}

func (x *CreateHomestayResp) GetHomestay() *Homestay {
	if x != nil {
		return x.Homestay
	}
	return nil
}

type HomestayBusinessBossDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" form:"ID" query:"ID"`
}

func (x *HomestayBusinessBossDetailReq) Reset() {
	*x = HomestayBusinessBossDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayBusinessBossDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayBusinessBossDetailReq) ProtoMessage() {}

func (x *HomestayBusinessBossDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayBusinessBossDetailReq.ProtoReflect.Descriptor instead.
func (*HomestayBusinessBossDetailReq) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{7}
}

func (x *HomestayBusinessBossDetailReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type HomestayBusinessBossDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty" form:"mobile" query:"mobile"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty" form:"nickname" query:"nickname"`
	Sex      int32  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty" form:"sex" query:"sex"`
	Avatar   string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty" form:"avatar" query:"avatar"`
	Info     string `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty" form:"info" query:"info"`
}

func (x *HomestayBusinessBossDetailResp) Reset() {
	*x = HomestayBusinessBossDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homestay_web_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayBusinessBossDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayBusinessBossDetailResp) ProtoMessage() {}

func (x *HomestayBusinessBossDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_homestay_web_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayBusinessBossDetailResp.ProtoReflect.Descriptor instead.
func (*HomestayBusinessBossDetailResp) Descriptor() ([]byte, []int) {
	return file_homestay_web_proto_rawDescGZIP(), []int{8}
}

func (x *HomestayBusinessBossDetailResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomestayBusinessBossDetailResp) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *HomestayBusinessBossDetailResp) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *HomestayBusinessBossDetailResp) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *HomestayBusinessBossDetailResp) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *HomestayBusinessBossDetailResp) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_homestay_web_proto protoreflect.FileDescriptor

var file_homestay_web_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77,
	0x65, 0x62, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x08, 0x48, 0x6f,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x36, 0x0a, 0x16, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x16, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x6f, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x66, 0x6f, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x30, 0x0a, 0x13, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x12, 0x48, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a,
	0x08, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x22, 0x91, 0x02, 0x0a, 0x0f, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x36, 0x0a,
	0x16, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x42, 0x6f, 0x73, 0x73, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x68,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42,
	0x6f, 0x73, 0x73, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x10, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x68, 0x6f, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x79, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x73, 0x22,
	0x47, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x52, 0x08,
	0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32,
	0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e,
	0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x79, 0x22, 0x2f, 0x0a, 0x1d, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0xa2, 0x01, 0x0a, 0x1e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x03, 0x0a, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x79, 0x12, 0x70, 0x0a, 0x0e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0xd2, 0xc1, 0x18, 0x17,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79,
	0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x68, 0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x54, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77,
	0x65, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x12, 0x92, 0x01, 0x0a, 0x1a, 0x68, 0x6f, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x79, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77,
	0x65, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x42, 0x6f, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x2f, 0x62, 0x6f, 0x73, 0x73, 0x42, 0x35, 0x5a, 0x33,
	0x53, 0x74, 0x61, 0x79, 0x45, 0x61, 0x73, 0x65, 0x47, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x2f, 0x62, 0x69, 0x7a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x5f,
	0x77, 0x65, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_homestay_web_proto_rawDescOnce sync.Once
	file_homestay_web_proto_rawDescData = file_homestay_web_proto_rawDesc
)

func file_homestay_web_proto_rawDescGZIP() []byte {
	file_homestay_web_proto_rawDescOnce.Do(func() {
		file_homestay_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_homestay_web_proto_rawDescData)
	})
	return file_homestay_web_proto_rawDescData
}

var file_homestay_web_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_homestay_web_proto_goTypes = []interface{}{
	(*Homestay)(nil),                       // 0: homestay_web.Homestay
	(*HomestayDetailReq)(nil),              // 1: homestay_web.HomestayDetailReq
	(*HomestayDetailResp)(nil),             // 2: homestay_web.HomestayDetailResp
	(*HomestayListReq)(nil),                // 3: homestay_web.HomestayListReq
	(*HomestayListResp)(nil),               // 4: homestay_web.HomestayListResp
	(*CreateHomestayReq)(nil),              // 5: homestay_web.CreateHomestayReq
	(*CreateHomestayResp)(nil),             // 6: homestay_web.CreateHomestayResp
	(*HomestayBusinessBossDetailReq)(nil),  // 7: homestay_web.HomestayBusinessBossDetailReq
	(*HomestayBusinessBossDetailResp)(nil), // 8: homestay_web.HomestayBusinessBossDetailResp
	(*base.Empty)(nil),                     // 9: base.Empty
}
var file_homestay_web_proto_depIdxs = []int32{
	0, // 0: homestay_web.HomestayDetailResp.homestay:type_name -> homestay_web.Homestay
	0, // 1: homestay_web.HomestayListResp.homestays:type_name -> homestay_web.Homestay
	0, // 2: homestay_web.CreateHomestayReq.homestay:type_name -> homestay_web.Homestay
	0, // 3: homestay_web.CreateHomestayResp.homestay:type_name -> homestay_web.Homestay
	1, // 4: homestay_web.homestay.homestayDetail:input_type -> homestay_web.HomestayDetailReq
	3, // 5: homestay_web.homestay.homestayList:input_type -> homestay_web.HomestayListReq
	5, // 6: homestay_web.homestay.createHomestay:input_type -> homestay_web.CreateHomestayReq
	7, // 7: homestay_web.homestay.homestayBusinessBossDetail:input_type -> homestay_web.HomestayBusinessBossDetailReq
	2, // 8: homestay_web.homestay.homestayDetail:output_type -> homestay_web.HomestayDetailResp
	4, // 9: homestay_web.homestay.homestayList:output_type -> homestay_web.HomestayListResp
	9, // 10: homestay_web.homestay.createHomestay:output_type -> base.Empty
	8, // 11: homestay_web.homestay.homestayBusinessBossDetail:output_type -> homestay_web.HomestayBusinessBossDetailResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_homestay_web_proto_init() }
func file_homestay_web_proto_init() {
	if File_homestay_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_homestay_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Homestay); i {
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
		file_homestay_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayDetailReq); i {
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
		file_homestay_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayDetailResp); i {
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
		file_homestay_web_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayListReq); i {
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
		file_homestay_web_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayListResp); i {
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
		file_homestay_web_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHomestayReq); i {
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
		file_homestay_web_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHomestayResp); i {
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
		file_homestay_web_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayBusinessBossDetailReq); i {
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
		file_homestay_web_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayBusinessBossDetailResp); i {
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
			RawDescriptor: file_homestay_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_homestay_web_proto_goTypes,
		DependencyIndexes: file_homestay_web_proto_depIdxs,
		MessageInfos:      file_homestay_web_proto_msgTypes,
	}.Build()
	File_homestay_web_proto = out.File
	file_homestay_web_proto_rawDesc = nil
	file_homestay_web_proto_goTypes = nil
	file_homestay_web_proto_depIdxs = nil
}
