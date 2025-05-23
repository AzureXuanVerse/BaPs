// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type CafeInteractWithCharacterResponse struct {
	*ResponsePacket

	CafeDB         *CafeDB
	CharacterDB    *CharacterDB
	ParcelResultDB *ParcelResultDB
}

func (x *CafeInteractWithCharacterResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *CafeInteractWithCharacterResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
