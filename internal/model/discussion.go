package model

import (
	"aie/internal/providers"
	"context"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateDiscusstionRequest struct {
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Like_Count     int       `json:"like_count"`
	Bookmark_Count int       `json:"bookmark_count" bson:"bookmark_count"`
	Comments       []Comment `json:"comments"`
	Exam           string    `json:"exam"`
	Liked_By       []string  `json:"liked_by" bson:"liked_by"`
	Bookmarked_By  []string  `json:"bookmarked_by" bson:"bookmarked_by"`
}

type Comment struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
}

type Discussion struct {
	DiscussionId   string    `json:"discussion_id" bson:"discussion_id"`
	UserId         string    `json:"user_id" bson:"user_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Like_Count     int       `json:"like_count"`
	Bookmark_Count int       `json:"bookmark_count" bson:"bookmark_count"`
	Comments       []Comment `json:"comments"`
	Exam           string    `json:"exam"`
	Liked_By       []string  `json:"liked_by" bson:"liked_by"`
	Bookmarked_By  []string  `json:"bookmarked_by" bson:"bookmarked_by"`
}

type CommentRequest struct {
	Comment string `json:"comment"`
}

// TODO: Do field validation before calling this method
func (createDiscussionData *CreateDiscusstionRequest) Create(userId string) error {
	log.Debug().Msg("Create discussion started")
	d := Discussion{
		DiscussionId:   uuid.New().String(),
		UserId:         userId,
		Title:          createDiscussionData.Title,
		Body:           createDiscussionData.Body,
		Like_Count:     createDiscussionData.Like_Count,
		Bookmark_Count: createDiscussionData.Bookmark_Count,
		Comments:       createDiscussionData.Comments,
		Exam:           createDiscussionData.Exam,
		Liked_By:       createDiscussionData.Liked_By,
		Bookmarked_By:  createDiscussionData.Bookmarked_By,
	}

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
	if existingDiscussion.Like_Count != discussion.Like_Count {
		existingDiscussion.Like_Count = discussion.Like_Count
		changesCount++
	}
	if existingDiscussion.Bookmark_Count != discussion.Bookmark_Count {
		existingDiscussion.Bookmark_Count = discussion.Bookmark_Count
		changesCount++
	}
	if len(discussion.Comments) > 0 && !slices.Equal(existingDiscussion.Comments, discussion.Comments) {
		existingDiscussion.Comments = discussion.Comments
		changesCount++
	}
	if len(discussion.Liked_By) > 0 && !slices.Equal(existingDiscussion.Liked_By, discussion.Liked_By) {
		existingDiscussion.Liked_By = discussion.Liked_By
		changesCount++
	}
	if len(discussion.Bookmarked_By) > 0 && !slices.Equal(existingDiscussion.Bookmarked_By, discussion.Bookmarked_By) {
		existingDiscussion.Bookmarked_By = discussion.Bookmarked_By
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
	filter := bson.D{{Key: "discussion_id", Value: id}}

	var result bson.M
	err := providers.DB.Collection("DISCUSSIONS").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	discussion := &Discussion{}
	bsonBytes, _ := bson.Marshal(result)
	_ = bson.Unmarshal(bsonBytes, discussion)

	return discussion, nil
}

func GetDiscussionsByExam(exam string, itemsCount, page int64) ([]*Discussion, error) {
	log.Debug().Msg("GetDiscussions started")

	// Calculate the skip and limit values
	skip := (page - 1) * itemsCount
	limit := itemsCount

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

func AddComment(id string, comment Comment, userId string) error {
	discussion, err := GetDiscussion(id)
	if err != nil {
		return err
	}
	storedComments := discussion.Comments
	updatedComments := append(storedComments, comment)
	discussion.Comments = updatedComments

	filter := bson.D{{Key: "discussion_id", Value: id}}
	update := bson.D{{Key: "$set", Value: discussion}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("DISCUSSIONS").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}

	return nil
}

func ToggleLike(userId string, discussionId string) error {
	discussion, err := GetDiscussion(discussionId)
	if err != nil {
		return err
	}

	usersWhoHaveLiked := discussion.Liked_By
	matchedUserIdIndex := slices.IndexFunc(usersWhoHaveLiked, func(c string) bool { return c == userId })
	if matchedUserIdIndex >= 0 {
		// User had already liked the post, deregister like
		discussion.Like_Count = discussion.Like_Count - 1
		discussion.Liked_By = append(discussion.Liked_By[:matchedUserIdIndex], discussion.Liked_By[matchedUserIdIndex+1:]...)
	} else {
		// User has not liked the post, register like
		discussion.Like_Count = discussion.Like_Count + 1
		discussion.Liked_By = append(discussion.Liked_By, userId)
	}

	filter := bson.D{{Key: "discussion_id", Value: discussionId}}
	update := bson.D{{Key: "$set", Value: discussion}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("DISCUSSIONS").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}
	return nil
}

func ToggleBookmark(userId string, discussionId string) error {
	discussion, err := GetDiscussion(discussionId)
	if err != nil {
		return err
	}

	usersWhoHaveBookmarked := discussion.Bookmarked_By
	matchedUserIdIndex := slices.IndexFunc(usersWhoHaveBookmarked, func(c string) bool { return c == userId })
	if matchedUserIdIndex >= 0 {
		// User had already bookmarked the post, deregister bookmark
		discussion.Bookmark_Count = discussion.Bookmark_Count - 1
		discussion.Bookmarked_By = append(discussion.Bookmarked_By[:matchedUserIdIndex], discussion.Bookmarked_By[matchedUserIdIndex+1:]...)
	} else {
		// User has not bookmarked the post, register bookmark
		discussion.Bookmark_Count = discussion.Bookmark_Count + 1
		discussion.Bookmarked_By = append(discussion.Bookmarked_By, userId)
	}

	filter := bson.D{{Key: "discussion_id", Value: discussionId}}
	update := bson.D{{Key: "$set", Value: discussion}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("DISCUSSIONS").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}
	return nil
}
