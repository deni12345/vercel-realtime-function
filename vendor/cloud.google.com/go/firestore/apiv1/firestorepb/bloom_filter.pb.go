// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/firestore/v1/bloom_filter.proto

package firestorepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A sequence of bits, encoded in a byte array.
//
// Each byte in the `bitmap` byte array stores 8 bits of the sequence. The only
// exception is the last byte, which may store 8 _or fewer_ bits. The `padding`
// defines the number of bits of the last byte to be ignored as "padding". The
// values of these "padding" bits are unspecified and must be ignored.
//
// To retrieve the first bit, bit 0, calculate: `(bitmap[0] & 0x01) != 0`.
// To retrieve the second bit, bit 1, calculate: `(bitmap[0] & 0x02) != 0`.
// To retrieve the third bit, bit 2, calculate: `(bitmap[0] & 0x04) != 0`.
// To retrieve the fourth bit, bit 3, calculate: `(bitmap[0] & 0x08) != 0`.
// To retrieve bit n, calculate: `(bitmap[n / 8] & (0x01 << (n % 8))) != 0`.
//
// The "size" of a `BitSequence` (the number of bits it contains) is calculated
// by this formula: `(bitmap.length * 8) - padding`.
type BitSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes that encode the bit sequence.
	// May have a length of zero.
	Bitmap []byte `protobuf:"bytes,1,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
	// The number of bits of the last byte in `bitmap` to ignore as "padding".
	// If the length of `bitmap` is zero, then this value must be `0`.
	// Otherwise, this value must be between 0 and 7, inclusive.
	Padding int32 `protobuf:"varint,2,opt,name=padding,proto3" json:"padding,omitempty"`
}

func (x *BitSequence) Reset() {
	*x = BitSequence{}
	mi := &file_google_firestore_v1_bloom_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BitSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BitSequence) ProtoMessage() {}

func (x *BitSequence) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1_bloom_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BitSequence.ProtoReflect.Descriptor instead.
func (*BitSequence) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1_bloom_filter_proto_rawDescGZIP(), []int{0}
}

func (x *BitSequence) GetBitmap() []byte {
	if x != nil {
		return x.Bitmap
	}
	return nil
}

func (x *BitSequence) GetPadding() int32 {
	if x != nil {
		return x.Padding
	}
	return 0
}

// A bloom filter (https://en.wikipedia.org/wiki/Bloom_filter).
//
// The bloom filter hashes the entries with MD5 and treats the resulting 128-bit
// hash as 2 distinct 64-bit hash values, interpreted as unsigned integers
// using 2's complement encoding.
//
// These two hash values, named `h1` and `h2`, are then used to compute the
// `hash_count` hash values using the formula, starting at `i=0`:
//
//	h(i) = h1 + (i * h2)
//
// These resulting values are then taken modulo the number of bits in the bloom
// filter to get the bits of the bloom filter to test for the given entry.
type BloomFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bloom filter data.
	Bits *BitSequence `protobuf:"bytes,1,opt,name=bits,proto3" json:"bits,omitempty"`
	// The number of hashes used by the algorithm.
	HashCount int32 `protobuf:"varint,2,opt,name=hash_count,json=hashCount,proto3" json:"hash_count,omitempty"`
}

func (x *BloomFilter) Reset() {
	*x = BloomFilter{}
	mi := &file_google_firestore_v1_bloom_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BloomFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomFilter) ProtoMessage() {}

func (x *BloomFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1_bloom_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomFilter.ProtoReflect.Descriptor instead.
func (*BloomFilter) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1_bloom_filter_proto_rawDescGZIP(), []int{1}
}

func (x *BloomFilter) GetBits() *BitSequence {
	if x != nil {
		return x.Bits
	}
	return nil
}

func (x *BloomFilter) GetHashCount() int32 {
	if x != nil {
		return x.HashCount
	}
	return 0
}

var File_google_firestore_v1_bloom_filter_proto protoreflect.FileDescriptor

var file_google_firestore_v1_bloom_filter_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x3f, 0x0a,
	0x0b, 0x42, 0x69, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x69,
	0x74, 0x6d, 0x61, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x62,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x62,
	0x69, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0xc8, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10,
	0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x3b, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0xa2,
	0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_v1_bloom_filter_proto_rawDescOnce sync.Once
	file_google_firestore_v1_bloom_filter_proto_rawDescData = file_google_firestore_v1_bloom_filter_proto_rawDesc
)

func file_google_firestore_v1_bloom_filter_proto_rawDescGZIP() []byte {
	file_google_firestore_v1_bloom_filter_proto_rawDescOnce.Do(func() {
		file_google_firestore_v1_bloom_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_v1_bloom_filter_proto_rawDescData)
	})
	return file_google_firestore_v1_bloom_filter_proto_rawDescData
}

var file_google_firestore_v1_bloom_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_firestore_v1_bloom_filter_proto_goTypes = []any{
	(*BitSequence)(nil), // 0: google.firestore.v1.BitSequence
	(*BloomFilter)(nil), // 1: google.firestore.v1.BloomFilter
}
var file_google_firestore_v1_bloom_filter_proto_depIdxs = []int32{
	0, // 0: google.firestore.v1.BloomFilter.bits:type_name -> google.firestore.v1.BitSequence
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_firestore_v1_bloom_filter_proto_init() }
func file_google_firestore_v1_bloom_filter_proto_init() {
	if File_google_firestore_v1_bloom_filter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_v1_bloom_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_v1_bloom_filter_proto_goTypes,
		DependencyIndexes: file_google_firestore_v1_bloom_filter_proto_depIdxs,
		MessageInfos:      file_google_firestore_v1_bloom_filter_proto_msgTypes,
	}.Build()
	File_google_firestore_v1_bloom_filter_proto = out.File
	file_google_firestore_v1_bloom_filter_proto_rawDesc = nil
	file_google_firestore_v1_bloom_filter_proto_goTypes = nil
	file_google_firestore_v1_bloom_filter_proto_depIdxs = nil
}
