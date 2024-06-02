package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/module/user/models/entity"
	"user-service/internal/pkg/helpers/errors"

	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type repositories struct {
	db  *sqlx.DB
	log *otelzap.Logger
}

// FindProfileByUserID implements Repositories.
func (r *repositories) FindProfileByUserID(ctx context.Context, userID int) (entity.Profile, error) {
	query := `SELECT id, user_id, first_name, last_name, address, district, city, state, country, region, phone, personal_id, type_personal_id, created_at, updated_at FROM profiles WHERE user_id = $1`

	var profile entity.Profile
	err := r.db.GetContext(ctx, &profile, query, userID)

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
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`

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
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`

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
	query := `INSERT INTO profiles (user_id, first_name, last_name, address, district, city, state, country, region, phone, personal_id, type_personal_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW()) ON CONFLICT (user_id) DO UPDATE SET first_name = $3, last_name = $4, address = $5, district = $6, city = $7, state = $8, country = $9, region = $10, phone = $11, personal_id = $12, type_personal_id = $13, updated_at = NOW() RETURNING id`

	err := r.db.QueryRowContext(ctx, query, payload.UserID, payload.FirstName, payload.LastName, payload.Address, payload.District, payload.City, payload.State, payload.Country, payload.Region, payload.Phone, payload.PersonalID, payload.TypePersonalID).Scan(&payload.ID)

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
	query := `INSERT INTO users (email, password, created_at) VALUES ($1, $2, NOW()) ON CONFLICT (email) DO UPDATE SET password = $2, updated_at = NOW() RETURNING id`

	err := r.db.QueryRowContext(ctx, query, payload.Email, payload.Password).Scan(&payload.ID)

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
	FindProfileByUserID(ctx context.Context, userID int) (entity.Profile, error)
	UpsertProfile(ctx context.Context, payload *entity.Profile) error
}

func New(db *sqlx.DB, log *otelzap.Logger) Repositories {
	return &repositories{
		db:  db,
		log: log,
	}
}
