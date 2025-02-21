package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/services"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func GetContentTrading() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/trending/%s/day?language=en-US", contentType)
		movie, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(movie.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		movieData, ok := result["results"].([]any)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid response format: 'results' not found"})
			return
		}

		rand.NewSource(time.Now().UnixNano())
		randomIndex := rand.Intn(len(movieData))

		c.JSON(http.StatusOK, gin.H{"content": movieData[randomIndex]})
	}
}

func GetContentTrailers() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/%s/%s/videos?language=en-US", contentType, id)
		trailers, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(trailers.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		trailerDatas, ok := result["results"].([]any)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid response format: 'results' not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"contents": trailerDatas})
	}
}

func GetContentDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/%s/%s?language=en-US", contentType, id)
		details, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(details.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"content": result})
	}
}

func GetContentSimilar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/%s/%s/similar?language=en-US&page=1", contentType, id)
		similar, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(similar.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		similarDatas, ok := result["results"].([]any)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid response format: 'results' not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"contents": similarDatas})
	}
}

func GetContentByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		contentType := c.Query("contentType")
		url := fmt.Sprintf("https://api.themoviedb.org/3/%s/%s?language=en-US&page=1", contentType, category)
		content, err := services.FetchFromTMDB(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var result map[string]any

		if err = json.Unmarshal(content.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		contentDatas, ok := result["results"].([]any)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid response format: 'results' not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"contents": contentDatas})
	}
}
