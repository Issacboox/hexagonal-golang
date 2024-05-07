package repository

import (
	m "bam/internal/app/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type IApproveRepository interface {
	RegisterOrdination(reg *m.RegisOrdinary) error
	FindOrdinationByID(id uint) (*m.RegisOrdinary, error)
	UpdateOrdination(reg *m.RegisOrdinary) error
	DeleteOrdination(id uint) error
	FindOrdinationByName(name string) ([]*m.RegisOrdinary, error)
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

func (r *ApproveRepository) RegisterOrdination(reg *m.RegisOrdinary) error {
	// Format the birthday to DD/MM/YYYY
	birthday, err := time.Parse("02/01/2006", reg.Birthday)
	if err != nil {
		return err
	}
	reg.Birthday = birthday.Format("02/01/2006")

	// Check if the gender is valid
	switch reg.Gender {
	case m.Man, m.Woman, m.PreferNotToSay, m.Alternative:
		// Valid gender, proceed
	default:
		return fmt.Errorf("invalid gender: %s", reg.Gender)
	}

	return r.db.Create(reg).Error
}

func (r *ApproveRepository) FindOrdinationByID(id uint) (*m.RegisOrdinary, error) {
	var reg m.RegisOrdinary
	result := r.db.First(&reg, id)
	return &reg, result.Error
}

func (r *ApproveRepository) UpdateOrdination(reg *m.RegisOrdinary) error {
	// Validate gender
	if reg.Gender != m.Man && reg.Gender != m.Woman && reg.Gender != m.PreferNotToSay && reg.Gender != m.Alternative {
		return errors.New("invalid gender")
	}

	// Validate birthday format
	_, err := time.Parse("02/01/2006", reg.Birthday)
	if err != nil {
		return errors.New("invalid birthday format, should be DD/MM/YYYY")
	}

	return r.db.Save(reg).Error
}

func (r *ApproveRepository) DeleteOrdination(id uint) error {
	return r.db.Delete(&m.RegisOrdinary{}, id).Error
}

func (r *ApproveRepository) FindOrdinationByName(name string) ([]*m.RegisOrdinary, error) {
	var regs []*m.RegisOrdinary
	result := r.db.Where("first_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&regs)
	if result.Error != nil {
		return nil, result.Error
	}
	return regs, nil
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
