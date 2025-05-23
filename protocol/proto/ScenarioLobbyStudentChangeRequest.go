// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ScenarioLobbyStudentChangeRequest struct {
	*RequestPacket

	LobbyStudents       []int64
	LobbyStudentsBefore []int64
}

func (x *ScenarioLobbyStudentChangeRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ScenarioLobbyStudentChangeRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
