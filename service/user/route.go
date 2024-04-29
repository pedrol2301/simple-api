package user

import (
	"net/http"
	"simple-api/types"
	"simple-api/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Register routes here
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Handle the login
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Handle the registration

	var payload types.RegisterUserPayload
	if err := utils.ParseJsonBody(r.Body, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if the user already exists

}
