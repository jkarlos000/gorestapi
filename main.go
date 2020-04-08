package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"restapi/database"
)

func main() {
	/*r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)*/
	databaseConnection := database.InitDB()

	// Logic
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)

	pachonsitos := [...]string{
		"Lirlis", "bb",
	}
	fmt.Printf("%T\n", pachonsitos)
	printString(pachonsitos[:]...)
}

func printString(x ...string) {
	for _, v := range x {
		fmt.Println(v)
	}
}
