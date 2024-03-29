package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/henlies/project/entity"
)

func GetPetID(c *fiber.Ctx) error {
	var pets entity.Pet
	id := c.Params("id")
	if err := entity.DB().Preload("Type").Preload("Gene").Raw("SELECT * FROM pets WHERE id = ?", id).Find(&pets).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": pets})
}

func CreatePet(c *fiber.Ctx) error {
	var pet entity.Pet
	var typex entity.Type
	var gene entity.Gene
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if tx := entity.DB().Where("id = ?", pet.TypeID).First(&typex); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Type not found"})
	}
	if tx := entity.DB().Where("id = ?", pet.GeneID).First(&gene); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gene not found"})
	}

	cp := entity.Pet{
		Name:     pet.Name,
		Type:     typex,
		Gene:     gene,
		Food:     pet.Food,
		Habit:    pet.Habit,
		Descript: pet.Descript,
		Pill:     pet.Pill,
		Pic:      pet.Pic,
	}
	if err := entity.DB().Create(&cp).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ======================================== UPDATE USER ========================================
	var user entity.ServiceUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	uptu := entity.ServiceUser{
		PetID: cp.ID,
	}
	if err := entity.DB().Where("id = ?", user.ID).Updates(&uptu).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": cp})
}

func UpdatePet(c *fiber.Ctx) error {
	var pet entity.Pet
	var typex entity.Type
	var gene entity.Gene
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if tx := entity.DB().Where("id = ?", pet.TypeID).First(&typex); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Type not found"})
	}
	if tx := entity.DB().Where("id = ?", pet.GeneID).First(&gene); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gene not found"})
	}

	up := entity.Pet{
		Name:     pet.Name,
		Type:     typex,
		Gene:     gene,
		Food:     pet.Food,
		Habit:    pet.Habit,
		Descript: pet.Descript,
		Pill:     pet.Pill,
		Pic:      pet.Pic,
	}
	if err := entity.DB().Where("id = ?", pet.ID).Updates(&up).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": pet.ID})
}
