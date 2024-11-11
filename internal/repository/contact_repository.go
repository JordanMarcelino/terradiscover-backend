package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/entity"
)

type ContactRepository interface {
	Search(ctx context.Context, params *dto.SearchContactRequest) ([]*entity.Contact, error)
	Save(ctx context.Context, contact *entity.Contact) error
}

type contactRepositoryImpl struct {
	db dbtx
}

func NewContactRepository(db dbtx) *contactRepositoryImpl {
	return &contactRepositoryImpl{
		db: db,
	}
}

func (r *contactRepositoryImpl) Search(ctx context.Context, params *dto.SearchContactRequest) ([]*entity.Contact, error) {
	query := `
		SELECT id, full_name, email, phone
		FROM contacts
		WHERE user_id = $1
	`
	args := []any{params.UserID}
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString(query)

	if params.Name != "" {
		queryBuilder.WriteString(fmt.Sprintf("AND full_name ILIKE $%d", len(args)+1))
		args = append(args, fmt.Sprintf("%%%s%%", params.Name))
	}
	if params.Email != "" {
		queryBuilder.WriteString(fmt.Sprintf("AND email ILIKE $%d", len(args)+1))
		args = append(args, fmt.Sprintf("%%%s%%", params.Email))
	}
	if params.Phone != "" {
		queryBuilder.WriteString(fmt.Sprintf("AND phone ILIKE $%d", len(args)+1))
		args = append(args, fmt.Sprintf("%%%s%%", params.Phone))
	}

	rows, err := r.db.QueryContext(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contacts := []*entity.Contact{}
	for rows.Next() {
		contact := new(entity.Contact)
		if err := rows.Scan(&contact.ID, &contact.FullName, &contact.Email, &contact.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *contactRepositoryImpl) Save(ctx context.Context, contact *entity.Contact) error {
	query := `
		INSERT INTO contacts(user_id, full_name, email, phone) VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return r.db.QueryRowContext(ctx, query, contact.UserID, contact.FullName, contact.Email, contact.Phone).
		Scan(&contact.ID)
}
