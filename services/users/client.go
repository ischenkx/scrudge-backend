package users

import "context"

type UserDto struct {
	Username, ID string
}

type LoginDto struct {
	Username, Password string
}

type RegisterDto struct {
	Username, Password string
}

type Client interface {
	Login(ctx context.Context, dto LoginDto) (UserDto, error)
	Register(ctx context.Context, dto RegisterDto) (UserDto, error)
	Get(ctx context.Context, id string) (UserDto, error)
	GetByName(ctx context.Context, name string) (UserDto, error)
	UpdateUsername(ctx context.Context, id, username string) error
	UpdatePassword(ctx context.Context, id, prevPassword, password string) error
}
