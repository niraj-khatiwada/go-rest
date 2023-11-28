package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/niraj-khatiwada/go-rest/api"
	"github.com/niraj-khatiwada/go-rest/routes"
	"github.com/niraj-khatiwada/go-rest/utils"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnv()

	r := chi.NewRouter()
	r.Use(utils.GetRouterMiddlewares()...)

	api.Api(r)
	routes.Routes(r)

	fmt.Printf("Server started at port %s \n", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(":9000", r); err != nil {
		fmt.Println(err)
		panic("server could not start.")

	}
}
