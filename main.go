package main

import (
	"fmt"

	"github.com/mahfuz05/go-banking-restapi/migrations"
)

func main() {
	fmt.Println("Hello")
	migrations.Migrate()
}
