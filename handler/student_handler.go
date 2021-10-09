package handler

import (
	"net/http"
	"strconv"
	"test-start/helper"
	"test-start/models/students"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentService students.Service
}

func NewUserHandler(studentService students.Service) *studentHandler {
	return &studentHandler{studentService}
}

func (h *studentHandler) StoreStudent(c *gin.Context) {
	var input students.StudentsInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		if input.Age == 0 {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Create Student failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Student failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPost, err := h.studentService.CreateStudent(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create Student failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Create Student Success", http.StatusOK, "success", newPost)
	c.JSON(http.StatusOK, response)

}

func (h *studentHandler) PutStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get Student Error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.studentService.GetStudentById(studentId)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error to get student", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// validate input update
	var input students.StudentsInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		if input.Age == 0 {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Create Student failed", http.StatusUnprocessableEntity, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Student failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// updateStudent

	updateStudent, err := h.studentService.UpdateStudentById(studentId, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Student failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update Student Success", http.StatusOK, "success", updateStudent)
	c.JSON(http.StatusOK, response)

}

func (h *studentHandler) GetSudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get Student Error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	student, err := h.studentService.GetStudentById(studentId)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error to get student", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Get Student", http.StatusOK, "success", student)
	c.JSON(http.StatusOK, response)
}

func (h *studentHandler) DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get Student Error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = h.studentService.DeleteStudentByID(studentId)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error to get student", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete Student", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
