// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AccountNicknameResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountDB *AccountDB
}

func (x *AccountNicknameResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountNicknameResponse) ProtoReflect() Message {
	return x
}

func (x *AccountNicknameResponse) GetProtocol() int32 {
	return mx.Protocol_Account_Nickname
}

func (x *AccountNicknameResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

