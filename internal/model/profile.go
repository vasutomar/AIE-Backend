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

type Profile struct {
	UserId     string    `json:"user_id" bson:"user_id"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Exams      []string  `json:"exams"`
	Salt       string    `json:"salt"`
	Groups     []string  `json:"groups"`
	Friends    []string  `json:"friends"`
	ProfilePic string    `json:"profile_pic" bson:"profile_pic"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type CondensedUser struct {
	UserId     string `json:"user_id" bson:"user_id"`
	ProfilePic string `json:"profile_pic" bson:"profile_pic"`
	Name       string `json:"name"`
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
		UserId:     result["user_id"].(string),
		Phone:      result["phone"].(string),
		Email:      result["email"].(string),
		Exams:      make([]string, 0),
		Salt:       result["salt"].(string),
		Groups:     make([]string, 0),
		Friends:    make([]string, 0),
		ProfilePic: result["profile_pic"].(string),
		Name:       result["name"].(string),
		CreatedAt:  result["created_at"].(primitive.DateTime).Time(),
		UpdatedAt:  result["updated_at"].(primitive.DateTime).Time(),
	}

	exams := result["exams"].(bson.A)
	for _, exam := range exams {
		profile.Exams = append(profile.Exams, exam.(string))
	}

	groups := result["groups"].(bson.A)
	for _, group := range groups {
		profile.Groups = append(profile.Groups, group.(string))
	}

	friends := result["friends"].(bson.A)
	for _, friend := range friends {
		profile.Friends = append(profile.Friends, friend.(string))
	}

	return profile, nil
}

func GetAllProfilesForAnExam(exam string, itemsCount, page int64) ([]*CondensedUser, error) {
	log.Debug().Msg("GetAllProfilesForAnExam started")

	// Calculate the skip and limit values
	skip := (page - 1) * itemsCount
	limit := itemsCount

	// Define options with skip and limit for pagination
	findOptions := options.Find().SetSkip(skip).SetLimit(limit)

	filter := bson.M{"exams": exam}
	// Define a filter (in this case, an empty filter to retrieve all documents)
	//filter := bson.D{{Key: "exams", Value: bson.D{}}}

	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("PROFILE").Find(contextWithTimeout, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	profiles := make([]*CondensedUser, 0)
	for _, result := range results {
		fetchedProfile := &Profile{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, fetchedProfile)

		condensed := CondensedUser{
			Name:       fetchedProfile.Name,
			ProfilePic: fetchedProfile.ProfilePic,
			UserId:     fetchedProfile.UserId,
		}

		profiles = append(profiles, &condensed)
	}

	return profiles, nil
}

func GetUsersByName(exam, name string) ([]*CondensedUser, error) {

	// TODO - Not matching regex, need to fix

	log.Debug().Msg("GetUsersByName started")

	filter := bson.M{
		"$and": []bson.M{
			{"exams": exam},
			{"name": bson.M{"$regex": name, "$options": "i"}}, // Case-insensitive partial match
		},
	}

	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("PROFILE").Find(contextWithTimeout, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	profiles := make([]*CondensedUser, 0)
	for _, result := range results {
		fetchedProfile := &Profile{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, fetchedProfile)

		condensed := CondensedUser{
			Name:       fetchedProfile.Name,
			ProfilePic: fetchedProfile.ProfilePic,
			UserId:     fetchedProfile.UserId,
		}

		profiles = append(profiles, &condensed)
	}

	return profiles, nil
}

func GetFriends(userId string) ([]CondensedUser, error) {
	profile, _ := GetProfile(userId)
	friends := profile.Friends

	filter := bson.M{
		"user_id": bson.M{
			"$in": friends,
		},
	}
	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("PROFILE").Find(contextWithTimeout, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	fetchedFriends := make([]CondensedUser, 0)
	for _, result := range results {
		profile := &Profile{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, profile)
		friend := CondensedUser{
			Name:       profile.Name,
			ProfilePic: profile.ProfilePic,
			UserId:     profile.UserId,
		}
		fetchedFriends = append(fetchedFriends, friend)
	}

	return fetchedFriends, nil
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

	if existingProfile.ProfilePic != profile.ProfilePic && profile.ProfilePic != "" {
		existingProfile.ProfilePic = profile.ProfilePic
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

	if len(profile.Friends) > 0 && !slices.Equal(existingProfile.Friends, profile.Friends) {
		appendedFriends := append(existingProfile.Friends, profile.Friends[0])
		existingProfile.Friends = appendedFriends
		changesCount++
	}

	if changesCount == 0 {
		return nil
	}

	existingProfile.UpdatedAt = time.Now()

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

func GetCondensedProfile(userId string) (CondensedUser, error) {
	existingProfile, err := GetProfile(userId)
	if err != nil {
		return CondensedUser{}, err
	}

	condensedUser := CondensedUser{
		UserId:     existingProfile.UserId,
		ProfilePic: existingProfile.ProfilePic,
		Name:       existingProfile.Name,
	}

	return condensedUser, nil
}
