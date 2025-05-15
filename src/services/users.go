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

func CreateUser(body tp.CreateUserReqBody) (int, error) {
	var user models.User
	user.Name = body.Name
	user.Surname = body.Surname
	user.Patronymic = body.Patronymic
	err := setUserDetails(&user)
	if err != nil {
		log.Error(fmt.Sprintf(`Failed set details %s`, user))
		return 0, err
	}
	tx := models.Db.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		log.Error(err)
		log.Error(fmt.Sprintf(`Failed to update %s\n`, user))
		return 0, err
	}
	tx.Commit()
	return int(user.ID), nil
}

func GetUserByID(userId int) (*models.User, error) {
	var user models.User
	err := models.Db.Model(&models.User{}).First(&user, "id = ?", userId).Error
	if err != nil {
		log.Error(err)
		log.Errorf(`Failed to find User with ID=%d`, userId)
		return &user, err
	}
	return &user, nil
}

func UpdateUser(userId int, body tp.UpdateUserReqBody) error {
	data, err := utils.StructToMap(body)
	if err != nil {
		log.Error(err)
		return err
	}
	tx := models.Db.Begin()
	err = models.Db.Model(&models.User{}).Where("id = ?", userId).Updates(data).Error
	if err != nil {
		tx.Rollback()
		log.Error(err)
		log.Errorf(`Failed to update User Id %d`, userId)
		return err
	}
	tx.Commit()
	return err
}

func EditUser(userId int, body tp.EditReqBody) error {
	tx := models.Db.Begin()
	err := models.Db.Model(&models.User{}).Where("id = ?", userId).Update(body.Field, body.Value).Error
	if err != nil {
		tx.Rollback()
		log.Error(err)
		log.Errorf(`Failed to edit User Id %d`, userId)
		return err
	}
	tx.Commit()
	return nil
}

func DeleteUser(userId int) error {
	user, err := GetUserByID(userId)
	if err != nil {
		log.Error(err)
		return err
	}
	tx := models.Db.Begin()
	err = models.Db.Delete(&user).Error
	if err != nil {
		tx.Rollback()
		log.Errorf(`Failed to delete %s\n`, user)
		return err
	}
	tx.Commit()
	return nil
}

func ReadUsers(query tp.ReadUsersReqQuery, pagination tp.Pagination) ([]tp.UserItem, error) {
	var items []tp.UserItem
	whereExpressions, _ := utils.ParseReadUsersQuery(query)
	err := models.Db.Model(
		&models.User{},
	).Where(
		clause.Where{Exprs: whereExpressions},
	).Limit(pagination.Limit).Offset(pagination.Offset).Order("id ASC").Find(&items).Error
	if err != nil {
		log.Error(err)
	}
	return items, err
}
