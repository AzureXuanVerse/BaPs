// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ContentSweepMultiSweepPresetListRequest struct {
	*RequestPacket
}

func (x *ContentSweepMultiSweepPresetListRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ContentSweepMultiSweepPresetListRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
