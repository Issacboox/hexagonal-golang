package model

import "github.com/google/uuid"

type Gender string

const (
	Man            Gender = "Man"
	Woman          Gender = "Woman"
	PreferNotToSay Gender = "Prefer not to say"
	Alternative    Gender = "Alternative gender"
)

type Status string

const (
	Waiting  Status = "waiting"
	Approved Status = "approved"
	Reject   Status = "reject"
	Cancel   Status = "cancel"
)

const (
	Active   Status = "Active"
	Inactive Status = "Inactive"
)

type RegisOrdinary struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"fname" validate:"required"`
	LastName  string    `json:"lname" validate:"required"`
	Birthday  string    `json:"bday" validate:"required"`
	Gender    Gender    `json:"gender" validate:"required"`
	Status    Status    `json:"status" gorm:"default:'waiting'"`
	Comment   *string   `json:"comment" gorm:"default:null"`
}

type Pagination struct {
	Page  int `json:"page"`
	View  int `json:"view"`
	Total int `json:"total"`
}

type UpdateStatusRequest struct {
	Status  string `json:"status"`
	Comment string `json:"comment"`
}
