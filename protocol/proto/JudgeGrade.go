// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type JudgeGrade int32

const (
	JudgeGrade_None     JudgeGrade = 0
	JudgeGrade_Miss     JudgeGrade = 1
	JudgeGrade_Attack   JudgeGrade = 2
	JudgeGrade_Critical JudgeGrade = 3
)

var (
	JudgeGrade_name = map[int32]string{
		0: "None",
		1: "Miss",
		2: "Attack",
		3: "Critical",
	}
	JudgeGrade_value = map[string]int32{
		"None":     0,
		"Miss":     1,
		"Attack":   2,
		"Critical": 3,
	}
)

func (x JudgeGrade) String() string {
	return JudgeGrade_name[int32(x)]
}

func (x JudgeGrade) Value(sr string) JudgeGrade {
	return JudgeGrade(JudgeGrade_value[sr])
}
