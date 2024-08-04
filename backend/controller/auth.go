package controller

import (
	"github.com/Sahilb315/trello_clone/database"
	"github.com/Sahilb315/trello_clone/logic"
	"github.com/Sahilb315/trello_clone/models"
	"github.com/Sahilb315/trello_clone/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignUpData struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Login(c *fiber.Ctx) error {

	db := database.Database.Db
	user := LoginData{}
	if err := c.BodyParser(&user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid login data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input data", nil)
	}

	var dbUser models.User
	if err := db.Table("users").Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return utils.SendResponse(c, fiber.StatusNotFound, "User not found", nil)
	}
	isPassSame := logic.ComparePassword(dbUser.Password, user.Password)

	if isPassSame {
		token, err := logic.CreateToken(dbUser)
		if err != nil {
			return utils.SendResponse(c, fiber.StatusInternalServerError, "Error creating token", nil)
		}
		return utils.SendResponse(c, fiber.StatusOK, "Login successful", map[string]string{
			"token": token,
		})
	} else {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "Incorrect Password", nil)
	}
}

func SignUp(c *fiber.Ctx) error {
	var userData SignUpData
	db := database.Database.Db
	if err := c.BodyParser(&userData); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid signup data", nil)
	}

	validate := validator.New()
	if err := validate.Struct(userData); err != nil {
		return utils.SendResponse(c, fiber.StatusBadRequest, "Invalid input data", nil)
	}

	hashedPassword, err := logic.HashPassword(userData.Password)
	if err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Error hashing password", nil)
	}

	newUser := models.User{
		FullName: userData.FullName,
		Email:    userData.Email,
		Password: hashedPassword,
	}
	var user models.User
	if err := db.Table("users").Where("email = ?", newUser.Email).Find(&user).Error; err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Error searching user", nil)
	}

	if user.ID != 0{
		return utils.SendResponse(c, fiber.StatusConflict, "Email already exists", nil)
	}

	if err := db.Table("users").Create(&newUser).Error; err != nil {
		return utils.SendResponse(c, fiber.StatusInternalServerError, "Error creating user", nil)
	}
	return utils.SendResponse(c, fiber.StatusCreated, "User created successfully", nil)
}
