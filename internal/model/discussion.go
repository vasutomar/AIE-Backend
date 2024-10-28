package model

import (
	"aie/internal/providers"
	"context"
	"slices"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Discussion struct {
	Title         string `json:"title"`
	Body          string `json:"body"`
	Username      string `json:"username"`
	Exam          string `json:"exam"`
	LikeCount     int64  `json:"like_count" bson:"like_count"`
	BookmarkCount int64  `json:"bookmark_count" bson:"bookmark_count"`
	// Array of comment ObjectIds
	Comments []string `json:"comments"`
}

// TODO: Do field validation before calling this method
func (d *Discussion) Create() error {
	log.Debug().Msg("Create discussion started")
	_, err := providers.DB.Collection("DISCUSSIONS").InsertOne(context.Background(), d)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDiscussion(id string, discussion *Discussion) error {
	log.Debug().Msg("UpdateDiscussion started")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	existingDiscussion, err := GetDiscussion(id)
	if err != nil {
		return err
	}

	changesCount := 0
	if existingDiscussion.Title != discussion.Title && discussion.Title != "" {
		existingDiscussion.Title = discussion.Title
		changesCount++
	}
	if existingDiscussion.Body != discussion.Body && discussion.Body != "" {
		existingDiscussion.Body = discussion.Body
		changesCount++
	}
	if existingDiscussion.Exam != discussion.Exam && discussion.Exam != "" {
		existingDiscussion.Exam = discussion.Exam
		changesCount++
	}
	if existingDiscussion.LikeCount != discussion.LikeCount {
		existingDiscussion.LikeCount = discussion.LikeCount
		changesCount++
	}
	if existingDiscussion.BookmarkCount != discussion.BookmarkCount {
		existingDiscussion.BookmarkCount = discussion.BookmarkCount
		changesCount++
	}
	if len(discussion.Comments) > 0 && !slices.Equal(existingDiscussion.Comments, discussion.Comments) {
		existingDiscussion.Comments = discussion.Comments
		changesCount++
	}

	if changesCount == 0 {
		return nil
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: discussion}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("DISCUSSIONS").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}

	return nil
}

func GetDiscussion(id string) (*Discussion, error) {
	log.Debug().Msg("GetDiscussion started")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}

	var result bson.M
	err = providers.DB.Collection("DISCUSSIONS").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	discussion := &Discussion{}
	bsonBytes, _ := bson.Marshal(result)
	_ = bson.Unmarshal(bsonBytes, discussion)

	return discussion, nil
}

func GetDiscussions(exam string, count, page int64) ([]*Discussion, error) {
	log.Debug().Msg("GetDiscussions started")

	// Calculate the skip and limit values
	skip := (page - 1) * count
	limit := count

	// Define options with skip and limit for pagination
	findOptions := options.Find().SetSkip(skip).SetLimit(limit)

	// Define a filter (in this case, an empty filter to retrieve all documents)
	filter := bson.D{{Key: "exam", Value: exam}}

	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("DISCUSSIONS").Find(contextWithTimeout, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	discussions := make([]*Discussion, 0)
	for _, result := range results {
		discussion := &Discussion{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, discussion)
		discussions = append(discussions, discussion)
	}

	return discussions, nil
}
