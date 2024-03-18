package entity

import (
	"database/sql"
)

type User struct {
	ID        int          `json:"id" db:"id"`
	Email     string       `json:"email" db:"email"`
	Password  string       `json:"password" db:"password"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type Profile struct {
	ID             int          `json:"id" db:"id"`
	UserID         int          `json:"user_id" db:"user_id"`
	FirstName      string       `json:"first_name" db:"first_name"`
	LastName       string       `json:"last_name" db:"last_name"`
	Address        string       `json:"address" db:"address"`
	District       string       `json:"district" db:"district"`
	City           string       `json:"city" db:"city"`
	State          string       `json:"state" db:"state"`
	Country        string       `json:"country" db:"country"`
	Region         string       `json:"region" db:"region"` // Continent
	Phone          string       `json:"phone" db:"phone"`
	PersonalID     string       `json:"personal_id" db:"personal_id"`
	TypePersonalID string       `json:"type_personal_id" db:"type_personal_id"` // DNI, NIE, Passport
	CreatedAt      sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
