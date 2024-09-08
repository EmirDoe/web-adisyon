package controllers

import (
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"webadisyon.com/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	// SecretKey is the key used to sign the JWT
	SecretKey = os.Getenv("JWT_SECRET")
)

type RegisterRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Role     int    `json:"role"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AuthStatus struct {
	IsAuthenticated bool   `json:"isAuthenticated"`
	Message         string `json:"message"`
	UserID          string `json:"userId"`
}

func Register(c *fiber.Ctx) error {
	/*activeUserRole, err := CheckUserRole(c)

	if activeUserRole != 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized to create user",
		})
	}*/

	request := &RegisterRequest{}
	if err := c.BodyParser(request); err != nil {
		return err

	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	user := models.User{
		Name:     request.Name,
		UserName: request.UserName,
		Password: password,
	}

	userExists, err := models.GetUserByUserName(request.UserName)

	if err != nil {
		return err
	}

	if userExists.ID != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	if err := models.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(request)
}

func ValidateUser(c *fiber.Ctx) AuthStatus {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return AuthStatus{IsAuthenticated: false, Message: "Unauthorized", UserID: "0"}
	}

	claims := token.Claims.(*jwt.StandardClaims)
	user, err := models.GetUserByID(claims.Issuer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = models.User{}
			return AuthStatus{IsAuthenticated: false, Message: "User not found", UserID: user.ID}
		}
	}

	return AuthStatus{IsAuthenticated: true, Message: "Authorized", UserID: user.ID}
}

func Login(c *fiber.Ctx) error {

	request := &LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return err

	}

	auth := ValidateUser(c)
	if auth.IsAuthenticated {
		return c.Status(401).JSON(fiber.Map{
			"message": "User already logged in",
			"userId":  auth.UserID,
		})
	}

	user, err := models.GetUserByUserName(request.UserName)

	if err != nil {
		return err
	}

	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {

		return c.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"AuthStatus": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := models.GetUserByID(claims.Issuer)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"userID":  user.ID,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"AuthStatus": "Authenticated",
	})
}

func Logout(c *fiber.Ctx) error {
	if c.Cookies("jwt") == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "User not logged in",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	_, err = models.GetUserByID(claims.Issuer)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Next()
}

func CheckUserRole(c *fiber.Ctx) (int, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return -1, c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := models.GetUserByID(claims.Issuer)

	if err != nil {
		return -1, c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return user.Role, nil
}
