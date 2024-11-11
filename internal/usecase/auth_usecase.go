package usecase

import (
	"context"

	"github.com/jordanmarcelino/terradiscover-backend/internal/apperror"
	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/entity"
	"github.com/jordanmarcelino/terradiscover-backend/internal/repository"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/encryptutils"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/jwtutils"
)

type AuthUseCase interface {
	Login(ctx context.Context, request *dto.UserLoginRequest) (string, error)
	Register(ctx context.Context, request *dto.UserRegisterRequest) (*dto.UserResponse, error)
}

type authUseCaseImpl struct {
	jwtUtil         jwtutils.JwtUtil
	bcryptEncryptor encryptutils.BcryptEncryptor
	dataStore       repository.DataStore
	userRepository  repository.UserRepository
}

func NewAuthUseCase(
	jwtUtil jwtutils.JwtUtil,
	bcryptEncryptor encryptutils.BcryptEncryptor,
	dataStore repository.DataStore,
	userRepository repository.UserRepository,
) *authUseCaseImpl {
	return &authUseCaseImpl{
		jwtUtil:         jwtUtil,
		bcryptEncryptor: bcryptEncryptor,
		dataStore:       dataStore,
		userRepository:  userRepository,
	}
}

func (u *authUseCaseImpl) Login(ctx context.Context, request *dto.UserLoginRequest) (string, error) {
	user, err := u.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return "", apperror.NewServerError(err)
	}
	if user == nil {
		return "", apperror.NewInvalidCredentialError()
	}

	if ok := u.bcryptEncryptor.Check(request.Password, user.Password); !ok {
		return "", apperror.NewInvalidCredentialError()
	}

	token, err := u.jwtUtil.Sign(user.ID)
	if err != nil {
		return "", apperror.NewServerError(err)
	}

	return token, err
}

func (u *authUseCaseImpl) Register(ctx context.Context, request *dto.UserRegisterRequest) (*dto.UserResponse, error) {
	res := &dto.UserResponse{Email: request.Email}
	err := u.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
		userRepository := ds.User()

		user, err := userRepository.FindByEmail(ctx, request.Email)
		if err != nil {
			return apperror.NewServerError(err)
		}
		if user != nil {
			return apperror.NewUserAlreadyExistsError()
		}

		hashPwd, err := u.bcryptEncryptor.Hash(request.Password)
		if err != nil {
			return apperror.NewServerError(err)
		}

		newUser := &entity.User{Email: request.Email, Password: hashPwd}
		if err := userRepository.Save(ctx, newUser); err != nil {
			return apperror.NewServerError(err)
		}

		res.ID = newUser.ID
		return nil
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}
