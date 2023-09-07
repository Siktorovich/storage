package main

import (
	"fmt"
	"log"

	"github.com/Siktorovich/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	getFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(st.files)

	fmt.Println("it works", getFile)
}
