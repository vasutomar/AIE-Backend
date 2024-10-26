package model

import (
	"aie/internal/providers"
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Discussion struct {
	Username       string   `json:"username"`
	Title          string   `json:"title"`
	Body           string   `json:"body"`
	Like_Count     int      `json:"like_count"`
	Bookmark_Count int      `json:"bookmark_count"`
	Comments       []string `json:"comments"`
}

func GetDiscussions(exam string) (*[]Discussion, error) {
	log.Debug().Msg("getDiscussions started")
	filter := bson.D{{Key: "exam", Value: exam}}

	cursor, err := providers.DB.Collection("DISCUSSIONS").Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// Unpacks the cursor into a slice
	var results []Discussion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return &results, err
}
