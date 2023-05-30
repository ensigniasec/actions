// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: spdx/v1/spdx.proto

package spdxv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseListVersion string                 `protobuf:"bytes,1,opt,name=license_list_version,json=licenseListVersion,proto3" json:"license_list_version,omitempty"`
	Creators           []*Creator             `protobuf:"bytes,2,rep,name=creators,proto3" json:"creators,omitempty"`
	Created            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Comment            string                 `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreationInfo) Reset() {
	*x = CreationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreationInfo) ProtoMessage() {}

func (x *CreationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreationInfo.ProtoReflect.Descriptor instead.
func (*CreationInfo) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{0}
}

func (x *CreationInfo) GetLicenseListVersion() string {
	if x != nil {
		return x.LicenseListVersion
	}
	return ""
}

func (x *CreationInfo) GetCreators() []*Creator {
	if x != nil {
		return x.Creators
	}
	return nil
}

func (x *CreationInfo) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CreationInfo) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Creator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatorType string `protobuf:"bytes,2,opt,name=creator_type,json=creatorType,proto3" json:"creator_type,omitempty"`
}

func (x *Creator) Reset() {
	*x = Creator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Creator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Creator) ProtoMessage() {}

func (x *Creator) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Creator.ProtoReflect.Descriptor instead.
func (*Creator) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{1}
}

func (x *Creator) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Creator) GetCreatorType() string {
	if x != nil {
		return x.CreatorType
	}
	return ""
}

type ExternalRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceCategory string `protobuf:"bytes,1,opt,name=reference_category,json=referenceCategory,proto3" json:"reference_category,omitempty"`
	ReferenceType     string `protobuf:"bytes,2,opt,name=reference_type,json=referenceType,proto3" json:"reference_type,omitempty"`
	ReferenceLocator  string `protobuf:"bytes,3,opt,name=reference_locator,json=referenceLocator,proto3" json:"reference_locator,omitempty"`
	Comment           string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ExternalRef) Reset() {
	*x = ExternalRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalRef) ProtoMessage() {}

func (x *ExternalRef) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalRef.ProtoReflect.Descriptor instead.
func (*ExternalRef) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{2}
}

func (x *ExternalRef) GetReferenceCategory() string {
	if x != nil {
		return x.ReferenceCategory
	}
	return ""
}

func (x *ExternalRef) GetReferenceType() string {
	if x != nil {
		return x.ReferenceType
	}
	return ""
}

func (x *ExternalRef) GetReferenceLocator() string {
	if x != nil {
		return x.ReferenceLocator
	}
	return ""
}

func (x *ExternalRef) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VersionInfo      string         `protobuf:"bytes,3,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	DownloadLocation string         `protobuf:"bytes,4,opt,name=download_location,json=downloadLocation,proto3" json:"download_location,omitempty"`
	SourceInfo       string         `protobuf:"bytes,5,opt,name=source_info,json=sourceInfo,proto3" json:"source_info,omitempty"`
	LicenseConcluded string         `protobuf:"bytes,6,opt,name=license_concluded,json=licenseConcluded,proto3" json:"license_concluded,omitempty"`
	LicenseDeclared  string         `protobuf:"bytes,7,opt,name=license_declared,json=licenseDeclared,proto3" json:"license_declared,omitempty"`
	CopyrightText    string         `protobuf:"bytes,8,opt,name=copyright_text,json=copyrightText,proto3" json:"copyright_text,omitempty"`
	ExternalRefs     []*ExternalRef `protobuf:"bytes,9,rep,name=external_refs,json=externalRefs,proto3" json:"external_refs,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{3}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *Package) GetDownloadLocation() string {
	if x != nil {
		return x.DownloadLocation
	}
	return ""
}

func (x *Package) GetSourceInfo() string {
	if x != nil {
		return x.SourceInfo
	}
	return ""
}

func (x *Package) GetLicenseConcluded() string {
	if x != nil {
		return x.LicenseConcluded
	}
	return ""
}

func (x *Package) GetLicenseDeclared() string {
	if x != nil {
		return x.LicenseDeclared
	}
	return ""
}

func (x *Package) GetCopyrightText() string {
	if x != nil {
		return x.CopyrightText
	}
	return ""
}

func (x *Package) GetExternalRefs() []*ExternalRef {
	if x != nil {
		return x.ExternalRefs
	}
	return nil
}

// Relationship is a Relationship section of an SPDX Document
type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 11.1: Relationship
	// Cardinality: optional, one or more; one per Relationship
	//
	//	one mandatory for SPDX Document with multiple packages
	//
	// RefA and RefB are first and second item
	// Relationship is type from 11.1.1
	SpdxElementId      *DocElementID `protobuf:"bytes,1,opt,name=spdx_element_id,json=spdxElementId,proto3" json:"spdx_element_id,omitempty"`
	RelatedSpdxElement *DocElementID `protobuf:"bytes,2,opt,name=related_spdx_element,json=relatedSpdxElement,proto3" json:"related_spdx_element,omitempty"`
	Comment            string        `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	RelationshipType   string        `protobuf:"bytes,4,opt,name=relationship_type,json=relationshipType,proto3" json:"relationship_type,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{4}
}

func (x *Relationship) GetSpdxElementId() *DocElementID {
	if x != nil {
		return x.SpdxElementId
	}
	return nil
}

func (x *Relationship) GetRelatedSpdxElement() *DocElementID {
	if x != nil {
		return x.RelatedSpdxElement
	}
	return nil
}

func (x *Relationship) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Relationship) GetRelationshipType() string {
	if x != nil {
		return x.RelationshipType
	}
	return ""
}

type DocElementID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentRefId string     `protobuf:"bytes,1,opt,name=document_ref_id,json=documentRefId,proto3" json:"document_ref_id,omitempty"`
	ElementRefId  *ElementID `protobuf:"bytes,2,opt,name=element_ref_id,json=elementRefId,proto3" json:"element_ref_id,omitempty"`
	SpecialId     string     `protobuf:"bytes,3,opt,name=special_id,json=specialId,proto3" json:"special_id,omitempty"`
}

func (x *DocElementID) Reset() {
	*x = DocElementID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocElementID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocElementID) ProtoMessage() {}

func (x *DocElementID) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocElementID.ProtoReflect.Descriptor instead.
func (*DocElementID) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{5}
}

func (x *DocElementID) GetDocumentRefId() string {
	if x != nil {
		return x.DocumentRefId
	}
	return ""
}

func (x *DocElementID) GetElementRefId() *ElementID {
	if x != nil {
		return x.ElementRefId
	}
	return nil
}

func (x *DocElementID) GetSpecialId() string {
	if x != nil {
		return x.SpecialId
	}
	return ""
}

type ElementID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ElementID) Reset() {
	*x = ElementID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElementID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElementID) ProtoMessage() {}

func (x *ElementID) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElementID.ProtoReflect.Descriptor instead.
func (*ElementID) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{6}
}

func (x *ElementID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpdxVersion       string          `protobuf:"bytes,1,opt,name=spdx_version,json=spdxVersion,proto3" json:"spdx_version,omitempty"`
	DataLicense       string          `protobuf:"bytes,2,opt,name=data_license,json=dataLicense,proto3" json:"data_license,omitempty"`
	Name              string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DocumentNamespace string          `protobuf:"bytes,5,opt,name=document_namespace,json=documentNamespace,proto3" json:"document_namespace,omitempty"`
	CreationInfo      *CreationInfo   `protobuf:"bytes,6,opt,name=creation_info,json=creationInfo,proto3" json:"creation_info,omitempty"`
	Packages          []*Package      `protobuf:"bytes,7,rep,name=packages,proto3" json:"packages,omitempty"`
	Relationships     []*Relationship `protobuf:"bytes,8,rep,name=relationships,proto3" json:"relationships,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spdx_v1_spdx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_spdx_v1_spdx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_spdx_v1_spdx_proto_rawDescGZIP(), []int{7}
}

func (x *Document) GetSpdxVersion() string {
	if x != nil {
		return x.SpdxVersion
	}
	return ""
}

func (x *Document) GetDataLicense() string {
	if x != nil {
		return x.DataLicense
	}
	return ""
}

func (x *Document) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Document) GetDocumentNamespace() string {
	if x != nil {
		return x.DocumentNamespace
	}
	return ""
}

func (x *Document) GetCreationInfo() *CreationInfo {
	if x != nil {
		return x.CreationInfo
	}
	return nil
}

func (x *Document) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *Document) GetRelationships() []*Relationship {
	if x != nil {
		return x.Relationships
	}
	return nil
}

var File_spdx_v1_spdx_proto protoreflect.FileDescriptor

var file_spdx_v1_spdx_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x70, 0x64, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x30, 0x0a, 0x14, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x46, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc8, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x65,
	0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70,
	0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x66, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x73, 0x22,
	0xdd, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x3d, 0x0a, 0x0f, 0x73, 0x70, 0x64, 0x78, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x64, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x0d, 0x73, 0x70, 0x64, 0x78, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x47, 0x0a, 0x14, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x64, 0x78, 0x5f,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x70, 0x64,
	0x78, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x8f, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x49,
	0x64, 0x22, 0x1b, 0x0a, 0x09, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba,
	0x02, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x70, 0x64, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x70, 0x64, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x70,
	0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0d, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x42, 0x8a, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x70, 0x64,
	0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x31, 0x65, 0x6e, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x61, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73,
	0x70, 0x64, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x64, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x70, 0x64, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x53, 0x70, 0x64, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x70, 0x64, 0x78, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08,
	0x53, 0x70, 0x64, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spdx_v1_spdx_proto_rawDescOnce sync.Once
	file_spdx_v1_spdx_proto_rawDescData = file_spdx_v1_spdx_proto_rawDesc
)

func file_spdx_v1_spdx_proto_rawDescGZIP() []byte {
	file_spdx_v1_spdx_proto_rawDescOnce.Do(func() {
		file_spdx_v1_spdx_proto_rawDescData = protoimpl.X.CompressGZIP(file_spdx_v1_spdx_proto_rawDescData)
	})
	return file_spdx_v1_spdx_proto_rawDescData
}

var file_spdx_v1_spdx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spdx_v1_spdx_proto_goTypes = []interface{}{
	(*CreationInfo)(nil),          // 0: spdx.v1.CreationInfo
	(*Creator)(nil),               // 1: spdx.v1.Creator
	(*ExternalRef)(nil),           // 2: spdx.v1.ExternalRef
	(*Package)(nil),               // 3: spdx.v1.Package
	(*Relationship)(nil),          // 4: spdx.v1.Relationship
	(*DocElementID)(nil),          // 5: spdx.v1.DocElementID
	(*ElementID)(nil),             // 6: spdx.v1.ElementID
	(*Document)(nil),              // 7: spdx.v1.Document
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_spdx_v1_spdx_proto_depIdxs = []int32{
	1, // 0: spdx.v1.CreationInfo.creators:type_name -> spdx.v1.Creator
	8, // 1: spdx.v1.CreationInfo.created:type_name -> google.protobuf.Timestamp
	2, // 2: spdx.v1.Package.external_refs:type_name -> spdx.v1.ExternalRef
	5, // 3: spdx.v1.Relationship.spdx_element_id:type_name -> spdx.v1.DocElementID
	5, // 4: spdx.v1.Relationship.related_spdx_element:type_name -> spdx.v1.DocElementID
	6, // 5: spdx.v1.DocElementID.element_ref_id:type_name -> spdx.v1.ElementID
	0, // 6: spdx.v1.Document.creation_info:type_name -> spdx.v1.CreationInfo
	3, // 7: spdx.v1.Document.packages:type_name -> spdx.v1.Package
	4, // 8: spdx.v1.Document.relationships:type_name -> spdx.v1.Relationship
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_spdx_v1_spdx_proto_init() }
func file_spdx_v1_spdx_proto_init() {
	if File_spdx_v1_spdx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spdx_v1_spdx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreationInfo); i {
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
		file_spdx_v1_spdx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Creator); i {
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
		file_spdx_v1_spdx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalRef); i {
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
		file_spdx_v1_spdx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_spdx_v1_spdx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relationship); i {
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
		file_spdx_v1_spdx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocElementID); i {
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
		file_spdx_v1_spdx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElementID); i {
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
		file_spdx_v1_spdx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
			RawDescriptor: file_spdx_v1_spdx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spdx_v1_spdx_proto_goTypes,
		DependencyIndexes: file_spdx_v1_spdx_proto_depIdxs,
		MessageInfos:      file_spdx_v1_spdx_proto_msgTypes,
	}.Build()
	File_spdx_v1_spdx_proto = out.File
	file_spdx_v1_spdx_proto_rawDesc = nil
	file_spdx_v1_spdx_proto_goTypes = nil
	file_spdx_v1_spdx_proto_depIdxs = nil
}
