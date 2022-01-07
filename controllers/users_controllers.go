package controllers

import (
	"net/http"
	"strconv"

	"github.com/chpushpa/student/domain/httperrors"
	"github.com/chpushpa/student/domain/students"
	"github.com/chpushpa/student/services"
	"github.com/gin-gonic/gin"
)

var (
	StudentsController = studentsController{}
)

type studentsController struct{}

func respond(c *gin.Context, isXml bool, httpCode int, body interface{}) {
	if isXml {
		c.XML(httpCode, body)
		return
	}
	c.JSON(httpCode, body)
}

func (controller studentsController) Create(c *gin.Context) {
	var student students.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid json body")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	createdStudent, err := services.StudentService.Create(student)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	// return created user
	c.JSON(http.StatusCreated, createdStudent)
}

func (controller studentsController) Get(c *gin.Context) {
	isXml := c.GetHeader("Accept") == "application/xml"

	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("invalid user id")
		respond(c, isXml, httpErr.Code, httpErr)
		return
	}

	user, getErr := services.StudentService.Get(userId)
	if getErr != nil {
		respond(c, isXml, getErr.Code, getErr)
		return
	}
	respond(c, isXml, http.StatusOK, user)
}
