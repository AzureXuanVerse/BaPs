// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CampaignMainStageStrategySkipResultResponse struct{
	message ProtoMessage
	ResponsePacket

    TacticRank int64
    CampaignStageHistoryDB *CampaignStageHistoryDB
    LevelUpCharacterDBs []*CharacterDB
    ParcelResultDB *ParcelResultDB
    FirstClearReward []*ParcelInfo
    ThreeStarReward []*ParcelInfo
}

func (x *CampaignMainStageStrategySkipResultResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CampaignMainStageStrategySkipResultResponse) ProtoReflect() Message {
	return x
}

func (x *CampaignMainStageStrategySkipResultResponse) GetProtocol() int32 {
	return mx.Protocol_Campaign_MainStageStrategySkipResult
}

func (x *CampaignMainStageStrategySkipResultResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

