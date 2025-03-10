// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeRankUpRequest struct{
	message ProtoMessage
	RequestPacket

    AccountServerId int64
    CafeDBId int64
    ConsumeRequestDB *ConsumeRequestDB
}

func (x *CafeRankUpRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeRankUpRequest) ProtoReflect() Message {
	return x
}

func (x *CafeRankUpRequest) GetProtocol() int32 {
	return mx.Protocol_Cafe_RankUp
}

func (x *CafeRankUpRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

