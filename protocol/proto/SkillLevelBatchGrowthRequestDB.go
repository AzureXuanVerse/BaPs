// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type SkillLevelBatchGrowthRequestDB struct {
	SkillSlot    string
	Level        int32
	ReplaceInfos []*SelectTicketReplaceInfo
}

func (x *SkillLevelBatchGrowthRequestDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
