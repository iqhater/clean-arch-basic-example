package infrastructure

import (
	"testing"
)

const DB_PATH = "../../db.txt"

func TestReadDataFromFile(t *testing.T) {

	// arrange
	path := DB_PATH

	// act
	result := readDataFromFile(path)

	// assert
	if string(result) != "👋 Hello Gopher" {
		t.Errorf("Excpect: 👋 Hello Gopher Have: %s", string(result))
	}
}

func TestGetGreet(t *testing.T) {

	// arrange
	repo := NewGreetRepository(DB_PATH)
	greeting := "👋 Hello Gopher"

	// act
	result, _ := repo.GetGreet()

	// assert
	if result.Title != greeting {
		t.Errorf("Expected: %s, got: %s", greeting, result.Title)
	}

	if result == nil {
		t.Errorf("Expected: %v, got: %v", nil, result)
	}
}
