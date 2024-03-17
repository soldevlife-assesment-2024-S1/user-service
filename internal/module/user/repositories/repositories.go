package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/module/user/models/entity"
	"user-service/internal/pkg/helpers/errors"
	"user-service/internal/pkg/log"

	"github.com/jmoiron/sqlx"
)

type repositories struct {
	db  *sqlx.DB
	log log.Logger
}

// FindProfileByID implements Repositories.
func (r *repositories) FindProfileByID(ctx context.Context, id int) (entity.Profile, error) {
	query := `SELECT id, user_id, address, district, city, state, country, region, phone, personal_id, type_personal_id, created_at, updated_at FROM profiles WHERE id = $1`

	var profile entity.Profile
	err := r.db.GetContext(ctx, &profile, query, id)

	// check row not found
	if err == sql.ErrNoRows {
		return profile, errors.NotFound("record not found")
	}

	if err != nil {
		return profile, errors.InternalServerError(fmt.Sprintf("error finding profile by id: %s", err.Error()))
	}

	return profile, nil
}

// FindUserByEmail implements Repositories.
func (r *repositories) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE email = $1`

	var user entity.User
	err := r.db.GetContext(ctx, &user, query, email)

	// check row not found
	if err == sql.ErrNoRows {
		return user, nil
	}

	if err != nil {
		return user, errors.InternalServerError(fmt.Sprintf("error finding user by email: %s", err.Error()))
	}

	return user, nil
}

// FindUserByID implements Repositories.
func (r *repositories) FindUserByID(ctx context.Context, id int) (entity.User, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at FROM users WHERE id = $1`

	var user entity.User

	err := r.db.GetContext(ctx, &user, query, id)

	// check row not found
	if err == sql.ErrNoRows {
		return user, nil
	}

	if err != nil {
		return user, errors.InternalServerError(fmt.Sprintf("error finding user by id: %s", err.Error()))
	}

	return user, nil
}

// UpsertProfile implements Repositories.
func (r *repositories) UpsertProfile(ctx context.Context, payload *entity.Profile) error {
	query := `INSERT INTO profiles (id, user_id, address, district, city, state, country, region, phone, personal_id, type_personal_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW()) ON CONFLICT (id) DO UPDATE SET address = $2, district = $3, city = $4, state = $5, country = $6, region = $7, phone = $8, personal_id = $9, type_personal_id = $10, updated_at = NOW() RETURNING id`

	err := r.db.QueryRowContext(ctx, query, payload.ID, payload.UserID, payload.Address, payload.District, payload.City, payload.State, payload.Country, payload.Region, payload.Phone, payload.PersonalID, payload.TypePersonalID).Scan(&payload.ID)

	// check row not found
	if err == sql.ErrNoRows {
		return errors.NotFound("record not found")
	}

	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error upserting profile on repository: %s", err.Error()))
	}

	return nil
}

// UpsertUser implements Repositories.
func (r *repositories) UpsertUser(ctx context.Context, payload *entity.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password, created_at) VALUES ($1, $2, $3, $4, NOW()) ON CONFLICT (email) DO UPDATE SET first_name = $1, last_name = $2, password = $3, updated_at = NOW() RETURNING id`

	err := r.db.QueryRowContext(ctx, query, payload.FirstName, payload.LastName, payload.Email, payload.Password).Scan(&payload.ID)

	// check row not found
	if err == sql.ErrNoRows {
		return errors.NotFound("record not found")
	}

	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error upserting user on repository: %s", err.Error()))
	}

	return nil
}

type Repositories interface {
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
	FindUserByID(ctx context.Context, id int) (entity.User, error)
	UpsertUser(ctx context.Context, payload *entity.User) error
	FindProfileByID(ctx context.Context, id int) (entity.Profile, error)
	UpsertProfile(ctx context.Context, payload *entity.Profile) error
}

func New(db *sqlx.DB, log log.Logger) Repositories {
	return &repositories{
		db:  db,
		log: log,
	}
}
