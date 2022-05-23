package main

import (
	// "context"
	"fmt"

	"github.com/TutorialEdge/go-rest-api-course/internal/db"
)

//Run - si responsible for the instantiation and
//startup for our go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("go rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}



