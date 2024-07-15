package web

type RegisterRequest struct {
	Username    string `validate:"required" json:"username"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	RoleId      int    `validate:"required" json:"role_id"`
}
