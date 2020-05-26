package main

import (
	"fmt"

	"github.com/mahfuz05/bankingtest/migrations"
)

func main() {
	fmt.Println("Hello")
	migrations.Migrate()
}
