// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameDreamMakerNewGameResponse struct {
	*ResponsePacket

	InfoDB       *MiniGameDreamMakerInfoDB
	ParameterDBs []*MiniGameDreamMakerParameterDB
}

func (x *MiniGameDreamMakerNewGameResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameDreamMakerNewGameResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
