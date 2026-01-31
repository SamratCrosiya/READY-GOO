package controllers

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/yourusername/magic-stream/models"
    "go.mongodb.org/mongo-driver/bson"
    "net/http"
    "time"
)

func GetAllMovies() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var movies []models.Movie
        defer cancel()

        results, err := movieCollection.Find(ctx, bson.M{})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        for results.Next(ctx) {
            var singleMovie models.Movie
            results.Decode(&singleMovie)
            movies = append(movies, singleMovie)
        }

        c.JSON(http.StatusOK, movies)
    }
}
