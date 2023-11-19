package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/henlies/project/entity"
	"golang.org/x/crypto/bcrypt"
)

// - ไว้เข้ารหัส
func SetupPasswordHash(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

func ListUsers(c *fiber.Ctx) error {
	var users []entity.User
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Blood").
		Preload("Signin").Raw("SELECT * FROM users").Find(&users).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": users})
}

func GetUser(c *fiber.Ctx) error {
	var user entity.User
	id := c.Params("id")
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Blood").
		Preload("Signin").Raw("SELECT * FROM users WHERE id = ?", id).Find(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
}

func CreateUser(c *fiber.Ctx) error {
	var user entity.User
	var prefix entity.Prefix
	var gender entity.Gender
	var blood entity.Blood
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if tx := entity.DB().Where("id = ?", user.PrefixID).First(&prefix); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Prefix not found"})
	}
	if tx := entity.DB().Where("id = ?", user.GenderID).First(&gender); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gender not found"})
	}
	if tx := entity.DB().Where("id = ?", user.BloodID).First(&blood); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Blood not found"})
	}

	var userrole entity.Role
	if err := entity.DB().Model(&entity.Role{}).Where("name = ?", "ผู้ใช้งานระบบ").First(&userrole).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User role not found"})
	}
	createusersignin := entity.Signin{
		User: user.User,
		Pass: SetupPasswordHash(user.Pass),
		Role: userrole,
	}

	cu := entity.User{
		Prefix:    prefix,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Nickname:  user.Nickname,
		Gender:    gender,
		Phone:     user.Phone,
		Email:     user.Email,
		Birth:     user.Birth,
		Age:       user.Age,
		Blood:     blood,
		Descript:  user.Descript,
		User:      user.User,
		Pass:      SetupPasswordHash(user.Pass),
		Pic:       user.Pic,
		Signin:    createusersignin,
	}
	if err := entity.DB().Create(&cu).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": cu})
}

// func UpdateUser(c *fiber.Ctx) error {
// 	var user entity.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	uu := entity.User{
// 		Firstname: user.Firstname,
// 		Lastname:  user.Lastname,
// 	}
// 	if err := entity.DB().Where("id = ?", user.ID).Updates(&uu).Error; err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
// }

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": id})
}