package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Alamat string `json:"alamat"`
}

var posts = []Post{
	{ID: 1, Name: "Koswara", Alamat: "Samarang"},
	{ID: 2, Name: "Jajang", Alamat: "Bayongbong"},
}

var nextID = 3

// GET /posts Ambil semua data
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

// POST /posts tambah data baru
func CreatePost(c *gin.Context) {
	var newPost Post

	if err := c.ShouldBindJSON(&newPost); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

	newPost.ID = nextID
	nextID++
	posts = append(posts, newPost)

	c.JSON(http.StatusCreated, newPost)
}
