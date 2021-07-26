package Controllers
import (
	"fmt"
	"github.com/Exercise2/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetUsers ... Get all students
func GetUsers(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllUsers(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//CreateUser ... Create student
func CreateUser(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateUser(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//GetUserByID ... Get student by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//UpdateUser ... Update the student information
func UpdateUser(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	error_ := c.BindJSON(&student)
	if error_ != nil {
		fmt.Println(error_.Error())
	}
	err = Models.UpdateUser(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//DeleteUser ... Delete the student record
func DeleteUser(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
