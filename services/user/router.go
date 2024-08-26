package user

import (
	"net/http"
	"standardApi/types"
	"standardApi/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandlerLogin).Methods("POST")
	router.HandleFunc("/register", h.HandlerRegister).Methods("POST")

}

func (h *Handler) HandlerLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if the user exists
	// if it doesnt we create the new user
}
