// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventImageRequest struct {
	*RequestPacket

	EventId int64
}

func (x *EventImageRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventImageRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
