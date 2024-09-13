package main

import (
	"context"
	"fmt"

	"github.com/RalphTan37/inventory-crud-api/application"
)

func main() {
	//new instance of the application
	app := application.New()

	//start the application
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
}
