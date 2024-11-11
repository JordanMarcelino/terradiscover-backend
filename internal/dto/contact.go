package dto

import "github.com/jordanmarcelino/terradiscover-backend/internal/entity"

type ContactResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CreateContactRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,phone_number"`
	UserID   int64
}

type SearchContactRequest struct {
	Email  string `form:"email"`
	Name   string `form:"name"`
	Phone  string `form:"phone"`
	Page   int64  `form:"page" binding:"gte=1"`
	Size   int64  `form:"size" binding:"gte=1,max=20"`
	UserID int64
}

func ConvertToContactResponses(contacts []*entity.Contact) []*ContactResponse {
	res := []*ContactResponse{}
	for _, contact := range contacts {
		res = append(res, ConvertToContactResponse(contact))
	}
	return res
}

func ConvertToContactResponse(contact *entity.Contact) *ContactResponse {
	return &ContactResponse{
		ID:       contact.ID,
		FullName: contact.FullName,
		Email:    contact.Email,
		Phone:    contact.Phone,
	}
}
