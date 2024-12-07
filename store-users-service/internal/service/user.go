package service

import (
	"context"
	"users-service/internal/domain/user"
	interfaces "users-service/internal/repository/interface"
	services "users-service/internal/service/interface"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) services.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, req user.Request) (id string, err error) {
	data := user.Entity{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Roles:   req.Roles,
	}
	id, err = us.userRepository.Create(ctx, data)
	return
}

func (us *UserService) ListUsers(ctx context.Context) (res []user.Response, err error) {
	data, err := us.userRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = user.ParseFromEntities(data)
	return
}

func (us *UserService) GetUser(ctx context.Context, id string) (res user.Response, err error) {
	data, err := us.userRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = user.ParseFromEntity(data)
	return
}

func (us *UserService) DeleteUser(ctx context.Context, id string) (err error) {
	err = us.userRepository.Delete(ctx, id)
	return
}

func (us *UserService) UpdateUser(ctx context.Context, id string, req user.Request) (err error) {
	data := user.Entity{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Roles:   req.Roles,
	}
	err = us.userRepository.Update(ctx, id, data)
	return
}

func (us *UserService) SearchUser(ctx context.Context, filter, value string) (res []user.Response, err error) {
	if !user.IsValidFilter(filter) || value == "" {
		err = user.ErrorInvalidSearch
		return
	}
	data, err := us.userRepository.Search(ctx, filter, value)
	if err != nil {
		return
	}
	res = user.ParseFromEntities(data)
	return
}
