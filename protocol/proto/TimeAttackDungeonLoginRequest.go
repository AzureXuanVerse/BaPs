// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type TimeAttackDungeonLoginRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *TimeAttackDungeonLoginRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *TimeAttackDungeonLoginRequest) ProtoReflect() Message {
	return x
}

func (x *TimeAttackDungeonLoginRequest) GetProtocol() int32 {
	return mx.Protocol_TimeAttackDungeon_Login
}

func (x *TimeAttackDungeonLoginRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

