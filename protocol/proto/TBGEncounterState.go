// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type TBGEncounterState int32

const (
	TBGEncounterState_None      TBGEncounterState = 0
	TBGEncounterState_Active    TBGEncounterState = 1
	TBGEncounterState_Disposing TBGEncounterState = 2
)

var (
	TBGEncounterState_name = map[int32]string{
		0: "None",
		1: "Active",
		2: "Disposing",
	}
	TBGEncounterState_value = map[string]int32{
		"None":      0,
		"Active":    1,
		"Disposing": 2,
	}
)

func (x TBGEncounterState) String() string {
	return TBGEncounterState_name[int32(x)]
}

func (x TBGEncounterState) Value(sr string) TBGEncounterState {
	return TBGEncounterState(TBGEncounterState_value[sr])
}
