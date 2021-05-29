package user

import (
	"artics-api/src/internal/domain/models"
)

// User entity
type User struct {
	models.User
	Password       string
	FollowingCount int
	FollowerCount  int
}
