package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	mygopackage "github.com/ok200paul/my-go-package"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	client, err := mygopackage.New(mygopackage.WithLogger(logger))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println(client.DoSomething())
}
