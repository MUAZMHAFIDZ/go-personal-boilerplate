// services/post_service.go

package services

import (
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/utils"
)

func CreatePost(post *models.Post) (*models.Post, error) {
	db := utils.InitDB()

	err := db.Create(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPosts() ([]models.Post, error) {
	db := utils.InitDB()

	var posts []models.Post
	err := db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func UpdatePost(id string, post *models.Post) (*models.Post, error) {
    db := utils.InitDB()

    var existingPost models.Post
    if err := db.First(&existingPost, "id = ?", id).Error; err != nil {
        return nil, err
    }

    existingPost.Title = post.Title
    existingPost.Description = post.Description
    if post.Image != "" {
        existingPost.Image = post.Image
    }

    err := db.Save(&existingPost).Error
    if err != nil {
        return nil, err
    }

    return &existingPost, nil
}


func DeletePost(id string) error {
	db := utils.InitDB()

	err := db.Delete(&models.Post{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPostByID(id string) (*models.Post, error) {
	db := utils.InitDB()

	var post models.Post
	if err := db.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
