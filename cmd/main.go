package main

import (
	"fmt"
	"go-service/internal/api"
	"go-service/internal/handler"
	"go-service/internal/service"
	"net/http"
	_ "net/http/pprof"
	"time"
)


func runServer(addr string){
	config := api.Config{
		URL: "https://varankin_dev.elma365.ru/api/extensions/2a38760e-083a-4dd0-aebc-78b570bfd3c7/script/users", 
		Token: "",
	}

	api := api.NewApi(&config)
	services := service.NewService(api)
	handlers := handler.NewHandler(services)

	mux := http.NewServeMux()
	mux.HandleFunc("/get-users/", handlers.GetUsers)
	mux.HandleFunc("/", handlers.RootHandler)

	srv := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("start at " + addr)

	srv.ListenAndServe()
}

func main() {
	runServer(":8080")
}


// type UsersHandler struct{
// }
// func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	config := api.Config{
// 		URL: "https://varankin_dev.elma365.ru/api/extensions/2a38760e-083a-4dd0-aebc-78b570bfd3c7/script/users", 
// 		Token: "",
// 	}

// 	api := api.NewApi(&config)
// 	services := service.NewService(api)
// 	users := services.GetUsers()

// 	fmt.Fprintln(w, users)
// }