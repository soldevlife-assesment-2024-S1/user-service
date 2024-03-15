package response

type RegisterResponse struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
}

type GetUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateUserResponse struct {
	ID int `json:"id"`
}

type ValidateTokenResponse struct {
	Valid bool `json:"valid"`
}

type GetProfileResponse struct {
	ID             int    `json:"id"`
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
