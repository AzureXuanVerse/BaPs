// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type CheatCharacterCustomPreset struct {
	UniqueId            int64
	StarGrade           int32
	Level               int32
	ExSkillLevel        int32
	PublicSkillLevel    int32
	PassiveSkillLevel   int32
	ExPassiveSkillLevel int32
	Equipments          []*CheatEquipmentCustomPreset
	Weapon              *CheatWeaponCustomPreset
}

func (x *CheatCharacterCustomPreset) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
