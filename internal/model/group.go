package model

import (
	"aie/internal/providers"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Group struct {
	Group_id      string          `json:"group_id"  bson:"group_id"`
	Admin         string          `json:"admin"`
	Members       []CondensedUser `json:"members"`
	Name          string          `json:"name"`
	Group_type    string          `json:"group_type" bson:"group_type"`
	Exam          string          `json:"exam"`
	Documents     []string        `json:"documents"`
	About         string          `json:"about"`
	Group_picture string          `json:"group_pic" bson:"group_pic"`
}

type CreateGroupRequest struct {
	Members       []CondensedUser `json:"members"`
	Name          string          `json:"name"`
	Group_type    string          `json:"group_type" bson:"group_type"`
	Exam          string          `json:"exam"`
	About         string          `json:"about"`
	Group_picture string          `json:"group_pic" bson:"group_pic"`
}

func (group *CreateGroupRequest) Create(userId string) (string, error) {
	log.Debug().Msgf("Create group started: %v", group)

	// Create group
	newGroup := Group{
		Group_id:      uuid.New().String(),
		Admin:         userId,
		Name:          group.Name,
		Group_type:    group.Group_type,
		Members:       group.Members,
		Exam:          group.Exam,
		Documents:     make([]string, 0),
		About:         group.About,
		Group_picture: group.Group_picture,
	}
	log.Debug().Msgf("Creating group: %v", newGroup)
	_, err := providers.DB.Collection("GROUPS").InsertOne(context.Background(), newGroup)
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("Group created successfully: %v", newGroup)
	return newGroup.Group_id, nil
}

func GetGroups(groupIds []string) ([]*Group, error) {
	log.Debug().Msg("GetGroups started")

	// Define a filter (in this case, an empty filter to retrieve all documents)
	filter := bson.M{
		"group_id": bson.M{
			"$in": groupIds,
		},
	}
	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("GROUPS").Find(contextWithTimeout, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	fetchedGroups := make([]*Group, 0)
	for _, result := range results {
		group := &Group{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, group)
		fetchedGroups = append(fetchedGroups, group)
	}

	return fetchedGroups, nil
}
