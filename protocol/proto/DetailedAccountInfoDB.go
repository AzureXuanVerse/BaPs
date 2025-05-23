// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type DetailedAccountInfoDB struct {
	AccountId                      int64
	Nickname                       string
	Level                          int64
	ClanName                       string
	Comment                        string
	FriendCount                    int64
	FriendCode                     string
	RepresentCharacterUniqueId     int64
	CharacterCount                 int64
	LastNormalCampaignClearStageId int64
	LastHardCampaignClearStageId   int64
	ArenaRanking                   int64
	RaidRanking                    int64
	RaidTier                       int32
	EliminateRaidRanking           int64
	EliminateRaidTier              int32
	AssistCharacterDBs             []*AssistCharacterDB
}

func (x *DetailedAccountInfoDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
