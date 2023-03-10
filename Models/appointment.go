package Models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DateTime    string `json:"date"`
	TimeBlockID uint
	DoctorID    uint    `json:"doctor_id"`
	PatientID   uint    `json:"patient_id"`
	PatientName string  `json:"patient_name" gorm:"-"`
	Treatment   string  `json:"treatment"`
	Price       float64 `json:"price"`
	IsCompleted bool    `json:"is_completed"`
	IsPaid      bool    `json:"is_paid"`
	ToothID     uint    `json:"tooth_id"`
	ToothCode   string  `json:"tooth_code" gorm:"-"`
}
