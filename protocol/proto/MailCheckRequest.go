// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type MailCheckRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *MailCheckRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *MailCheckRequest) ProtoReflect() Message {
	return x
}

func (x *MailCheckRequest) GetProtocol() int32 {
	return mx.Protocol_Mail_Check
}

func (x *MailCheckRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

