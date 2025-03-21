// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CampaignEnterMainStageResponse struct{
	message ProtoMessage
	ResponsePacket

    SaveDataDB *CampaignMainStageSaveDB
}

func (x *CampaignEnterMainStageResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CampaignEnterMainStageResponse) ProtoReflect() Message {
	return x
}

func (x *CampaignEnterMainStageResponse) GetProtocol() int32 {
	return mx.Protocol_Campaign_EnterMainStage
}

func (x *CampaignEnterMainStageResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

