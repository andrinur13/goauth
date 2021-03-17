package user

import "time"

type User struct {
	ID        int
	Username  string
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
