package model

import (
	"aie/internal/providers"
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Member struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Group struct {
	Group_id    string   `json:"group_id"`
	Admin       string   `json:"admin"`
	Members     []Member `json:"members"`
	MemberCount int      `json:"member_count"`
	Name        string   `json:"name"`
	Color       string   `json:"color"`
	Group_type  string   `json:"group_type"`
	Exam        string   `json:"exam"`
	Documents   []string `json:"documents"`
}

type CreateGroupRequest struct {
	Members     []Member `json:"members"`
	Name        string   `json:"name"`
	Color       string   `json:"color"`
	Group_type  string   `json:"group_type"`
	Exam        string   `json:"exam"`
	MemberCount int      `json:"member_count"`
}

func (group *CreateGroupRequest) Create(userId string) (string, error) {
	log.Debug().Msgf("Create group started: %v", group)

	// Create group
	newGroup := Group{
		Group_id:    uuid.New().String(),
		Admin:       userId,
		Name:        group.Name,
		Color:       group.Color,
		Group_type:  group.Group_type,
		Members:     group.Members,
		Exam:        group.Exam,
		Documents:   make([]string, 0),
		MemberCount: group.MemberCount,
	}
	log.Debug().Msgf("Creating group: %v", newGroup)
	_, err := providers.DB.Collection("GROUPS").InsertOne(context.Background(), newGroup)
	if err != nil {
		return "", err
	}

	log.Debug().Msgf("Group created successfully: %v", newGroup)
	return newGroup.Group_id, nil
}
