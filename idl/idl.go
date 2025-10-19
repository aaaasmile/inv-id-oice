package idl

import "time"

var (
	Appname = "Inv-ido-voice"
	Buildnr = "00.001.20251019-00"
)

type User struct {
	Id             int
	Name           string
	Email          string
	Password       string
	RememberToken  string
	LastLoggedInAt time.Time
	Locale         string
	Enabled        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
