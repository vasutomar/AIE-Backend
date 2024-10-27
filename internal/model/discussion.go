package model

import (
	"aie/internal/providers"
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Discussion struct {
	Username       string   `json:"username"`
	Title          string   `json:"title"`
	Body           string   `json:"body"`
	Like_Count     int      `json:"like_count"`
	Bookmark_Count int      `json:"bookmark_count"`
	Comments       []string `json:"comments"`
}

func GetDiscussions(exam string, items, page int64) (*[]Discussion, error) {
	log.Debug().Msg("getDiscussions started")
	filter := bson.D{{Key: "exam", Value: exam}}

	options := options.Find()
	options.SetLimit(items)
	options.SetSkip((page - 1) * items)
	cursor, err := providers.DB.Collection("DISCUSSIONS").Find(context.TODO(), filter, options)
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
