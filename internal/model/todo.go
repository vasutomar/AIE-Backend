package model

import (
	"aie/internal/providers"
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	Id       string `json:"todo_id" bson:"todo_id"`
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	UserId   string `json:"user_id" bson:"user_id"`
	State    string `json:"state"`
}

type CreateTodoRequest struct {
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	State    string `json:"state"`
}

func GetTodoData(id string) (*Todo, error) {
	log.Debug().Msg("GetTodo started")
	filter := bson.D{{Key: "todo_id", Value: id}}

	var result bson.M
	err := providers.DB.Collection("TODO").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	todo := &Todo{}
	bsonBytes, _ := bson.Marshal(result)
	_ = bson.Unmarshal(bsonBytes, todo)

	return todo, nil
}

// TODO: Do field validation before calling this method
func (todo *CreateTodoRequest) Create(userId string) error {
	log.Debug().Msg("Create Todo started")
	d := Todo{
		UserId:   userId,
		Title:    todo.Title,
		State:    todo.State,
		Deadline: todo.Deadline,
		Id:       uuid.New().String(),
	}

	_, err := providers.DB.Collection("TODO").InsertOne(context.Background(), d)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodo(todo *Todo, todoId string) error {
	log.Debug().Msg("UpdateTodo started")
	objectId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return err
	}

	existingTodo, err := GetTodoData(todoId)
	if err != nil {
		return err
	}

	changesCount := 0
	if existingTodo.Title != todo.Title && todo.Title != "" {
		existingTodo.Title = todo.Title
		changesCount++
	}
	if existingTodo.Deadline != todo.Deadline && todo.Deadline != "" {
		existingTodo.Deadline = todo.Deadline
		changesCount++
	}
	if existingTodo.State != todo.State && todo.State != "" {
		existingTodo.State = todo.State
		changesCount++
	}

	if changesCount == 0 {
		return nil
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: todo}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("DISCUSSIONS").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}

	return nil
}
