package model

import (
	"aie/internal/providers"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
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
}

func GetTodoData(id string) ([]Todo, error) {
	log.Debug().Msg("GetTodo started")

	filter := bson.M{
		"user_id": id,
	}
	// Query MongoDB
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cursor, err := providers.DB.Collection("TODO").Find(contextWithTimeout, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contextWithTimeout)

	// Iterate through the documents and collect them
	var results []bson.M
	if err = cursor.All(contextWithTimeout, &results); err != nil {
		return nil, err
	}

	fetchedTodos := make([]Todo, 0)
	for _, result := range results {
		todo := &Todo{}
		bsonBytes, _ := bson.Marshal(result)
		_ = bson.Unmarshal(bsonBytes, todo)
		newTodo := Todo{
			Title:    todo.Title,
			Deadline: todo.Deadline,
			State:    todo.State,
			UserId:   todo.UserId,
			Id:       todo.Id,
		}
		fetchedTodos = append(fetchedTodos, newTodo)
	}

	return fetchedTodos, nil
}

// TODO: Do field validation before calling this method
func (todo *CreateTodoRequest) Create(userId string) error {
	log.Debug().Msg("Create Todo started")
	d := Todo{
		UserId:   userId,
		Title:    todo.Title,
		State:    "not-started",
		Deadline: todo.Deadline,
		Id:       uuid.New().String(),
	}

	_, err := providers.DB.Collection("TODO").InsertOne(context.Background(), d)
	if err != nil {
		return err
	}

	return nil
}

func GetSingleTodo(id string) (*Todo, error) {
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

func UpdateTodo(todo *Todo, todoId string) error {
	log.Debug().Msg("UpdateTodo started")

	existingTodo, err := GetSingleTodo(todoId)
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

	filter := bson.D{{Key: "todo_id", Value: todoId}}
	update := bson.D{{Key: "$set", Value: existingTodo}}
	updateOpts := options.Update().SetUpsert(false)

	_, err = providers.DB.Collection("TODO").UpdateOne(context.Background(), filter, update, updateOpts)
	if err != nil {
		return err
	}

	return nil
}
