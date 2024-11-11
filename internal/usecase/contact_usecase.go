package usecase

import (
	"context"

	"github.com/jordanmarcelino/terradiscover-backend/internal/apperror"
	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/entity"
	"github.com/jordanmarcelino/terradiscover-backend/internal/repository"
)

type ContactUseCase interface {
	Search(ctx context.Context, request *dto.SearchContactRequest) ([]*dto.ContactResponse, *dto.PageMetaData, error)
	Save(ctx context.Context, request *dto.CreateContactRequest) (*dto.ContactResponse, error)
}

type contactUseCaseImpl struct {
	contactRepository repository.ContactRepository
}

func NewContactUseCase(contactRepository repository.ContactRepository) *contactUseCaseImpl {
	return &contactUseCaseImpl{
		contactRepository: contactRepository,
	}
}

func (u *contactUseCaseImpl) Search(ctx context.Context, request *dto.SearchContactRequest) ([]*dto.ContactResponse, *dto.PageMetaData, error) {
	panic("")
}

func (u *contactUseCaseImpl) Save(ctx context.Context, request *dto.CreateContactRequest) (*dto.ContactResponse, error) {
	contact := &entity.Contact{
		UserID:   request.UserID,
		FullName: request.FullName,
		Email:    request.Email,
		Phone:    request.Phone,
	}
	if err := u.contactRepository.Save(ctx, contact); err != nil {
		return nil, apperror.NewServerError(err)
	}

	return dto.ConvertToContactResponse(contact), nil
}
