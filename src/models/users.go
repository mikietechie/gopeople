package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `gorm:"default:'';" json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

func (user User) String() string {
	return fmt.Sprintf(`<User %d - %s>`, user.ID, user.Name)
}

func (user User) Fullname() string {
	return fmt.Sprintf(`%s %s %s`, user.Name, user.Patronymic, user.Surname)
}
