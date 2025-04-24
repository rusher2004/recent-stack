package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rusher2004/recent-stack/stack"
)

type DocOpen struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type DocNav struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type DocSave struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type SharePerson struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type ShareIn struct {
	ID        string      `json:"id"`
	From      SharePerson `json:"from"`
	Timestamp time.Time   `json:"timestamp"`
}

type ShareOut struct {
	ID         string        `json:"id"`
	Recipients []SharePerson `json:"recipients"`
	Timestamp  time.Time     `json:"timestamp"`
}

func main() {
	// Demonstrate DocOpen stack
	fmt.Println("=== DocOpen Stack Demo ===")
	demoDocOpenStack()
	fmt.Println()

	// Demonstrate DocNav stack
	fmt.Println("=== DocNav Stack Demo ===")
	demoDocNavStack()
	fmt.Println()

	// Demonstrate DocSave stack
	fmt.Println("=== DocSave Stack Demo ===")
	demoDocSaveStack()
	fmt.Println()

	// Demonstrate SharePerson stack
	fmt.Println("=== SharePerson Stack Demo ===")
	demoSharePersonStack()
	fmt.Println()

	// Demonstrate ShareIn stack
	fmt.Println("=== ShareIn Stack Demo ===")
	demoShareInStack()
	fmt.Println()

	// Demonstrate ShareOut stack
	fmt.Println("=== ShareOut Stack Demo ===")
	demoShareOutStack()
	fmt.Println()
}

// Read items from a JSON file
func readItemsFromJson[T any](filename string) ([]T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return items, nil
}

// Push items to a stack and demonstrate stack operations
func demoStackOperations[T any](s *stack.Stack[T], items []T, itemType string) {
	// Push items to stack
	fmt.Printf("Pushing %d %s items to stack\n", len(items), itemType)
	for _, item := range items {
		s.Push(stack.Item[T]{
			ID:   fmt.Sprintf("%s-%s", itemType, time.Now().Format("20060102-150405.000")),
			Type: itemType,
			Data: item,
		})
		// Add a small delay to ensure unique IDs
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("Stack size: %d\n", s.Size())

	// Peek at top item
	if topItem, ok := s.Peek(); ok {
		fmt.Printf("Top item ID: %s, Type: %s, Data: %v\n", topItem.ID, topItem.Type, topItem.Data)
	} else {
		fmt.Println("Stack is empty, cannot peek")
	}

	// Pop a few items
	fmt.Println("Popping 3 items from stack:")
	for i := 0; i < 3 && !s.IsEmpty(); i++ {
		if item, ok := s.Pop(); ok {
			fmt.Printf("Popped item ID: %s, Type: %s, Data: %v\n", item.ID, item.Type, item.Data)
		}
	}

	fmt.Printf("Stack size after popping: %d\n", s.Size())
}

func demoDocOpenStack() {
	items, err := readItemsFromJson[DocOpen]("data/DocOpen.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	docOpenStack := stack.New[DocOpen]()
	demoStackOperations(docOpenStack, items, "DocOpen")
}

func demoDocNavStack() {
	items, err := readItemsFromJson[DocNav]("data/DocNav.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	docNavStack := stack.New[DocNav]()
	demoStackOperations(docNavStack, items, "DocNav")
}

func demoDocSaveStack() {
	items, err := readItemsFromJson[DocSave]("data/DocSave.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	docSaveStack := stack.New[DocSave]()
	demoStackOperations(docSaveStack, items, "DocSave")
}

func demoSharePersonStack() {
	items, err := readItemsFromJson[SharePerson]("data/SharePerson.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	sharePersonStack := stack.New[SharePerson]()
	demoStackOperations(sharePersonStack, items, "SharePerson")
}

func demoShareInStack() {
	items, err := readItemsFromJson[ShareIn]("data/ShareIn.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	shareInStack := stack.New[ShareIn]()
	demoStackOperations(shareInStack, items, "ShareIn")
}

func demoShareOutStack() {
	items, err := readItemsFromJson[ShareOut]("data/ShareOut.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	shareOutStack := stack.New[ShareOut]()
	demoStackOperations(shareOutStack, items, "ShareOut")
}
