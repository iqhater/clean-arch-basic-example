package infrastructure

import (
	"io"
	"log"
	"os"
)

// static check interface implementation
var _ GreeterRepo = (*GreetFileDB)(nil)

// repository
type GreeterRepo interface {
	GetGreet() *UserGreetPublic
}

// db implementation
type GreetFileDB struct {
	path string
}

// init
func NewGreetRepository(db string) *GreetFileDB {
	return &GreetFileDB{
		path: db,
	}
}

// get from file/db implementation
func (db *GreetFileDB) GetGreet() *UserGreetPublic {
	return &UserGreetPublic{
		Title: string(readDataFromFile(db.path)),
	}
}

func readDataFromFile(path string) []byte {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
