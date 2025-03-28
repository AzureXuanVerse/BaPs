// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeRemoveFurnitureResponse struct{
	message ProtoMessage
	ResponsePacket

    CafeDB *CafeDB
    FurnitureDBs []*FurnitureDB
}

func (x *CafeRemoveFurnitureResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeRemoveFurnitureResponse) ProtoReflect() Message {
	return x
}

func (x *CafeRemoveFurnitureResponse) GetProtocol() int32 {
	return mx.Protocol_Cafe_Remove
}

func (x *CafeRemoveFurnitureResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

