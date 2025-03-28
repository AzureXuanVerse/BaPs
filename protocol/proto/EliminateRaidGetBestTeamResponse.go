// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EliminateRaidGetBestTeamResponse struct{
	message ProtoMessage
	ResponsePacket

	RaidTeamSettingDBsDict map[string][]*RaidTeamSettingDB
}

func (x *EliminateRaidGetBestTeamResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EliminateRaidGetBestTeamResponse) ProtoReflect() Message {
	return x
}

func (x *EliminateRaidGetBestTeamResponse) GetProtocol() int32 {
	return mx.Protocol_EliminateRaid_GetBestTeam
}

func (x *EliminateRaidGetBestTeamResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

