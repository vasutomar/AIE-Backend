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
	Exam           string   `json:"exam"`
	Liked_By       []string `json:"liked_by"`
	Bookmarked_By  []string `json:"bookmarked_by"`
}

func GetDiscussions(exam string, items, page int64) (*[]bson.M, error) {
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
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return &results, err
}

func CreateDiscussion(discussion Discussion) error {
	log.Debug().Msg("CreateDiscussion started")
	_, err := providers.DB.Collection("DISCUSSIONS").InsertOne(context.Background(), discussion)
	if err != nil {
		return err
	}

	return nil
}
