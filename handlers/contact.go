package handlers

import (
	"testcontact/database"
	"testcontact/models"
	"testcontact/utils"

	"github.com/gofiber/fiber/v2"
)

type Status struct {
	Code int32 `json:"code"`
}

func AddContact(c *fiber.Ctx) error {
	contact := new(models.Contact)
	if err := c.BodyParser(contact); err != nil {
		return err
	}

	if contact.Name == "" || contact.Tel == "" {
		utils.SendJSON(c, Status{Code: 400})
	}

	res := database.DBConn.First(&contact, "tel = ?", contact)
	if res.RowsAffected != 0 {
		utils.SendJSON(c, Status{Code: 208})
	}

	database.DBConn.Create(&contact)
	return utils.SendJSON(c, Status{Code: 201})
}

func RemoveContact(c *fiber.Ctx) error {
	contact := new(models.Contact)
	if err := c.BodyParser(contact); err != nil {
		return err
	}

	if contact.Tel == "" {
		return utils.SendJSON(c, Status{Code: 400})
	}

	res := database.DBConn.Where("tel = ?", contact.Tel).Unscoped().Delete(&contact)
	if res.RowsAffected == 0 {
		return utils.SendJSON(c, Status{Code: 204})
	}
	return utils.SendJSON(c, Status{Code: 200})
}

func EditContact(c *fiber.Ctx) error {

	return c.SendString("Edit")
}

func GetContacts(c *fiber.Ctx) error {
	ref := new([]models.Contact)
	database.DBConn.Find(&ref)
	return utils.SendJSON(c, ref)
}