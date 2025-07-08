package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tbm5k/tss/util/formatter"
	"gorm.io/gorm"
)

type UserHandler struct {
    repository *UserRepository
}

func New(db *gorm.DB) *UserHandler {
    return &UserHandler{
        repository: NewUserRepository(db),
    }
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {

	users, err := h.repository.List()
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	formatter.DataFormatter(w, users.ToDtos(), http.StatusOK)
}

func (h *UserHandler) Read(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }


	user, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "User not found", http.StatusNotFound)
	}

	formatter.DataFormatter(w, user.ToDto(), http.StatusCreated)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    // decode req body to json
    user := &UserDto{}
    if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		formatter.ErrorFormatter(w, "Error processing", http.StatusInternalServerError)
    }

	newUser := user.ToModel()
	newUser.ID = uuid.New()

	saved, err := h.repository.Create(newUser)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
	}

	formatter.DataFormatter(w, saved, http.StatusCreated)
}

