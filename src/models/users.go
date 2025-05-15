package models

import (
	"fmt"

	"github.com/mikietechie/gopeople/tp"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"name" gorm:"index"`                    // Single-field index
	Surname     string    `json:"surname" gorm:"index"`                 // Single-field index
	Patronymic  string    `json:"patronymic" gorm:"default:'';"`        // No index (rarely filtered)
	Gender      tp.Gender `json:"gender" gorm:"type:gender_enum;index"` // Single-field index
	Age         int       `json:"age" gorm:"index"`                     // Single-field index
	Nationality string    `json:"nationality" gorm:"index"`             // Single-field index
}

func (user User) String() string {
	return fmt.Sprintf(`<User %d - %s>`, user.ID, user.Name)
}

func (user User) Fullname() string {
	return fmt.Sprintf(`%s %s %s`, user.Name, user.Patronymic, user.Surname)
}
