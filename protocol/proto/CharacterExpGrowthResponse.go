// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type CharacterExpGrowthResponse struct {
	*ResponsePacket

	CharacterDB       *CharacterDB
	AccountCurrencyDB *AccountCurrencyDB
	ConsumeResultDB   *ConsumeResultDB
}

func (x *CharacterExpGrowthResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *CharacterExpGrowthResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
