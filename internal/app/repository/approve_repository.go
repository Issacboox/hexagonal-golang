package repository

import (
	m "bam/internal/app/model"

	"gorm.io/gorm"
)

type IApproveRepository interface {
	RegisterOrdination(reg *m.RegisOrdinary) error
	FindOrdinationByID(id uint) (*m.RegisOrdinary, error)
	UpdateOrdination(reg *m.RegisOrdinary) error
	DeleteOrdination(id uint) error
	FindOrdinationByName(firstName, lastName string) (*m.RegisOrdinary, error)
	FindOrdinations() ([]*m.RegisOrdinary, error)
	FindOrdinationByStatus(status string) (*m.RegisOrdinary, error)
	UpdateOrdinationStatus(id uint, status, comment string) error
}

type ApproveRepository struct {
	db             *gorm.DB
	ApproveActions IApproveRepository
}

func NewApproveRepository(db *gorm.DB) *ApproveRepository {
	return &ApproveRepository{db: db}
}

// repository/approve_repository.go
func (r *ApproveRepository) RegisterOrdination(reg *m.RegisOrdinary) error {
    return r.db.Create(reg).Error
}

func (r *ApproveRepository) FindOrdinationByID(id uint) (*m.RegisOrdinary, error) {
	var reg m.RegisOrdinary
	result := r.db.First(&reg, id)
	return &reg, result.Error
}

func (r *ApproveRepository) UpdateOrdination(reg *m.RegisOrdinary) error {
	return r.db.Save(reg).Error
}

func (r *ApproveRepository) DeleteOrdination(id uint) error {
	return r.db.Delete(&m.RegisOrdinary{}, id).Error
}

func (r *ApproveRepository) FindOrdinationByName(firstName, lastName string) (*m.RegisOrdinary, error) {
	var reg m.RegisOrdinary
	result := r.db.Where("fname = ? OR lname = ?", firstName, lastName).First(&reg)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reg, nil
}

func (r *ApproveRepository) FindOrdinations() ([]*m.RegisOrdinary, error) {
	var regs []*m.RegisOrdinary
	result := r.db.Find(&regs)
	return regs, result.Error
}

func (r *ApproveRepository) FindOrdinationByStatus(status string) (*m.RegisOrdinary, error) {
	var reg m.RegisOrdinary
	result := r.db.Where("status = ?", status).First(&reg)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reg, nil
}

func (r *ApproveRepository) UpdateOrdinationStatus(id uint, status, comment string) error {
	return r.db.Model(&m.RegisOrdinary{}).Where("id = ?", id).Updates(map[string]interface{}{"status": status, "comment": comment}).Error
}
