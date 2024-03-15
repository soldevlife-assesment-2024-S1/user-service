package entity

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Profile struct {
	ID             int
	UserID         int
	Address        string
	District       string
	City           string
	State          string
	Country        string
	Region         string // Continent
	Phone          string
	PersonalID     string
	TypePersonalID string // DNI, NIE, Passport
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
