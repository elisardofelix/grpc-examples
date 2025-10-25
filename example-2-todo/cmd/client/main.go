package main

import (
	"context"
	"log"

	"github.com/elisardofelix/grpc-examples/example-2-todo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "buy groceries"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task1.GetId())

	task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "walk the dog"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task2.GetId())

	// List tasks
	tasks, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatalf("could not list tasks: %v", err)
	}
	log.Println("Current tasks:")
	for _, task := range tasks.GetTasks() {
		log.Printf("- [%s] %s", task.Id, task.Task)
	}

	// Complete a task
	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: task1.GetId()})
	if err != nil {
		log.Fatalf("could not complete task: %v", err)
	}
	log.Printf("Completed task with ID: %s", task1.GetId())

	// Add another task
	task3, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "have breakfast"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added task - id: %s", task3.GetId())

	// List tasks
	tasks, err = client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatalf("could not list tasks: %v", err)
	}
	log.Println("Current tasks:")
	for _, task := range tasks.GetTasks() {
		log.Printf("- [%s] %s", task.Id, task.Task)
	}

}
