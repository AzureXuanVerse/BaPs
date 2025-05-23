// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type AccountBanErrorPacket struct {
	*ResponsePacket

	ErrorCode WebAPIErrorCode
	BanReason string
}

func (x *AccountBanErrorPacket) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *AccountBanErrorPacket) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
