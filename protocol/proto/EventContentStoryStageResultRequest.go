// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentStoryStageResultRequest struct {
	*RequestPacket

	EventContentId int64
	StageUniqueId  int64
}

func (x *EventContentStoryStageResultRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentStoryStageResultRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
