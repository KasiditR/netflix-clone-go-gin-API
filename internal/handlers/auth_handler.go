package handlers

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/database"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/models"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/tokens"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"math/rand"
	"net/http"
	"time"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			return
		}

		if user.Email == nil || user.Password == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
			return
		}

		emailCount, err := database.CountDocument(bson.M{"email": user.Email}, &user)
		if err != nil || emailCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exist"})
			return
		}

		rand.NewSource(time.Now().UnixNano())
		profilePics := []string{"/avatar1.png", "/avatar2.png", "/avatar3.png"}
		randomIndex := rand.Intn(len(profilePics))
		password := utils.HashPassword(*user.Password)

		user.ID = bson.NewObjectID()
		user.Password = &password
		user.Image = &profilePics[randomIndex]
		user.SearchHistory = make([]models.SearchHistory, 0)

		token, refreshToken, err := tokens.TokenGenerator(user.ID.Hex(), *user.Email, *user.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		_, err = database.InsertOne(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "the user did not get created"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":   token,
			"refreshToken":  refreshToken,
			"id":            user.ID,
			"email":         user.Email,
			"image":         user.Image,
			"searchHistory": user.SearchHistory,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			return
		}

		if user.Email == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
			return
		}

		var foundUser models.User

		err := database.FindOne(bson.M{"email": user.Email}, &foundUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": msg})
			return
		}

		token, refreshToken, err := tokens.TokenGenerator(foundUser.ID.Hex(), *foundUser.Email, *foundUser.Image)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"accessToken":   token,
			"refreshToken":  refreshToken,
			"id":            foundUser.ID,
			"email":         foundUser.Email,
			"image":         foundUser.Image,
			"searchHistory": foundUser.SearchHistory,
		})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
	}
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("id")
		email, _ := c.Get("email")
		image, _ := c.Get("image")

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"email":    email,
			"image":    image,
		})
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		refreshToken, ok := body["refreshToken"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
			return
		}
		payload, msg := tokens.ValidateToken(refreshToken)
		if msg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		}

		token, _, err := tokens.TokenGenerator(payload.ID, payload.Email, payload.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"accessToken": token})
	}
}
