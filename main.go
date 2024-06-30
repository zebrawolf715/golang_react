package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"` //mongoDBでbsonを使う
	Completed bool   `json:"completed" `
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("hello world")

	//envファイル
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loding .env file:", err)
	}

	//データベース関連
	MONGODB_URI := os.Getenv("MONGODB_URI") //データベースURI
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", grtTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

}
