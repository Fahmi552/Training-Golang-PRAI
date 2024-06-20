package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"Training/Session-6-db-pgx2/entity"
	"Training/Session-6-db-pgx2/service"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	user := &entity.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(*user, nil)

		createdUser, err := userService.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, *user, createdUser)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(entity.User{}, errors.New("failed to create user"))

		createdUser, err := userService.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
		assert.Equal(t, entity.User{}, createdUser)
	})
}

func TestGetUserByID(t *testing.T) {
	mockRepo := &MockIUserRepository{}
	userService := service.NewUserService(mockRepo)

	user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
	createdUser, err := userService.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	t.Run("GetUserByID - Success", func(t *testing.T) {
		retrievedUser, err1 := userService.GetUserByID(context.Background(), user.ID)
		assert.NoError(t, err1)
		assert.Equal(t, createdUser.ID, retrievedUser.ID)
		assert.Equal(t, createdUser.Name, retrievedUser.Name)
	})

	t.Run("GetUserByID - UserNotFound", func(t *testing.T) {
		_, err := userService.GetUserByID(context.Background(), user.ID)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

// func TestUpdateUser(t *testing.T) {
// 	mockRepo := &MockIUserRepository{}
// 	userService := service.NewUserService(mockRepo)

// 	user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
// 	createdUser := userService.CreateUser(context.Background(), user)

// 	t.Run("UpdateUser - Success", func(t *testing.T) {
// 		updatedUser := entity.User{Name: "Updated", Email: "updated@example.com", Password: "password"}
// 		result, err := userService.UpdateUser(createdUser.ID, updatedUser)

// 		assert.NoError(t, err)
// 		assert.Equal(t, "Updated", result.Name)
// 		assert.Equal(t, "updated@example.com", result.Email)
// 	})

// 	t.Run("UpdateUser - UserNotFound", func(t *testing.T) {
// 		updatedUser := entity.User{Name: "Updated", Email: "updated@example.com", Password: "password"}
// 		_, err := userService.UpdateUser(99, updatedUser)

// 		assert.Error(t, err)
// 		assert.Equal(t, "user not found", err.Error())
// 	})
// }

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockIUserRepository(ctrl)

	userService := service.NewUserService(mockRepo)

	// user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
	mockRepo.EXPECT().DeleteUser(gomock.Any(), 1).Return(nil)
	mockRepo.EXPECT().DeleteUser(gomock.Any(), 99).Return(nil)
	mockRepo.EXPECT().DeleteUser(gomock.Any(), 2).Return(errors.New("db error"))
	err := userService.DeleteUser(context.Background(), 1)
	assert.NoError(t, err)

	err2 := userService.DeleteUser(context.Background(), 2)
	assert.Error(t, err2)

	t.Run("DeleteUser - Success", func(t *testing.T) {
		err := userService.DeleteUser(context.Background(), 99)
		assert.NoError(t, err)
	})

	// t.Run("DeleteUser - UserNotFound", func(t *testing.T) {
	// 	err := userService.DeleteUser(context.Background(), user.ID)
	// 	assert.Error(t, err)
	// 	assert.Equal(t, "user not found", err.Error())
	// })
}

// func TestGetAllUsers(t *testing.T) {
// 	mockRepo := &MockIUserRepository{}
// 	userService := service.NewUserService(mockRepo)

// 	user1 := &entity.User{Name: "Test1", Email: "test1@example.com", Password: "password"}
// 	user2 := &entity.User{Name: "Test2", Email: "test2@example.com", Password: "password"}

// 	userService.CreateUser(user1)
// 	userService.CreateUser(user2)

// 	t.Run("GetAllUsers - Success", func(t *testing.T) {
// 		users := userService.GetAllUsers()
// 		assert.Equal(t, 2, len(users))
// 		assert.Equal(t, "Test1", users[0].Name)
// 		assert.Equal(t, "Test2", users[1].Name)
// 	})
// }
