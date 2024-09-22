package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Major string `json:"major"`
}

var students = []Student{
	{ID: 1, Name: "John Doe", Age: 20, Major: "Computer Science"},
	{ID: 2, Name: "Jane Smith", Age: 22, Major: "Mathematics"},
}

func getStudentByID(id int) *Student {
	for i, s := range students {
		if s.ID == id {
			return &students[i]
		}
	}
	return nil
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	router.GET("/get-students", func(c *gin.Context) {
		c.JSON(http.StatusOK, students)
	})

	router.GET("/get-student-detail/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := getStudentByID(id)
		if student == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		c.JSON(http.StatusOK, student)
	})

	router.POST("/add-student", func(c *gin.Context) {
		var newStudent Student
		if err := c.ShouldBindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newStudent.ID = len(students) + 1
		students = append(students, newStudent)
		c.JSON(http.StatusCreated, newStudent)
	})

	router.PUT("/update-student/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := getStudentByID(id)
		if student == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, student)
	})

	router.DELETE("/delete-student/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		for i, s := range students {
			if s.ID == id {
				students = append(students[:i], students[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
	})

	router.Run(":8080") // Start the server on port 8080
}
