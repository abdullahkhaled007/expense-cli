package cmd

import (
	"errors"
	"fmt"
	"golang/internal/expense"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("usage: expense <command> [args]")
		return
	}
	command := strings.ToLower(os.Args[1])
	args := os.Args[2:]

	// Use a simple JSON-backed persistent store in the project directory
	svc := expense.NewServiceWithFile("expenses.json")

	var err error
	switch command {
	case "list", "ls":
		err = ListCmd(args, svc)
	case "add":
		err = AddCmd(args, svc)
	case "delete", "del", "rm":
		err = DeleteCmd(args, svc)
	case "help", "h":
		printHelp()
		return
	default:
		fmt.Printf("unknown command: %s\n", command)
		fmt.Println("available commands: Add, List, Delete")
		return
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func ListCmd(args []string, svc *expense.Service) error {
	items, err := svc.List()
	if err != nil {
		return err
	}

	for _, e := range items {
		fmt.Printf("%d\t%s\t%.2f\t%s\n",
			e.ID, e.Name, e.Price, e.Date.Format("2006-01-02"),
		)
	}

	return nil
}

func AddCmd(args []string, svc *expense.Service) error {
	if len(args) < 2 {
		return errors.New("usage: add <name> <price> [description]")
	}

	name := args[0]
	if name == "" {
		return errors.New("name must not be empty")
	}

	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("invalid price: %w", err)
	}
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}

	desc := ""
	if len(args) > 2 {
		desc = args[2]
	}

	e, err := svc.Add(name, price, desc)
	if err != nil {
		return err
	}

	fmt.Printf("Added expense: %d %s %.2f\n", e.ID, e.Name, e.Price)

	return nil
}

func DeleteCmd(args []string, svc *expense.Service) error {
	if len(args) < 1 {
		return errors.New("usage: delete <id>")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	deleted, err := svc.Delete(id)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted expense: %d %s %.2f\n", deleted.ID, deleted.Name, deleted.Price)

	return nil
}

func printHelp() {
	fmt.Println("Usage: expense <command> [args]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list, ls           List all expenses")
	fmt.Println("  add <name> <price> Add a new expense")
	fmt.Println("  delete <id>        Delete an expense")
}
