// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameTableBoardSyncResponse struct {
	*ResponsePacket

	SaveDB *TBGBoardSaveDB
}

func (x *MiniGameTableBoardSyncResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameTableBoardSyncResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
