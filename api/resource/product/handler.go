package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tbm5k/tss/util/formatter"
	"gorm.io/gorm"
)

type ProductHandler struct {
    repository *ProductRepository
}

func New(db *gorm.DB) *ProductHandler {
    return &ProductHandler{
        repository: NewProductRepository(db),
    }
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {

	products, err := h.repository.List()
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	formatter.DataFormatter(w, products.toDtos(), http.StatusOK)
}

func (h *ProductHandler) Read(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }


	product, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Product not found", http.StatusNotFound)
	}

	formatter.DataFormatter(w, product.ToDto(), http.StatusCreated)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
    // decode req body to json

	fmt.Println(r.Body)

    product := &ProductDto{}
    if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		formatter.ErrorFormatter(w, "Error processing", http.StatusInternalServerError)
		return
    }

	fmt.Println(product)

	newProduct := product.ToModel()
	newProduct.ID = uuid.New()

	fmt.Println(newProduct)

	saved, err := h.repository.Create(newProduct)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
	}

	formatter.DataFormatter(w, saved, http.StatusCreated)
}

