// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type ClanJoinOption int32

const (
	ClanJoinOption_Free       ClanJoinOption = 0
	ClanJoinOption_Permission ClanJoinOption = 1
	ClanJoinOption_All        ClanJoinOption = 2
)

var (
	ClanJoinOption_name = map[int32]string{
		0: "Free",
		1: "Permission",
		2: "All",
	}
	ClanJoinOption_value = map[string]int32{
		"Free":       0,
		"Permission": 1,
		"All":        2,
	}
)

func (x ClanJoinOption) String() string {
	return ClanJoinOption_name[int32(x)]
}

func (x ClanJoinOption) Value(sr string) ClanJoinOption {
	return ClanJoinOption(ClanJoinOption_value[sr])
}
