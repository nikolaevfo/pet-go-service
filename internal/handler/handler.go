package handler

import (
	"fmt"
	"github.com/nikolaevfo/pet-go-service/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

type Post struct{
	message string
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request){
	resp := h.service.GetUsers()

	fmt.Fprintln(w, resp)
}

func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request){
	// res := make(chan Post)
	// go getPost(res)
	fmt.Fprintln(w, "Server is working")
}

func getPost(out chan Post){
	post := Post{message: "Server is working"}
	out <- post
}




// type UsersHandler struct{
// 	service *service.Service
// }

// type RootHandler struct{
// 	service *service.Service
// }

// func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	resp := h.service.GetUsers()

// 	fmt.Fprintln(w, resp)
// }

// func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "Server is working")
// }

// func NewRootHandler(service *service.Service) *RootHandler{
// 	return &RootHandler{
// 		service: service,
// 	}
// }

// func NewUsersHandler(service *service.Service) *UsersHandler{
// 	return &UsersHandler{
// 		service: service,
// 	}
// }