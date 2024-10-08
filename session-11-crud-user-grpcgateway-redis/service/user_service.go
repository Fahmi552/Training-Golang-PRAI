package service

import (
	"Training/session-11-crud-user-grpcgateway-redis/entity"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const redisUserByIDKey = "user:%d"

//New USer service membuat instance baru dari user service

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
	rdb      *redis.Client
}

// NewUserService membuat instance baru dari userService
// func NewUserService(userRepo IUserRepository) IUserService {
// 	return &userService{userRepo: userRepo}
// }

// yang baru
func NewUserService(userRepo IUserRepository, rdb_ *redis.Client) IUserService {
	return &userService{userRepo: userRepo, rdb: rdb_}
}

// CreateUser membuat pengguna baru
func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}

	// insert ke redis
	rediskey := fmt.Sprintf(redisUserByIDKey, createdUser.ID)
	createdUserJSON, err := json.Marshal(createdUser)
	if err != nil {
		log.Println("gagal marshal json")
	}
	if err := s.rdb.Set(ctx, rediskey, createdUserJSON, 100000*time.Second).Err(); err != nil {
		log.Println("error when set redis key", rediskey)
	}

	return createdUser, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	// user, err := s.userRepo.GetUserByID(ctx, id)
	// if err != nil {
	// 	return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
	// }

	// return user, nil

	var user entity.User
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	val, err := s.rdb.Get(ctx, redisKey).Result()
	log.Println(redisKey)
	if err == nil {
		log.Println("data tersedia di redis")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Println("gagal marshall data di redis, coba ambil ke database")
		}
	}
	if err != nil {
		log.Println("data tidak tersedia di redis, ambil dari database")
		user, err = s.userRepo.GetUserByID(ctx, id)
		if err != nil {
			log.Println("gagal ambil data di database")
			return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
		}
	}

	return user, nil
}

// GetUserByID mendapatkan pengguna berdasarkan Email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan Email: %v", err)
	}
	return user, nil
}

// UpdateUser memperbarui data pengguna
func (s *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	// err := s.userRepo.DeleteUser(ctx, id)
	// if err != nil {
	// 	return fmt.Errorf("gagal menghapus pengguna: %v", err)
	// }
	// return nil

	// delete redis key
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	if err := s.rdb.Del(ctx, redisKey).Err(); err != nil {
		log.Println("gagal delete key redis", redisKey)
	}
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}
