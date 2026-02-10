package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go-practice/registry"
	"go-practice/usecase/todo"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "uncomplete":
		uncomplete(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func uncomplete(args []string) {
	fs := flag.NewFlagSet("uncomplete", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	id := fs.String("id", "", "todo id (required)")
	store := fs.String("store", "memory", "memory|db") // 差し替え体験用にゃ

	if err := fs.Parse(args); err != nil {
		os.Exit(2)
	}
	if *id == "" {
		fmt.Fprintln(os.Stderr, "--id is required")
		os.Exit(2)
	}

	// ここがクリーンアーキの旨味：同じusecaseを別の入口で使うにゃ
	c, err := registry.NewCLIContainer(*store)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init container: %v\n", err)
		os.Exit(1)
	}

	if err := c.Uncomplete.Execute(context.Background(), todo.UncompleteInput{ID: *id}); err != nil {
		fmt.Fprintf(os.Stderr, "uncomplete failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("ok")
}

func usage() {
	fmt.Println(`usage:
  go run ./cmd/cli uncomplete --id <todo-id> [--store memory|db]

examples:
  go run ./cmd/cli uncomplete --id 1 --store memory
  go run ./cmd/cli uncomplete --id 1 --store db
`)
}
