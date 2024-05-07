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

func (s *ApproveService) UpdateOrdination(user *m.RegisOrdinary) error {
	// Find the existing user
	existingUser, err := s.repo.FindOrdinationByID(user.ID)
	if err != nil {
		return err
	}

	// Update only the fields that are non-empty
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}
	if user.Birthday != "" {
		existingUser.Birthday = user.Birthday
	}
	if user.Gender != "" {
		existingUser.Gender = user.Gender
	}
	if user.Status != "" {
		existingUser.Status = user.Status
	}
	if user.Comment != nil {
		existingUser.Comment = user.Comment
	}

	// Save the updated user
	return s.repo.UpdateOrdination(existingUser)
}

func (s *ApproveService) DeleteOrdination(id uint) error {
	return s.repo.DeleteOrdination(id)
}

// service
func (s *ApproveService) FindOrdinationByName(name string) ([]*m.RegisOrdinary, error) {
    return s.repo.FindOrdinationByName(name)
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
