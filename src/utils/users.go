package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mikietechie/gopeople/config"
	"github.com/mikietechie/gopeople/tp"
	"gorm.io/gorm/clause"
)

func ParseReadUsersQuery(query tp.ReadUsersReqQuery) ([]clause.Expression, error) {
	whereExpressions := []clause.Expression{clause.Neq{Column: "id", Value: 0}}

	if query.Name != "" {
		whereExpressions = append(whereExpressions, clause.Eq{Column: "name", Value: query.Name})
	}
	if query.Surname != "" {
		whereExpressions = append(whereExpressions, clause.Eq{Column: "surname", Value: query.Surname})
	}
	if query.Age != 0 {
		whereExpressions = append(whereExpressions, clause.Eq{Column: "age", Value: query.Age})
	}
	if query.Nationality != "" {
		whereExpressions = append(whereExpressions, clause.Eq{Column: "nationality", Value: query.Nationality})
	}
	if query.Gender != "" {
		whereExpressions = append(whereExpressions, clause.Eq{Column: "gender", Value: query.Gender})
	}
	return whereExpressions, nil
}

func GetAgeByName(name string) (int, error) {
	agent := fiber.Get(fmt.Sprintf(config.AGE_API_URL, name))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error(err.Error())
		}
		return 0, errors.New("call to api.agify failed")
	}
	var data struct {
		Age int `json:"age"`
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}
	return data.Age, nil
}

func GetNationalityByName(name string) (string, error) {
	agent := fiber.Get(fmt.Sprintf(config.NATION_API_URL, name))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error(err.Error())
		}
		return "", errors.New("call to api.agify failed")
	}
	var data struct {
		Country []struct {
			CountryId string `json:"country_id"`
		} `json:"country"`
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	if len(data.Country) == 0 {
		return "", errors.New("country not found")
	}
	return data.Country[0].CountryId, nil
}

func GetGenderByName(name string) (tp.Gender, error) {
	agent := fiber.Get(fmt.Sprintf(config.GENDER_API_URL, name))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error(err.Error())
		}
		return "", errors.New("call to api.genderize failed")
	}
	var data struct {
		Gender string `json:"gender"`
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	gender := tp.Other
	if data.Gender == "male" {
		gender = tp.Male
	} else if data.Gender == "female" {
		gender = tp.Female
	}
	return gender, nil
}
