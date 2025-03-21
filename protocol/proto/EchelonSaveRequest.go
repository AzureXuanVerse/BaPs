// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EchelonSaveRequest struct{
	message ProtoMessage
	RequestPacket

    EchelonDB *EchelonDB
    AssistUseInfos []*ClanAssistUseInfo
    IsPractice bool
}

func (x *EchelonSaveRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EchelonSaveRequest) ProtoReflect() Message {
	return x
}

func (x *EchelonSaveRequest) GetProtocol() int32 {
	return mx.Protocol_Echelon_Save
}

func (x *EchelonSaveRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

