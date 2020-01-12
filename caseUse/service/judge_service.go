package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type JudgeServiceImpl struct {
	judgeRepo caseUse.JudgeRepository
}

func NewJudgeServiceImpl(opRepo caseUse.JudgeRepository) *JudgeServiceImpl {
	return &JudgeServiceImpl{judgeRepo: opRepo}
}

func (osi *JudgeServiceImpl) CreateJudge(judge *entity.Judge) (*entity.Judge, []error) {
	judge, err1 := osi.judgeRepo.CreateJudge(judge)
	if len(err1) > 0 {
		panic(err1)
	}
	return judge, err1
}
