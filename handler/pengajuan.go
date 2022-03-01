package handler

import (
	"fmt"
	"go-api-koperasi/helper"
	"go-api-koperasi/pengajuan"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pengajuanHandler struct {
	pengajuanService pengajuan.Service
}

func NewPengajuanHandler(pengajuanService pengajuan.Service) *pengajuanHandler {
	return &pengajuanHandler{pengajuanService}
}

func (h *pengajuanHandler) FindAll(c *gin.Context) {

	loggedinUser, err := h.pengajuanService.FindAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Data failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Get Data failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var bukti []pengajuan.FormatterBuktiJaminan

	formatter := pengajuan.FormatAlls(loggedinUser, bukti)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) Jaminan(c *gin.Context) {
	var input pengajuan.UpdateJaminanPengajuan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add Jamainan  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinTeacher, err := h.pengajuanService.UpdateJaminanPengajuan(input)
	buktijaminan, err := h.pengajuanService.FindBuktiAll(input.NoPengajuan)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Update Bantuan failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	buktiformat := pengajuan.FormatBuktiAlls(buktijaminan)

	formatter := pengajuan.FormatUser(loggedinTeacher, buktiformat)

	response := helper.APIResponse("Successfuly Update Bantuan", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) FindByID(c *gin.Context) {
	value := c.Param("id")
	number, err := strconv.ParseUint(value, 10, 32)
	finalIntNum := int(number)

	loggedinTeacher, err := h.pengajuanService.FindByID(finalIntNum)
	buktijaminan, err := h.pengajuanService.FindBuktiAll(finalIntNum)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get By Id failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	buktiformat := pengajuan.FormatBuktiAlls(buktijaminan)

	formatter := pengajuan.FormatUser(loggedinTeacher, buktiformat)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) AddBuktiJaminan(c *gin.Context) {
	file, err := c.FormFile("avatar_bukti")
	if err != nil {
		fmt.Println(err)
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(" test Failed to upload bukti avatar ", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)

	value := c.PostForm("id_pengajuan")
	number, err := strconv.ParseUint(value, 10, 32)
	id_pengajuan := int(number)
	input := pengajuan.AddBuktiJaminan{
		IdPengajuan: id_pengajuan,
		Bukti:       file.Filename,
	}

	buktijaminan, err := h.pengajuanService.AddBuktiJaminan(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get By Id failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			data := gin.H{"is_uploaded": false}
			response := helper.APIResponse("Failed to upload bukti avatar", http.StatusBadRequest, "error", data)

			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	buktiformat := pengajuan.FormatBukti(buktijaminan)
	response := helper.APIResponse("Successfuly", http.StatusOK, "success", buktiformat)
	c.JSON(http.StatusOK, response)

}
