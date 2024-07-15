package domain

import "time"

type Customer struct {
	Id           int
	UserId       int
	FullName     string
	BirthPlace   string
	BirthDate    time.Time
	Salary       float32
	IdentityCard string
	SelfiePhoto  string
	Pin          string
}
