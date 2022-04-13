package controllers

import (
	//"encoded/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/common/easy-lease/pkg/models"
	"github.com/common/easy-lease/pkg/utils"
	"github.com/gorilla/mux"
)

// GetSocieties godoc
// @Summary Get details of all societies
// @Description Get details of all societies
// @Tags societies
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.Society
// @Router /societies [get]
func GetSociety(w http.ResponseWriter, r *http.Request) {
	newSocieties := models.GetAllSocieties()
	res, _ := json.Marshal(newSocieties)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSocietyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyId := vars["societyId"]
	ID, err := strconv.ParseInt(societyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	societyDetails, _ := models.GetSocietyById(ID)

	res, _ := json.Marshal(societyDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateSociety(w http.ResponseWriter, r *http.Request) {
	CreateSociety := &models.Society{}
	utils.ParseBody(r, CreateSociety)
	society := CreateSociety.CreateSociety()
	res, _ := json.Marshal(society)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteSociety(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	societyId := vars["societyId"]
	ID, err := strconv.ParseInt(societyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	society := models.DeleteSociety(ID)
	res, _ := json.Marshal(society)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSociety(w http.ResponseWriter, r *http.Request) {
	var updateSociety = &models.Society{}
	utils.ParseBody(r, updateSociety)
	vars := mux.Vars(r)
	societyId := vars["societyId"]
	ID, err := strconv.ParseInt(societyId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	societyDetails, db := models.GetSocietyById(ID)

	var societyName string = updateSociety.SocietyName
	var societyAddress string = updateSociety.SocietyAddress
	var societyCity string = updateSociety.SocietyCity
	var societyAmenities string = updateSociety.SocietyAmenities
	var societyImg string = updateSociety.SocietyImg

	if societyName != "" {
		societyDetails.SocietyName = societyName
	}

	if societyAddress != "" {
		societyDetails.SocietyAddress = societyAddress
	}

	if societyCity != "" {
		societyDetails.SocietyCity = societyCity
	}

	if societyAmenities != "" {
		societyDetails.SocietyAmenities = societyAmenities
	}

	if societyImg != "" {
		societyDetails.SocietyImg = societyImg
	}

	db.Save(&societyDetails)

	res, _ := json.Marshal(societyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
