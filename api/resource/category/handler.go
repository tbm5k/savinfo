package category

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tbm5k/tss/util/formatter"
	"gorm.io/gorm"
)

type CategoryHandler struct {
    repository *CategoryRepository
}

func New(db *gorm.DB) *CategoryHandler {
    return &CategoryHandler{
        repository: NewCategoryRepository(db),
    }
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {

	categorys, err := h.repository.List()
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	formatter.DataFormatter(w, categorys.ToDtos(), http.StatusOK)
}

func (h *CategoryHandler) AveragePrice(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }

	category, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Category not found", http.StatusNotFound)
	}

	var total uint
	for _, c := range category.Products {
		total += c.Price
	}

	log.Println(total, len(category.Products))
	avg := float32(total) / float32(len(category.Products))

	formatter.DataFormatter(w, avg, http.StatusCreated)
}

func (h *CategoryHandler) Read(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }

	category, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Category not found", http.StatusNotFound)
	}

	formatter.DataFormatter(w, category.ToDto(), http.StatusCreated)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
    // decode req body to json
    category := &CategoryDto{}
    if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		formatter.ErrorFormatter(w, "Error processing", http.StatusInternalServerError)
    }

	newCategory := category.ToModel()
	newCategory.ID = uuid.New()

	saved, err := h.repository.Create(newCategory)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
	}

	formatter.DataFormatter(w, saved, http.StatusCreated)
}

