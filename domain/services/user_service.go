package services

import (
	"github.com/alirezazeynali75/simple-to-do-app/domain"
	"github.com/alirezazeynali75/simple-to-do-app/domain/ports"
)

type UserService struct {
	userRepo ports.UserRepository
	jwt *domain.JwtManager
}