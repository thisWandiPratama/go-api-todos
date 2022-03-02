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
	var jaminanbarang pengajuan.FormaterJaminanBarang

	formatter := pengajuan.FormatAlls(loggedinUser, jaminanbarang, bukti)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) FindPersetujuanAll(c *gin.Context) {

	loggedinUser, err := h.pengajuanService.FindPersetujuanAll()
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
	var jaminanbarang pengajuan.FormaterJaminanBarang

	formatter := pengajuan.FormatAlls(loggedinUser, jaminanbarang, bukti)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) SearchJamianAll(c *gin.Context) {
	var input pengajuan.SearchJaminan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Search Data", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.pengajuanService.SearchJamianAll(input)
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
	var jaminanbarang pengajuan.FormaterJaminanBarang

	formatter := pengajuan.FormatAlls(loggedinUser, jaminanbarang, bukti)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) AddJaminan(c *gin.Context) {
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
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		petugas, err := h.pengajuanService.FindByID(loggedinTeacher.IdPengajuan)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var notifikasi pengajuan.AddNotifikasi
		notifikasi.IdPengajuan = loggedinTeacher.IdPengajuan
		notifikasi.Petugas = petugas.Petugas
		_, err = h.pengajuanService.AddNotifikasi(notifikasi)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	formatter := pengajuan.FormatJaminanBarang(loggedinTeacher)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) AddJaminanTanah(c *gin.Context) {
	var input pengajuan.AddJaminanTanah

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add Jamainan  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinTeacher, err := h.pengajuanService.AddJaminanTanah(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	} else {
		petugas, err := h.pengajuanService.FindByID(loggedinTeacher.IdPengajuan)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		var notifikasi pengajuan.AddNotifikasi
		notifikasi.IdPengajuan = loggedinTeacher.IdPengajuan
		notifikasi.Petugas = petugas.Petugas
		_, err = h.pengajuanService.AddNotifikasi(notifikasi)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	formatter := pengajuan.FormatJaminanTanah(loggedinTeacher)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

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

	if loggedinTeacher.Jaminan == "Tanah" {
		jaminantanah, err := h.pengajuanService.FindJaminanTanahByID(finalIntNum)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Get By Id failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		jaminanbarangs := pengajuan.FormaterJaminanTanah(jaminantanah)
		fmt.Println(jaminanbarangs)
		formatter := pengajuan.FormatUserJaminanTanah(loggedinTeacher, jaminanbarangs, buktiformat)
		response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)

	} else {
		jaminanbarang, err := h.pengajuanService.FindJaminanBarangByID(finalIntNum)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}

			response := helper.APIResponse("Get By Id failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		jaminanbarangs := pengajuan.FormaterJaminanBarang(jaminanbarang)
		formatter := pengajuan.FormatUser(loggedinTeacher, jaminanbarangs, buktiformat)
		response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

		c.JSON(http.StatusOK, response)
	}
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

	value1 := c.PostForm("status_draf")
	number1, err := strconv.ParseUint(value1, 10, 32)
	status_draf := int(number1)

	input := pengajuan.AddBuktiJaminan{
		IdPengajuan: id_pengajuan,
		Bukti:       file.Filename,
		StatusDraf:  status_draf,
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

func (h *pengajuanHandler) FindAllByStatusDraf(c *gin.Context) {

	loggedinUser, err := h.pengajuanService.FindAllByStatusDraf()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Get Draf failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Get Draf failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var jaminanbarang pengajuan.FormaterJaminanBarang
	var bukti []pengajuan.FormatterBuktiJaminan
	formatter := pengajuan.FormatAlls(loggedinUser, jaminanbarang, bukti)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *pengajuanHandler) FindNotifikasiAll(c *gin.Context) {

	loggedinUser, err := h.pengajuanService.FindNotifikasiAll()
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

	formatter := pengajuan.FormatNotifikasiAlls(loggedinUser)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
