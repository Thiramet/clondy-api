package handler

import (
	"clondy/model"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Signup ...
func (h *Handler) Signup(c echo.Context) (err error) {

	// Bind
	u := &model.User{ID: primitive.NewObjectID()}
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No Content email or password"}
	}

	fmt.Println(u.Email)

	// Check username
	user, _ := h.Collector.GetUser(bson.M{"email": u.Email})
	if user != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "User already exists "}
	}

	// GenerateFromPasswords
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error GenerateFromPassword"}
	}
	u.Password = string(hash)
	u.CreatedAt = time.Now()

	// Save users
	io, err := h.Collector.InsertUser(u)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Import data has a problem"}
	}

	return c.JSON(http.StatusOK, io)

}

// Login ..
func (h *Handler) Login(c echo.Context) (err error) {

	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No Content email or password"}
	}

	// Find Email
	user, err := h.Collector.GetUser(bson.M{"email": u.Email})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Username"}
	}

	//CompareHashAndPassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims //payload
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign and get the complete encoded token as a string using the secret //header
	tokenString, err := token.SignedString([]byte(os.Getenv("K_JWT")))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Error while generating token,Try again"}
	}

	user.Token = tokenString
	user.Password = ""
	user.Email = ""
	user.Username = ""

	return c.JSON(http.StatusOK, user)

}

// GetUser ..
func (h *Handler) GetUser(c echo.Context) (err error) {

	userID := c.Param("id")
	ObjectID, _ := primitive.ObjectIDFromHex(userID)

	fmt.Println(ObjectID)

	user, _ := h.Collector.GetUser(bson.M{"_id": ObjectID})
	user.Password = ""
	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}

// GetUsers ..
func (h *Handler) GetUsers(c echo.Context) (err error) {

	// userID := c.Param("id")
	// ObjectID, _ := primitive.ObjectIDFromHex(userID)

	// fmt.Println(ObjectID)

	user, _ := h.Collector.GetUser(bson.M{})
	user.Password = ""
	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}
