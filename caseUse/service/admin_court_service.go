package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type AdminCourtServiceImpl struct {
	courtRepo caseUse.CourtRepository
}

func NewAdminCourtServiceImpl(csRepo caseUse.CourtRepository) *AdminCourtServiceImpl {
	return &AdminCourtServiceImpl{courtRepo: csRepo}
}

func (acs *AdminCourtServiceImpl) CreateCourt(court *entity.Court) (*entity.Court, []error) {
	return nil, nil
}
func (acs *AdminCourtServiceImpl) UpdateCourt(court *entity.Court) (*entity.Court, []error) {
	return nil, nil
}

// DeleteCourt(id int) error
func (acs *AdminCourtServiceImpl) CreateAdmin(admin *entity.Admin) (*entity.Admin, []error) {
	return nil, nil
}