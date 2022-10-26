package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Kind int

const (
	NewPatient Kind = iota
	FollowUp
)

type Doctor struct {
	bun.BaseModel `bun:"table:doctor"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id" form:"id"`
	FirstName string `bun:"first_name" json:"first_name"`
	LastName  string `bun:"last_name" json:"last_name"`
}

type Appointment struct {
	bun.BaseModel `bun:"table:appointment"`

	ID               int64     `bun:"id,pk,autoincrement"`
	PatientFirstname string    `bun:"p_first_name"`
	PatientLastName  string    `bun:"p_last_name"`
	Date             time.Time `bun:",nullzero,notnull"`
	Kind             Kind      `bun:"kind"`
	DoctorId         int64     `bun:"doctor_id"`
}

// type DoctorReq struct {
// 	ID        int64  `json:"id" form:"id"`
// 	FirstName string `json:"first_name"`
// 	Lastname  string `json:"last_name"`
// }

type AppointmentReq struct {
	ID               int64     `json:"id" binding:"required"`
	PatientFirstname string    `json:"patient_first_name"`
	PatientLastName  string    `json:"patient_last_name"`
	Date             time.Time `json:"date_time"`
	Kind             Kind      `json:"kind"`
}

// func Doctor_to_DoctorReq(obj *Doctor) DoctorReq {
// 	return DoctorReq{obj.ID, obj.FirstName, obj.LastName}
// }
//
// func DoctorReq_to_Doctor(obj *DoctorReq) Doctor {
// 	return Doctor{FirstName: obj.FirstName, LastName: obj.Lastname}
// }

func Appt_to_ApptReq(obj *Appointment) AppointmentReq {
	return AppointmentReq{obj.ID, obj.PatientFirstname, obj.PatientLastName, obj.Date, obj.Kind}
}

func ApptReq_to_Appt(obj *AppointmentReq) Appointment {
	return Appointment{PatientFirstname: obj.PatientFirstname, PatientLastName: obj.PatientLastName,
		Date: obj.Date, Kind: obj.Kind}
}
