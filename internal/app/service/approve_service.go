package service

import (
	m "bam/internal/app/model"
	r "bam/internal/app/repository"
)

type ApproveService struct {
	repo r.IApproveRepository
}

func NewApproveService(repo r.IApproveRepository) *ApproveService {
	return &ApproveService{repo: repo}
}

func (s *ApproveService) GetUserByID(id uint) (*m.RegisOrdinary, error) {
	return s.repo.FindOrdinationByID(id)
}

func (s *ApproveService) RegisterOrdination(reg *m.RegisOrdinary) error {
	return s.repo.RegisterOrdination(reg)
}

// service/approve_service.go
func (s *ApproveService) FindOrdinationByID(id uint) (*m.RegisOrdinary, error) {
	return s.repo.FindOrdinationByID(id)
}

func (s *ApproveService) UpdateOrdination(reg *m.RegisOrdinary) error {
	return s.repo.UpdateOrdination(reg)

}
func (s *ApproveService) DeleteOrdination(id uint) error {
	return s.repo.DeleteOrdination(id)
}

func (s *ApproveService) FindOrdinationByName(firstName, lastName string) (*m.RegisOrdinary, error) {
	return s.repo.FindOrdinationByName(firstName, lastName)
}

// service/approve_service.go
func (s *ApproveService) FindOrdinations() ([]*m.RegisOrdinary, error) {
    return s.repo.FindOrdinations()
}

func (s *ApproveService) FindOrdinationByStatus(status string) (*m.RegisOrdinary, error) {
	return s.repo.FindOrdinationByStatus(status)
}

func (s *ApproveService) UpdateOrdinationStatus(id uint, status, comment string) error {
	return s.repo.UpdateOrdinationStatus(id, status, comment)
}

