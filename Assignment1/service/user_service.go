package service

import (
	"Training/Assignment1/entity"
	"context"
	"fmt"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser_service(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID_service(ctx context.Context, id int) (entity.User, error)
	UpdateUser_service(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser_service(ctx context.Context, id int) error
	GetAllUsers_service(ctx context.Context) ([]entity.User, error)
	GetUserWithLatestSubmission_service(ctx context.Context, userID int) (entity.UserWithRiskProfile, error)
}

type IUserRepository interface {
	CreateUser_repo(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID_repo(ctx context.Context, id int) (entity.User, error)
	UpdateUser_repo(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser_repo(ctx context.Context, id int) error
	GetAllUsers_repo(ctx context.Context) ([]entity.User, error)
	GetUserWithLatestSubmission_repo(ctx context.Context, userID int) (entity.User, entity.Submission, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateUser_service(ctx context.Context, user *entity.User) (entity.User, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdUser, err := u.userRepo.CreateUser_repo(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}
	return createdUser, nil
}

func (u *userService) GetUserByID_service(ctx context.Context, id int) (entity.User, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	user, err := u.userRepo.GetUserByID_repo(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
	}
	return user, nil
}

func (u *userService) UpdateUser_service(ctx context.Context, id int, user entity.User) (entity.User, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	updatedUser, err := u.userRepo.UpdateUser_repo(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser, nil
}

func (u *userService) DeleteUser_service(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	// err := u.userRepo.DeleteUser_repo(ctx, id)
	// if err != nil {
	// 	return fmt.Errorf("gagal menghapus pengguna: %v", err)
	// }
	// return nil

	err := u.userRepo.DeleteUser_repo(ctx, id)
	if err != nil {
		// Tangani error spesifik pengguna tidak ditemukan
		if err.Error() == fmt.Sprintf("pengguna dengan ID %d tidak ditemukan", id) {
			return fmt.Errorf("tidak ada pengguna dengan ID %d yang ditemukan", id)
		}
		return fmt.Errorf("gagal menghapus pengguna dengan ID %d: %v", id, err)
	}
	return nil
}

func (u *userService) GetAllUsers_service(ctx context.Context) ([]entity.User, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := u.userRepo.GetAllUsers_repo(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}

// userService.go
func (s *userService) GetUserWithLatestSubmission_service(ctx context.Context, userID int) (entity.UserWithRiskProfile, error) {
	user, submission, err := s.userRepo.GetUserWithLatestSubmission_repo(ctx, userID)
	if err != nil {
		return entity.UserWithRiskProfile{}, err
	}

	// Buat objek untuk menyimpan informasi yang akan dikembalikan
	userWithRiskProfile := entity.UserWithRiskProfile{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		RiskScore:    submission.RiskScore,
		RiskCategory: string(submission.RiskCategory),
		//RiskDefinition: submission.RiskDefinition,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userWithRiskProfile, nil
}
