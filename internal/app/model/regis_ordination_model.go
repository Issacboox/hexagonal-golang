package model

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
	ID        uint    `json:"id" gorm:"primaryKey"`
	FirstName string  `json:"fname" validate:"required"`
	LastName  string  `json:"lname" validate:"required"`
	Birthday  string  `json:"bday" validate:"required"`
	Gender    Gender  `json:"gender" validate:"required"`
	Status    Status  `json:"status" gorm:"default:'waiting'"`
	Comment   *string `json:"comment" gorm:"default:null"`
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
