// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type HexaDisplayInfo struct {
	Type            HexaDisplayType
	EntityId        int64
	UniqueId        int64
	Location        *HexLocation
	Parameter       int64
	StageRewardInfo *StrategyClearRewardInfo
}

func (x *HexaDisplayInfo) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
