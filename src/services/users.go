package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/mikietechie/gopeople/models"
	"github.com/mikietechie/gopeople/tp"
	"github.com/mikietechie/gopeople/utils"
	"gorm.io/gorm/clause"
)

func setUserDetails(u *models.User) error {
	var err error
	name := strings.TrimSpace(u.Name)
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	u.Nationality, err = utils.GetNationalityByName(name)
	if err != nil {
		log.Error(err)
		return err
	}
	u.Age, err = utils.GetAgeByName(name)
	if err != nil {
		log.Error(err)
		return err
	}
	u.Gender, err = utils.GetGenderByName(name)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func CreateUser(body tp.CreateUpdateUserReqBody) (int, error) {
	var user models.User
	user.Name = body.Name
	user.Surname = body.Surname
	user.Patronymic = body.Patronymic
	err := setUserDetails(&user)
	if err != nil {
		log.Error(fmt.Sprintf(`Failed set details %s`, user))
		log.Error(err)
		return 0, err
	}
	err = models.Db.Create(&user).Error
	if err != nil {
		log.Error(fmt.Sprintf(`Failed to update %s\n`, user))
		return 0, err
	}
	return int(user.ID), nil
}

func GetUserByID(userId int) (*models.User, error) {
	var user models.User
	err := models.Db.Model(&models.User{}).First(&user, "id = ?", userId).Error
	if err != nil {
		log.Errorf(`Failed to find User with ID=%d`, userId)
		return &user, err
	}
	return &user, nil
}

func UpdateUser(userId int, body tp.CreateUpdateUserReqBody) error {
	user, err := GetUserByID(userId)
	if err != nil {
		return err
	}
	user.Name = body.Name
	user.Surname = body.Surname
	user.Patronymic = body.Patronymic

	err = models.Db.Save(user).Error
	if err != nil {
		log.Errorf(`Failed to update %s\n`, user)
	}
	return err
}

func DeleteUser(userId int) error {
	user, err := GetUserByID(userId)
	if err != nil {
		return err
	}
	err = models.Db.Delete(&user).Error
	if err != nil {
		log.Errorf(`Failed to delete %s\n`, user)
	}
	return err
}

func ReadUsers(query tp.ReadUsersReqQuery, pagination tp.Pagination) ([]tp.UserItem, error) {
	var items []tp.UserItem
	whereExpressions, _ := utils.ParseReadUsersQuery(query)
	err := models.Db.Model(
		&models.User{},
	).Where(
		clause.Where{Exprs: whereExpressions},
	).Limit(pagination.Limit).Offset(pagination.Offset).Find(&items).Error
	return items, err
}
