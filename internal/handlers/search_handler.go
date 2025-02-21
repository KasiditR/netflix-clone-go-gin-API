package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/database"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/models"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/services"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SearchContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/search/%s?query=%s&include_adult=false&language=en-US&page=1", strings.ToLower(contentType), name)
		search, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(search.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		searchDatas, ok := result["results"].([]any)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid response format: 'results' not found"})
			return
		}

		if len(searchDatas) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("not found %s", name)})
			return
		}

		idVal, exists := utils.GeLocalValue(c, "id")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing user ID"})
			return
		}

		uid, err := bson.ObjectIDFromHex(idVal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID format"})
			return
		}

		var searchHistory models.SearchHistory
		if res, ok := searchDatas[0].(map[string]any); ok {
			if profilePath, ok := res["profile_path"].(string); ok && profilePath != "" {
				searchHistory.Image = &profilePath
			} else if knownFor, ok := res["known_for"].([]any); ok && len(knownFor) > 0 {
				if firstKnown, ok := knownFor[0].(map[string]any); ok {
					if posterPath, ok := firstKnown["poster_path"].(string); ok {
						searchHistory.Image = &posterPath
					}
				}
			} else if posterPath, ok := res["poster_path"].(string); ok && posterPath != "" {
				searchHistory.Image = &posterPath
			}

			if id, ok := res["id"].(float64); ok {
				searchHistory.ID = int(id)
			}

			if name, ok := res["name"].(string); ok {
				searchHistory.Title = &name
			} else if title, ok := res["title"].(string); ok {
				searchHistory.Title = &title
			}

			searchHistory.SearchType = &contentType

			dt := bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond))
			searchHistory.CreatedAt = &dt
		}

		filter := bson.M{"_id": uid}
		update := bson.M{"$push": bson.M{"searchHistory": searchHistory}}
		err = database.FindByIDAndUpdate("users", filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"results": searchDatas})
	}
}

func GetSearchHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		idVal, exists := utils.GeLocalValue(c, "id")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing user ID"})
			return
		}

		var user models.User
		err := database.FindByID(idVal, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"history": user.SearchHistory})
	}
}

func RemoveItemFromSearchHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		idVal := c.Param("id")
		id, err := strconv.Atoi(idVal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		idVal, exists := utils.GeLocalValue(c, "id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing user ID"})
			return
		}

		uid, err := bson.ObjectIDFromHex(idVal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID format"})
			return
		}

		err = database.FindByIDAndUpdate("users", bson.M{"_id": uid}, bson.M{
			"$pull": bson.M{
				"searchHistory": bson.M{
					"id": id,
				},
			},
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "item removed for search history"})
	}
}

func ClearSearchHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		idVal, exists := utils.GeLocalValue(c, "id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing user ID"})
			return
		}

		uid, err := bson.ObjectIDFromHex(idVal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID format"})
			return
		}

		err = database.FindByIDAndUpdate("users", bson.M{"_id": uid}, bson.M{
			"$set": bson.M{
				"searchHistory": []interface{}{},
			},
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "clear search history successfully"})
	}
}
