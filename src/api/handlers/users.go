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
// @Success				500 {object}	string
// @Router 				/users		[get]
func ReadUsers(c *fiber.Ctx) error {
	var pagination tp.Pagination
	var query tp.ReadUsersReqQuery
	err := c.QueryParser(&query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	pagination.Limit = c.QueryInt("limit", 10)
	pagination.Offset = c.QueryInt("offset", 0)
	data, err := services.ReadUsers(query, pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

// Create User godoc
// @Summary			Create User
// @Description		Create User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			body		body tp.CreateUserReqBody true "body"
// @Success			200			{object} int
// @Success			400 {object}	string
// @Success			500 {object}	string
// @Router			/users		[post]
func CreateUser(c *fiber.Ctx) error {
	var body tp.CreateUserReqBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	id, err := services.CreateUser(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
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
		return c.SendStatus(fiber.StatusNotFound)
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
// @Param			body body tp.UpdateUserReqBody true "User data"
// @Success			200 {object} int
// @Failure			400 {object} string
// @Failure			404
// @Success			500 {object}	string
// @Router			/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	var body tp.UpdateUserReqBody
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err = services.UpdateUser(id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(id)
}

// Edit User godoc
// @Summary			Edit User
// @Description		Edit User
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			id   path      int  true  "User ID"
// @Param			body body tp.EditReqBody true "User data"
// @Success			200 {object} int
// @Failure			400 {object} string
// @Failure			404
// @Success			500 {object}	string
// @Router			/users/{id} [patch]
func EditUser(c *fiber.Ctx) error {
	var body tp.EditReqBody
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err = services.EditUser(id, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
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
// @Success			500 {object}	string
// @Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	err = services.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
