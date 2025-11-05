package main

import (
	"fmt"
	"github.com/sxexesx/learn-backend/go/_projects/08_modules/server/internal/storage"
)

func main() {
	strg := storage.NewStorage()

	ff, err := strg.Upload("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Stored: %s", ff.Name)
}
