// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EliminateRaidLoginRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *EliminateRaidLoginRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EliminateRaidLoginRequest) ProtoReflect() Message {
	return x
}

func (x *EliminateRaidLoginRequest) GetProtocol() int32 {
	return mx.Protocol_EliminateRaid_Login
}

func (x *EliminateRaidLoginRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

