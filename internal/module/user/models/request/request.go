package request

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetUser struct {
	ID int `json:"id"`
}

type UpdateUser struct {
	ID    int    `json:"id" validate:"required"`
	Email string `json:"email" validate:"email"`
}

type CreateProfile struct {
	UserID         int    `json:"user_id" validate:"required"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name"`
	Address        string `json:"address"`
	District       string `json:"district"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Phone          string `json:"phone"`
	PersonalID     string `json:"personal_id"`
	TypePersonalID string `json:"type_personal_id"`
}

type ValidateToken struct {
	Token string `json:"token" validate:"required"`
}

type GetProfile struct {
	UserID int `json:"user_id" validate:"required"`
}

type UpdateProfile struct {
	ID             int    `json:"id" validate:"required"`
	UserID         int    `json:"user_id" validate:"required"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name"`
	Address        string `json:"address"`
	District       string `json:"district"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Phone          string `json:"phone"`
	PersonalID     string `json:"personal_id"`
	TypePersonalID string `json:"type_personal_id"`
}
