// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AccountSetRepresentCharacterAndCommentResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountDB *AccountDB
    RepresentCharacterDB *CharacterDB
}

func (x *AccountSetRepresentCharacterAndCommentResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountSetRepresentCharacterAndCommentResponse) ProtoReflect() Message {
	return x
}

func (x *AccountSetRepresentCharacterAndCommentResponse) GetProtocol() int32 {
	return mx.Protocol_Account_SetRepresentCharacterAndComment
}

func (x *AccountSetRepresentCharacterAndCommentResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

