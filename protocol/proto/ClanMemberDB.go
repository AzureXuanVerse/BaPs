// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "github.com/gucooing/BaPs/pkg/mx"
)

type ClanMemberDB struct{
    AccountId int64
    AccountLevel int64
    AccountNickName string
    ClanDBId int64
    RepresentCharacterUniqueId int64
    RepresentCharacterCostumeId int64
    AttendanceCount int64
    CafeComfortValue int64
    ClanSocialGrade ClanSocialGrade
    JoinDate mx.MxTime
    SocialGradeUpdateTime mx.MxTime
    LastLoginDate mx.MxTime
    GameLoginDate mx.MxTime
    AppliedDate mx.MxTime
    AttachmentDB *AccountAttachmentDB
}

