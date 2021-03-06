// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type OneofDescriptorProto struct {
	desc *descriptorpb.OneofDescriptorProto
}

func NewOneofDescriptorProto(name string) *OneofDescriptorProto {
	return &OneofDescriptorProto{
		desc: &descriptorpb.OneofDescriptorProto{
			Name: proto.String(name),
		},
	}
}

func (md *OneofDescriptorProto) GetName() string {
	return md.desc.GetName()
}

func (md *OneofDescriptorProto) Build() *descriptorpb.OneofDescriptorProto {
	return md.desc
}
