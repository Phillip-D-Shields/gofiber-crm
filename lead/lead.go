package lead

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/phillip-d-shields/go-fiber-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

func CreateLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(&lead); err != nil {
		c.Status(503).SendString(err.Error())
		return err
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).SendString("no lead with that id")
		return errors.New("no lead with that id")
	}
	db.Delete(&lead)
	c.SendString("lead successfully deleted")
	return nil
}
