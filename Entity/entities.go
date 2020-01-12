package entity

import (
	"time"

	"github.com/Surafeljava/gorm"
)

type Admin struct {
	ID       uint
	AdminId  string `gorm:"type:varchar(50);not null"`
	AdminPwd string `gorm:"type:varchar(50);not null"`
}

type UserType struct {
	gorm.Model
	UsrId  string
	UsrPwd string
}

type Case struct {
	ID            uint
	CaseNum       string `gorm:"type:varchar(50);not null"`
	CaseTitle     string `gorm:"type:varchar(50);not null"`
	CaseDesc      string `gorm:"type:varchar(50);not null"`
	CaseStatus    string `gorm:"type:varchar(50);not null"`
	CaseType      string `gorm:"type:varchar(50);not null"`
	CaseCreation  time.Time
	CaseCourtDate time.Time
	CaseJudge     string `gorm:"type:varchar(50);not null"`
}

type Relation struct {
	ID      uint
	CaseNum string `gorm:"type:varchar(255);not null"`
	PlId    string `gorm:"type:varchar(255);not null"`
	AcId    string `gorm:"type:varchar(255);not null"`
	JuId    string `gorm:"type:varchar(255);not null"`
}

type Decision struct {
	ID           uint
	CaseNum      string `gorm:"type:varchar(255);not null"`
	DecisionDate time.Time
	Decision     string `gorm:"type:varchar(255);not null"`
	DecisionDesc string `gorm:"type:varchar(255);not null"`
}

type Witness struct {
	ID          uint
	CaseNum     string `gorm:"type:varchar(255);not null"`
	WitnessDoc  string `gorm:"type:varchar(255);not null"`
	WitnessType string `gorm:"type:varchar(255);not null"`
}

type Judge struct {
	ID           uint
	JudgeId      string `gorm:"type:varchar(50);not null"`
	JudgePwd     string `gorm:"type:varchar(50);not null"`
	JudgeName    string `gorm:"type:varchar(50);not null"`
	JudgeGender  string `gorm:"type:varchar(50);not null"`
	JudgeAddress string `gorm:"type:varchar(50);not null"`
	JudgePhone   string `gorm:"type:varchar(50);not null"`
	JudgeType    string `gorm:"type:varchar(50);not null"`
	JudgePhoto   string `gorm:"type:varchar(50);not null"`
}

type Notification struct {
	ID             uint
	NotDescription string `gorm:"type:varchar(255);not null"`
	NotTitle       string `gorm:"type:varchar(255);not null"`
	NotLevel       string `gorm:"type:varchar(50);not null"`
	NotDate        time.Time
}

type Opponent struct {
	ID         uint
	OppId      string `gorm:"type:varchar(50);not null"`
	OppPwd     string `gorm:"type:varchar(50);not null"`
	OppType    string `gorm:"type:varchar(50);not null"`
	OppName    string `gorm:"type:varchar(50);not null"`
	OppGender  string `gorm:"type:varchar(50);not null"`
	OppBD      time.Time
	OppAddress string `gorm:"type:varchar(50);not null"`
	OppPhone   string `gorm:"type:varchar(50);not null"`
	OppPhoto   string `gorm:"type:varchar(50);not null"`
}

type SuccessMessage struct {
	Status  string
	Message string
}