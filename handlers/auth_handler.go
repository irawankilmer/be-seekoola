package handlers

import (
	"be-sakoola/config"
	"be-sakoola/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	//Validasi input dari body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cek email apa sudah terdaftar
	var eUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&eUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email sudah terdaftar!",
		})
		return
	}

	// Hash password
	hPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Bcrypt error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Hash Password gagal!",
		})
		return
	}

	// Ambil data role default
	var userRole models.Role
	config.DB.FirstOrCreate(&userRole, models.Role{
		Name: "admin",
	})

	// Simpan user baru ke database
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hPass),
		Roles: []models.Role{
			userRole,
		},
	}
	config.DB.Create(&user)

	// Respon
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi user baru telah berhasil",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"roles": user.Roles,
		},
	})
}
