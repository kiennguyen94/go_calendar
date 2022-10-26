package routs

import (
	"context"
	"errors"
	"fmt"
	models "kiennguyen94/go_calendar/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

var ctx context.Context
var db *bun.DB

func get_doctor(id int64) ([]models.Doctor, error) {
	var doctors []models.Doctor
	if id == 0 {
		err := db.NewSelect().Model(&doctors).OrderExpr("id ASC").Scan(ctx)
		if err != nil {
			msg := fmt.Sprintf("Could not fetch list of doctors from db [%s]", err.Error())
			// g.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return doctors, errors.New(msg)
		}
		return doctors, nil
	} else {
		err := db.NewSelect().Model(&doctors).Where("id = ?", id).Scan(ctx)
		if err != nil {
			msg := fmt.Sprintf("Could not fetch doctor id [%d] from db [%s]", id, err.Error())
			return doctors, errors.New(msg)
		}
		return doctors, nil
	}
}

// @BasePath /api/v1

// Get doctors godoc
// @Summary Get Doctors
// @Schemes
// @Description get doctors
// @Tags doctors
// @Accept json
// @Produce json
// @Success 200 {string} YYY
// @Router /doctor [get]
func GetDoctor(g *gin.Context) {
	// var doctor models.DoctorReq
	var doctor models.Doctor
	if err := g.ShouldBind(&doctor); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("doctor req [%v]", doctor)
	id := doctor.ID
	doctors, err := get_doctor(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("YYY fetched [%v]\n", doctors)
	// var doctor_req []models.DoctorReq
	// var doctor_req []models.DoctorReq
	// for _, e := range doctors {
	// 	doctor_req = append(doctor_req, models.Doctor_to_DoctorReq(&e))
	// }
	fmt.Printf("YYY fetched req [%v]\n", doctors)
	g.JSON(http.StatusOK, doctors)
	return
}

// @BasePath /api/v1

// Post doctors godoc
// @Summary Post Doctors
// @Schemes
// @Description post doctors
// @Tags doctors
// @Accept json
// @Produce json
// @Success 200 {string} XXX
// @Router /doctor [post]
func PostDoctor(g *gin.Context) {
	// var doctor models.DoctorReq
	var doctor models.Doctor
	if err := g.ShouldBindJSON(&doctor); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("XXX get request data %v\n", doctor)
	if doctor.FirstName == "" || doctor.LastName == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Post request needs first_name and last_name"})
		return
	}
	// doctor_mdl := models.DoctorReq_to_Doctor(&doctor)
	db.NewInsert().Model(&doctor).Exec(ctx)
	g.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func PostAppointment(g *gin.Context) {
	var app models.AppointmentReq
	if err := g.ShouldBindJSON(&app); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get doctor by id

}
