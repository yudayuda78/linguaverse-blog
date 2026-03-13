package repository

import (
	"fmt"
	"github.com/yudayuda78/linguaverse/internal/database"
	"github.com/yudayuda78/linguaverse/internal/model"
)

func GetAllPosts() ([]model.Post, error) {
	fmt.Println("Repository: GetAllPosts called")

	rows, err := database.DB.Query("SELECT id, title, slug, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {
		var post model.Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}