// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameRoadPuzzleTilePlaceResponse struct {
	*ResponsePacket

	RailTileToPlace *RoadPuzzleRailTileData
	SaveDB          *RoadPuzzleBoardSaveDB
	ParcelResultDB  *ParcelResultDB
}

func (x *MiniGameRoadPuzzleTilePlaceResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameRoadPuzzleTilePlaceResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
