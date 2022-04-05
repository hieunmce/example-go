package user

import (
	"context"
	"example.com/m/domain"
	"example.com/m/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// CreateData data for CreateUser
type CreateData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateRequest request struct for CreateUser
type CreateRequest struct {
	User CreateData `json:"user"`
}

// CreateResponse response struct for CreateUser
type CreateResponse struct {
	User  domain.User  `json:"user"`
	Token domain.Token `json:"token"`
}

// StatusCode customstatus code for success create User
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a User

//swagger
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			user = &domain.User{
				Name:     req.User.Name,
				Email:    req.User.Email,
				Password: req.User.Password,
			}
		)

		err := s.UserService.Create(ctx, user)
		if err != nil {
			return nil, err
		}
		var userResp = domain.User{
			Name:  user.Name,
			Email: user.Email,
			Model: user.Model,
		}

		tokenString, err := generateToken(7, jwt.MapClaims{
			"name":  userResp.Name,
			"email": userResp.Email,
		})
		refreshTokenString, err := generateToken(30, jwt.MapClaims{
			"name":  userResp.Name,
			"email": userResp.Email,
		})
		tokenResp := domain.Token{
			Token:        tokenString,
			RefreshToken: refreshTokenString,
		}
		return CreateResponse{User: userResp, Token: tokenResp}, nil
	}
}

// FindRequest request struct for Find a User
type FindRequest struct {
	UserID domain.UUID
}

// FindResponse response struct for Find a User
type FindResponse struct {
	User *domain.User `json:"user"`
}

// MakeFindEndPoint make endpoint for find User
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var userFind domain.User
		req := request.(FindRequest)
		userFind.ID = req.UserID

		user, err := s.UserService.Find(ctx, &userFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{User: user}, nil
	}
}

// FindAllRequest request struct for FindAll User
type FindAllRequest struct{}

// FindAllResponse request struct for find all User
type FindAllResponse struct {
	Users []domain.User `json:"users"`
}

// MakeFindAllEndpoint make endpoint for find all User
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		users, err := s.UserService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Users: users}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID    domain.UUID `json:"-"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	User UpdateData `json:"user"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	User domain.User `json:"user"`
}

// MakeUpdateEndpoint make endpoint for update a User
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			user = domain.User{
				Model: domain.Model{ID: req.User.ID},
				Name:  req.User.Name,
				Email: req.User.Email,
			}
		)

		res, err := s.UserService.Update(ctx, &user)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{User: *res}, nil
	}
}

// DeleteRequest request struct for delete a User
type DeleteRequest struct {
	UserID domain.UUID
}

// DeleteResponse response struct for Find a User
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a User
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			userFind = domain.User{}
			req      = request.(DeleteRequest)
		)
		userFind.ID = req.UserID

		err := s.UserService.Delete(ctx, &userFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
