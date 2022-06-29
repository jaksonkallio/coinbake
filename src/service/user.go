package service

import (
	"time"

	"github.com/jaksonkallio/coinbake/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	EmailAddress string
	LastAuthed   time.Time
	LastActive   time.Time
}

func FindUserByEmailAddress(emailAddress string) *User {
	user := User{}
	database.Handle().Where("email_address = ?", emailAddress).First(&user)
	return &user
}

// Marks the user as authed and active.
func (user *User) MarkLastAuthed() {
	database.Handle().Model(&user).Updates(
		map[string]interface{}{
			"last_authed": time.Now(),
			"last_active": time.Now(),
		},
	)
}

// Marks the user as active.
func (user *User) MarkLastActive() {
	database.Handle().Model(&user).Update("last_authed", time.Now())
}
