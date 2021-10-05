package grades

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	registerHandlers(r)
}

func registerHandlers(r *gin.Engine) {
	r.GET("/students", func(c *gin.Context) {
		studentsMutex.Lock()
		defer studentsMutex.Unlock()

		c.JSON(http.StatusOK, students)
	})

	r.GET("/students/:id", func(c *gin.Context) {
		studentsMutex.Lock()
		defer studentsMutex.Unlock()

		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		student, err := students.GetByID(int(studentId))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, student)
		}
	})

	r.GET("/students/:id/grades", func(c *gin.Context) {
		studentsMutex.Lock()
		defer studentsMutex.Unlock()

		studentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		student, err := students.GetByID(int(studentId))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, student.Grades)
		}
	})
}
