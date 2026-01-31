package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Genre struct {
    Name string `json:"name" bson:"name"`
}

type Ranking struct {
    RankingValue int    `json:"ranking_value" bson:"ranking_value"`
    RankingName  string `json:"ranking_name" bson:"ranking_name"`
}

type Movie struct {
    ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    ImdbID       string             `json:"imdb_id" bson:"imdb_id"`
    Title        string             `json:"title" bson:"title"`
    Genres       []Genre            `json:"genres" bson:"genres"`
    AdminReview  string             `json:"admin_review" bson:"admin_review"`
    Sentiment    string             `json:"sentiment" bson:"sentiment"`
    RankingValue int                `json:"ranking_value" bson:"ranking_value"`
}
