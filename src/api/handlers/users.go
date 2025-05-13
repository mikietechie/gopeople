package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mikietechie/gopeople/services"
	"github.com/mikietechie/gopeople/tp"
)

// Read Users godoc
// @Summary 			Read Users
// @Description			Read Users
// @Tags			 	users
// @Accept				json
// @Produce				json
// @Param				limit		query     int  false  "paginate with limit e.g 10"
// @Param				offset		query     int  false  "paginate with offset e.g 5"
// @Param				id			query     int  false  "search by id"
// @Param				name		query     string  false  "search by name"
// @Param				surname		query     string  false  "search by surname"
// @Param				patronymic	query     string  false  "search by patronymic"
// @Param				gender		query     string  false  "search by gender"
// @Param				age			query     string  false  "search by age"
// @Param				nationality	query     string  false  "search by nationality"
// @Success				200 {object}	[]tp.UserItem
// @Success				400 {object}	string
// @Router 				/users		[get]
func ReadUsers(c *fiber.Ctx) error {
	var pagination tp.Pagination
	var query tp.ReadUsersReqQuery
	c.QueryParser(&query)
	pagination.Limit = c.QueryInt("limit", 10)
	pagination.Offset = c.QueryInt("offset", 0)
	data, err := services.ReadUsers(query, pagination)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

// Create User godoc
// @Summary			Create User
// @Description		Create User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			body		body tp.CreateUpdateUserReqBody true "body"
// @Success			200			{object} int
// @Failure			400			{object} string
// @Router			/users		[post]
func CreateUser(c *fiber.Ctx) error {
	var body tp.CreateUpdateUserReqBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	id, err := services.CreateUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(id)
}

// Read User godoc
// @Summary			Read User
// @Description		Read User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			id   path      int  true  "User ID"
// @Success			200 {object} models.User
// @Failure			400 {object} string
// @Failure			404
// @Router			/users/{id} [get]
func ReadUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(404)
	}
	user, err := services.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// Update User godoc
// @Summary			Update User
// @Description		Update User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			id   path      int  true  "User ID"
// @Param			body body tp.CreateUpdateUserReqBody true "User data"
// @Success			200 {object} int
// @Failure			400 {object} string
// @Failure			404
// @Router			/users/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	var body tp.CreateUpdateUserReqBody
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(404)
	}
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err = services.UpdateUser(id, body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(id)
}

// Delete User godoc
// @Summary			Delete User
// @Description		Delete User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			id   path      int  true  "User ID"
// @Success			200
// @Failure			400 {object} string
// @Failure			404
// @Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(404)
	}
	err = services.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
