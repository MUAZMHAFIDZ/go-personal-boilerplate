// controllers/post_controller.go

package controllers

import (
	"fmt"
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/services"
	"net/http"
	"path/filepath"
	"time"
	"os"

	"github.com/gin-gonic/gin"
)

func uploadImage(c *gin.Context) (string, error) {
	// Ambil file dari form
	file, err := c.FormFile("image")
	if err != nil {
		return "", fmt.Errorf("image is required")
	}

	// Ambil nama file asli dan ekstensi
	ext := filepath.Ext(file.Filename)
	baseName := filepath.Base(file.Filename)
	// Tambahkan timestamp unik di belakang nama file
	timestamp := time.Now().Format("20060102_150405")
	uniqueFileName := fmt.Sprintf("%s_%s%s", baseName[:len(baseName)-len(ext)], timestamp, ext)

	// Tentukan lokasi penyimpanan dengan nama file unik
	imagePath := filepath.Join("uploads", uniqueFileName)

	// Simpan file
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	return imagePath, nil
}


func CreatePostHandler(c *gin.Context) {
	// Paksa parsing multipart form (max 10MB)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Ambil nilai dari form
	title := c.PostForm("title")
	description := c.PostForm("description")

	// Debugging
	fmt.Println("Title:", title)
	fmt.Println("Description:", description)

	if title == "" || description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Description are required"})
		return
	}

	imagePath, err := uploadImage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Buat objek post
	post := models.Post{
		Title:       title,
		Description: description,
		Image:       imagePath,
	}

	// Handle upload gambar jika ada
	// if imagePath, err := uploadImage(c); err == nil {
	// 	post.Image = imagePath
	// }

	// Simpan ke database
	result, err := services.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetPostsHandler(c *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func UpdatePostHandler(c *gin.Context) {
    id := c.Param("id") // Ambil ID dari parameter URL
    // Paksa parsing multipart form (karena kita juga menerima file)
    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
        return
    }
    // Ambil data dari form: title dan description
    title := c.PostForm("title")
    description := c.PostForm("description")
    if title == "" || description == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Description are required"})
        return
    }
    // Ambil post lama dari database untuk mendapatkan path gambar lama
    existingPost, err := services.GetPostByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
        return
    }
    // Jika ada gambar lama, hapus gambar tersebut
    if existingPost.Image != "" {
        err := os.Remove(existingPost.Image) // Hapus gambar lama
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete old image: %v", err)})
            return
        }
    }
    // Ambil gambar baru dengan menggunakan fungsi uploadImage
    var imagePath string
    imagePath, err = uploadImage(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
        return
    }
    // Buat objek Post untuk mengupdate
    post := models.Post{
        ID:          id,           // Pastikan ID tetap sama
        Title:       title,        // Set judul baru
        Description: description,  // Set deskripsi baru
        Image:       imagePath,    // Set gambar baru jika ada
    }
    // Panggil service untuk update post
    updatedPost, err := services.UpdatePost(id, &post)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Kirimkan response dengan post yang telah diupdate
    c.JSON(http.StatusOK, updatedPost)
}


func DeletePostHandler(c *gin.Context) {
	id := c.Param("id") // ID tetap sebagai string

	// Ambil data post berdasarkan ID
	post, err := services.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}
	// Hapus file gambar terkait jika ada
	if post.Image != "" {
		err := os.Remove(post.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete image: %v", err)})
			return
		}
	}
	// Hapus post dari database
	if err := services.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	// Kirim status sukses jika tidak ada error
	c.Status(http.StatusNoContent)
}


