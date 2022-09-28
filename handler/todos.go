package handler

import (
	"go-api-mahasiswa/helper"
	"go-api-mahasiswa/todos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todosHandler struct {
	todosService todos.Service
}

func NewTodoHandler(todoswaService todos.Service) *todosHandler {
	return &todosHandler{todoswaService}
}

func (h *todosHandler) AddTodos(c *gin.Context) {
	var input todos.AddTodo

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add Todo  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinTeacher, err := h.todosService.AddTodo(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := todos.FormatTodo(loggedinTeacher)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *todosHandler) FindAllTodo(c *gin.Context) {
	var input todos.AllTodo

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add Todo  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	alltodo, err := h.todosService.FindAllTodo(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed Get All Todo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := todos.FormatterAllTodo(alltodo)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *todosHandler) DeleteMahasiswa(c *gin.Context) {
	value := c.Param("id")
	number, err := strconv.ParseUint(value, 10, 32)
	finalIntNum := int(number)

	allmahasiswa, err := h.todosService.DeleteTodo(finalIntNum)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed Delete Todo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := todos.FormatTodo(allmahasiswa)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *todosHandler) UpdateTodo(c *gin.Context) {
	var input todos.UpdateTodo

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update Todo  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	alltodo, err := h.todosService.UpdateTodo(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed Update Todo", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := todos.FormatTodo(alltodo)

	response := helper.APIResponse("Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
