package controllers

import (
	"strconv"
	"time"

	"app/configs"
	"app/database"
	"app/helpers"
	"app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var SecretKey = configs.Env("JWT_SECRET_KEY")

var JwtCookieName = "session"

// Register Function
func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	password, _ := helpers.HashPassword(user.Password)

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	database.DB.Create(&newUser)

	return c.JSON(newUser)
}

// Login Function
func Login(c *fiber.Ctx) error {
	var data models.User
	// Parsing the data that has been entered by User
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	if err := helpers.CompareHashAndPassword(user.Password, data.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong",
		})
	}

	cookie := fiber.Cookie{
		Name:     JwtCookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Success",
	})
}

// User  Function
func User(c *fiber.Ctx) error {
	cookie := c.Cookies(JwtCookieName)

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

// Logout Function
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     JwtCookieName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Success",
	})
}
