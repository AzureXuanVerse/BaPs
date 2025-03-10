// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type StickerLoginResponse struct{
	message ProtoMessage
	ResponsePacket

    StickerBookDB *StickerBookDB
}

func (x *StickerLoginResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *StickerLoginResponse) ProtoReflect() Message {
	return x
}

func (x *StickerLoginResponse) GetProtocol() int32 {
	return mx.Protocol_Sticker_Login
}

func (x *StickerLoginResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

