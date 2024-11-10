package model

import (
	"aie/internal/providers"
	"context"
	"slices"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Profile struct {
	UserId string   `json:"user_id" bson:"user_id"`
	Phone  string   `json:"phone"`
	Email  string   `json:"email"`
	Exams  []string `json:"exams"`
	Salt   string   `json:"salt"`
	Groups []string `json:"groups"`
}

func CreateProfile(profile Profile) error {
	log.Debug().Msgf("creating profile for user: %v", profile.UserId)
	_, err := providers.DB.Collection("PROFILE").InsertOne(context.Background(), profile)
	if err != nil {
		return err
	}
	return nil
}

func GetProfile(userId string) (*Profile, error) {
	log.Debug().Msgf("profile get started: %v", userId)
	// Check if the user already exists
	filter := bson.D{{Key: "user_id", Value: userId}}

	var result bson.M
	err := providers.DB.Collection("PROFILE").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	profile := &Profile{
		UserId: result["user_id"].(string),
		Phone:  result["phone"].(string),
		Email:  result["email"].(string),
		Exams:  make([]string, 0),
		Salt:   result["salt"].(string),
		Groups: make([]string, 0),
	}

	exams := result["exams"].(bson.A)
	for _, exam := range exams {
		profile.Exams = append(profile.Exams, exam.(string))
	}

	return profile, nil
}

// Note: Modifyable fields: Phone, Email, Exams
// + add validations using regex for these fields before updating them.
func UpdateProfile(userId string, profile *Profile) error {
	existingProfile, err := GetProfile(userId)
	if err != nil {
		return err
	}

	changesCount := 0

	if existingProfile.Phone != profile.Phone && profile.Phone != "" {
		existingProfile.Phone = profile.Phone
		changesCount++
	}

	if existingProfile.Email != profile.Email && profile.Email != "" {
		existingProfile.Email = profile.Email
		changesCount++
	}

	if len(profile.Exams) > 0 && !slices.Equal(existingProfile.Exams, profile.Exams) {
		existingProfile.Exams = profile.Exams
		changesCount++
	}

	if len(profile.Groups) > 0 && !slices.Equal(existingProfile.Groups, profile.Groups) {
		appendedGroups := append(existingProfile.Groups, profile.Groups[0])
		existingProfile.Groups = appendedGroups
		changesCount++
	}

	if changesCount == 0 {
		return nil
	}

	// existingProfile is now updated
	updatedProfile := existingProfile
	log.Debug().Msgf("profile update started: %v", userId)
	filter := bson.D{{Key: "user_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: updatedProfile}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("PROFILE").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}

	return nil
}
