package controllers

import (
	_ "go/token"
	"strconv"
	"time"

	"github.com/ADEMOLA200/go_admin_application/database"
	"github.com/ADEMOLA200/go_admin_application/models"
	"github.com/ADEMOLA200/go_admin_application/util"
	 _"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["password"] != data["confirm_password"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(&user)
}


func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// Retrieve user from the database based on the email
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Check if the user exists
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Check if the provided password is empty
	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "password cannot be empty",
		})
	}

	// Compare the hashed password from the database with the provided password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	/******************************** JWT **************************************/
	// create a token
	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer: strconv.Itoa(int(user.Id)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// 	Audience: user.Email,
	// })

	// token, err := claims.SignedString([]byte("secret"))


	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	/******************************************************************************/

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Create cookies
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}


func User(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")

	id, _:= util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}


func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout succesful",
	})
}