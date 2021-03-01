package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sadiqultra/handyman/auth/jwt"
	"github.com/sadiqultra/handyman/auth/middleware"
	"github.com/sadiqultra/handyman/graph/generated"
	"github.com/sadiqultra/handyman/graph/model"
	"github.com/sadiqultra/handyman/repository/services"
	"github.com/sadiqultra/handyman/repository/users"
	"github.com/sadiqultra/handyman/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Phone = input.Phone
	user.Address = input.Address
	user.City = input.City
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateService(ctx context.Context, input model.NewService) (*model.Service, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.Service{}, fmt.Errorf("access denied")
	}

	var service services.Service
	service.Title = input.Title
	service.Description = input.Description
	service.Price = input.Price

	service.User = user
	serviceID := service.Save()
	grahpqlUser := &model.User{
		ID:      user.ID,
		Name:    user.Username,
		Address: user.Address,
		City:    user.City,
	}
	return &model.Service{
		ID: strconv.FormatInt(serviceID, 10), Title: service.Title,
		Description: service.Description, User: grahpqlUser}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &utils.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	var resultServices []*model.Service
	var dbServices []services.Service
	dbServices = services.GetAll()
	for _, service := range dbServices {
		grahpqlUser := &model.User{
			ID:      service.User.ID,
			Name:    service.User.Username,
			Address: service.User.Address,
			City:    service.User.City}

		resultServices = append(resultServices, &model.Service{
			ID:          service.ID,
			Title:       service.Title,
			Description: service.Description,
			User:        grahpqlUser})
	}
	return resultServices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
