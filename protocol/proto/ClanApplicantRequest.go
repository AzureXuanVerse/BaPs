// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ClanApplicantRequest struct {
	*RequestPacket

	OffSet int64
}

func (x *ClanApplicantRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ClanApplicantRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
