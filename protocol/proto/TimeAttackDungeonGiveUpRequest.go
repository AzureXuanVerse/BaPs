// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type TimeAttackDungeonGiveUpRequest struct{
	message ProtoMessage
	RequestPacket

    RoomId int64
}

func (x *TimeAttackDungeonGiveUpRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *TimeAttackDungeonGiveUpRequest) ProtoReflect() Message {
	return x
}

func (x *TimeAttackDungeonGiveUpRequest) GetProtocol() int32 {
	return mx.Protocol_TimeAttackDungeon_GiveUp
}

func (x *TimeAttackDungeonGiveUpRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

