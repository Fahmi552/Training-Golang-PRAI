package postgresgorm

import (
	"Training/Assignment1/entity"
	"Training/Assignment1/service"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db GormDBIface
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db GormDBIface) service.IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser_repo(ctx context.Context, user *entity.User) (entity.User, error) {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	return *user, nil
}

func (u *userRepository) GetUserByID_repo(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := u.db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) UpdateUser_repo(ctx context.Context, id int, user entity.User) (entity.User, error) {
	var existingUser entity.User
	if err := u.db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").First(&existingUser, id).Error; err != nil {
		log.Printf("Error finding user to update: %v\n", err)
		return entity.User{}, err
	}

	// Memperbarui informasi pengguna
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	if err := u.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}
	return existingUser, nil
}

func (u *userRepository) DeleteUser_repo(ctx context.Context, id int) error {
	// Pertama, cek apakah pengguna dengan ID tersebut ada
	var user entity.User
	if result := u.db.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Jika tidak ditemukan, kembalikan error spesifik
			return fmt.Errorf("pengguna dengan ID %d tidak ditemukan", id)
		}
		return result.Error
	}

	// Jika ditemukan, lanjutkan untuk menghapus pengguna tersebut
	if err := u.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

func (u *userRepository) GetAllUsers_repo(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}

// userRepository.go
func (u *userRepository) GetUserWithLatestSubmission_repo(ctx context.Context, userID int) (entity.User, entity.Submission, error) {
	var user entity.User
	var submission entity.Submission

	// Ambil user berdasarkan ID
	if err := u.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return user, submission, err
	}

	// Ambil submission terbaru berdasarkan user_id
	if err := u.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&submission).Error; err != nil {
		// Jika tidak ada submission ditemukan, kembalikan nilai kosong untuk submission
		if err == gorm.ErrRecordNotFound {
			return user, entity.Submission{}, nil
		}
		return user, submission, err
	}

	return user, submission, nil
}
