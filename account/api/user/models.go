package user

import (
	"github.com/smartbot/account/database"
)

type OnboardRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type User struct {
	ID             string              `json:"id"`
	Username       string              `json:"username"`
	Firstname      string              `json:"first_name"`
	Lastname       string              `json:"last_name"`
	PrimaryAddress string              `json:"primary_address"`
	CountryCode    string              `json:"country_code"`
	Mobile         string              `json:"mobile"`
	Role           database.UserRole   `json:"role"`
	Status         database.UserStatus `json:"status"`
}
