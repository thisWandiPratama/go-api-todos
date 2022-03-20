package handler

import (
	"go-api-mahasiswa/helper"
	"go-api-mahasiswa/mahasiswa"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mahasiswaHandler struct {
	mahasiswaService mahasiswa.Service
}

func NewMahasiswaHandler(mahasiswaService mahasiswa.Service) *mahasiswaHandler {
	return &mahasiswaHandler{mahasiswaService}
}

func (h *mahasiswaHandler) AddMahasiswa(c *gin.Context) {
	var input mahasiswa.AddMahasiswa

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add Mahasiswa  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinTeacher, err := h.mahasiswaService.AddMahasiswa(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := mahasiswa.FormatMahasiswa(loggedinTeacher)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
