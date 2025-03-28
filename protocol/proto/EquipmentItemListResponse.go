// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EquipmentItemListResponse struct{
	message ProtoMessage
	ResponsePacket

    EquipmentDBs []*EquipmentDB
}

func (x *EquipmentItemListResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EquipmentItemListResponse) ProtoReflect() Message {
	return x
}

func (x *EquipmentItemListResponse) GetProtocol() int32 {
	return mx.Protocol_Equipment_List
}

func (x *EquipmentItemListResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

