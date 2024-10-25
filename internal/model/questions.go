package model

import (
	"aie/internal/providers"
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetQuestions() ([]*OnboardingQuestion, error) {
	log.Debug().Msg("getQuestions started")
	filter := bson.D{}

	var result bson.M
	err := providers.DB.Collection("ONBOARDING").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Debug().Msgf("Error finding user: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("No questions found")
		} else {
			return nil, err
		}
	}

	questions := result["questions"].(bson.A)
	log.Debug().Msgf("Found %d questions: %v", len(questions), questions)
	questionsList := make([]*OnboardingQuestion, 0)
	for _, question := range questions {
		questionMap := question.(bson.M)
		q := &OnboardingQuestion{
			QuestionId: questionMap["questionId"].(string),
			Title:      questionMap["title"].(string),
			Type:       questionMap["type"].(string),
			Options:    make([]string, 0),
		}
		if questionMap["options"] != nil {
			options := questionMap["options"].(bson.A)
			for _, option := range options {
				q.Options = append(q.Options, option.(string))
			}
		}
		questionsList = append(questionsList, q)
	}
	return questionsList, nil
}
