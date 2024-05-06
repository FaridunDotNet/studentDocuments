package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Students struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Cours     string `gorm:"not null"`
	Direction string `gorm:"not null"`
	Group     string `gorm:"not null"`
}

func main() {

	dsn := "host=localhost user=postgres password=qwerty dbname=student_documents port=5432 sslmode=disable TimeZone=Asia/Dushanbe "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to the coonecting %s", err)
	}
	fmt.Println("connecting to the database")
	db.AutoMigrate(&Students{})

	router := gin.Default()

	router.POST("/students", func(c *gin.Context) {
		createStudents(c, db)
	})
	router.GET("/students", func(c *gin.Context) {
		fetchStudents(c, db)
	})
	router.GET("/students/:id", func(c *gin.Context) {
		fetchStudentById(c, db)
	})
	router.PUT("/students/:id", func(c *gin.Context) {
		updateStudents(c, db)
	})
	router.DELETE("/students/:id", func(c *gin.Context) {
		deleteStudents(c, db)
	})

	router.Run(":8080")
}

func createStudents(c *gin.Context, db *gorm.DB) {
	var student Students
	if err := c.BindJSON(&student); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func fetchStudents(c *gin.Context, db *gorm.DB) {
	var student []Students
	db.Find(&student)
	c.JSON(http.StatusOK, student)
}

func fetchStudentById(c *gin.Context, db *gorm.DB) {
	var student Students
	id := c.Param("id")
	db.First(&student, id)
	c.JSON(http.StatusOK, student)
}

func updateStudents(c *gin.Context, db *gorm.DB) {
	var student Students
	id := c.Param("id")
	if err := db.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errot": "Record not found"})
		return
	}
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&student)
	c.JSON(http.StatusOK, student)
}

func deleteStudents(c *gin.Context, db *gorm.DB) {
	var student Students
	id := c.Param("id")
	db.First(&student, id)
	db.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successful"})
}
