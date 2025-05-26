package handlers

import (
	"be-sakoola/config"
	"be-sakoola/internal/dto"
	"be-sakoola/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("sangat-rahasia")

func Register(c *gin.Context) {
	var req dto.RegisterRequest

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

func Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cari data user berdasarkan Email
	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau Password salah!",
		})
		return

	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau Password salah!",
		})
		return
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   user.Roles,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 jam
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal meng-generate token!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// Engke dihapus deui
func Profile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
