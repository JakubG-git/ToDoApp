package model

type UserCreateOrUpdateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest UserCreateOrUpdateRequest
