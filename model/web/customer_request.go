package web

type CustomerRequest struct {
	UserId       int     `validate:"required" json:"user_id"`
	FullName     string  `validate:"required" json:"full_name"`
	BirthPlace   string  `validate:"required" json:"birth_place"`
	BirthDate    string  `validate:"required" json:"birth_date"`
	Salary       float32 `validate:"required" json:"salary"`
	IdentityCard string  `validate:"required" json:"identity_card"`
	SelfiePhoto  string  `validate:"required" json:"selfie_photo"`
	Pin          string  `validate:"required" json:"pin"`
}
