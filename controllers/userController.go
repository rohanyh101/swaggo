package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roh4nyh/swaggo/database"
	helper "github.com/roh4nyh/swaggo/helpers"
	"github.com/roh4nyh/swaggo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	userdatabaseName   = "Cluster0"
	userCollectionName = "users"
)

var userValidate = validator.New()
var UserCollection *mongo.Collection = database.OpenCollection(userdatabaseName, userCollectionName)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword, foundUserPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintln("email or password is incorrect")
		check = false
	}

	return check, msg
}

// UserSignUp godoc
// @Summary user sign up
// @Schemes
// @Description sign up a user
// @Accept json
// @Produce json
// @SecurityDefinitions ApiKeyAuth
// @Param user body models.User true "User sign up request"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {string} map[string]interface{} "error": "email or password is incorrect"
// @Failure 500 {string} map[string]interface{} "error": "Error occurred while checking for email"
// @Router /auth/signup [post]
func UserSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := userValidate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for email"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()

		token, _ := helper.GenerateToken(user.UserId, *user.Firstname, *user.Lastname, *user.Email, *user.Role)
		user.Token = &token

		resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintln("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

// UserLogIn godoc
// @Summary user log in
// @Schemes
// @Description log in a user
// @Accept json
// @Produce json
// @SecurityDefinitions ApiKeyAuth
// @Param user body models.User true "User log in request"
// @Success 200 {string} models.User
// @Failure 400 {string} map[string]interface{} "error": "email or password is incorrect"
// @Failure 500 {string} map[string]interface{} "error": "email or password is incorrect"
// @Router /auth/login [post]
func UserLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		token, err := helper.GenerateToken(foundUser.UserId, *foundUser.Firstname, *foundUser.Lastname, *foundUser.Email, *foundUser.Role)
		if err != nil || token == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		helper.UpdateToken(token, foundUser.UserId)

		err = UserCollection.FindOne(ctx, bson.M{"user_id": foundUser.UserId}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)
	}
}

// GetUsers godoc
// @Summary get all users (ADMIN privilege required)
// @Schemes
// @Description get all users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} []models.User "List of users"
// @Failure 400 {string} map[string]interface{} "error": "UnAuthenticated to access this resource"
// @Failure 500 {string} map[string]interface{} "error": "Error occurred while listing users"
// @Failure 200 {string} map[string]interface{} "error": "no users available"
// @Router /users [get]
// @Security ApiKeyAuth
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var users []models.User

		cursor, err := UserCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing users"})
			return
		}

		if err = cursor.All(ctx, &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
			return
		}

		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no users available"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// GetUser godoc
// @Summary get a user profile
// @Schemes
// @Description get a user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} models.User "User profile"
// @Failure 400 {string} map[string]interface{} "error": "UnAuthenticated to access this resource"
// @Failure 500 {string} map[string]interface{} "error": "Error occurred while fetching user"
// @Failure 404 {string} map[string]interface{} "error": "User not found"
// @Router /profile [get]
// @Security ApiKeyAuth
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")

		if err := helper.MatchUserTypeToId(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		err := UserCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser godoc
// @Summary update a user profile
// @Schemes
// @Description update a user
// @Accept json
// @Produce json
// @Param user body models.User true "User update request"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {string} map[string]interface{} "error": "UnAuthenticated to access this resource"
// @Failure 500 {string} map[string]interface{} "error": "Error occurred while updating user"
// @Failure 404 {string} map[string]interface{} "error": "User not found"
// @Router /profile [put]
// @Security ApiKeyAuth
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")

		if err := helper.MatchUserTypeToId(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateObj := bson.M{}

		if user.Firstname != nil {
			updateObj["first_name"] = user.Firstname
		}

		if user.Lastname != nil {
			updateObj["last_name"] = user.Lastname
		}

		if user.Password != nil {
			password := HashPassword(*user.Password)
			updateObj["password"] = password
		}

		updateObj["updated_at"] = time.Now()

		filter := bson.M{"user_id": bson.M{"$eq": userId}}
		update := bson.M{"$set": updateObj}

		_, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}

// DeleteUser godoc
// @Summary delete a user profile
// @Schemes
// @Description delete a user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} map[string]interface{} "error": "UnAuthenticated to access this resource"
// @Failure 404 {string} map[string]interface{} "error": "User not found"
// @Failure 500 {string} map[string]interface{} "error": "Error occurred while deleting user"
// @Router /profile [delete]
// @Security ApiKeyAuth
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")

		if err := helper.MatchUserTypeToId(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := UserCollection.DeleteOne(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting user"})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
