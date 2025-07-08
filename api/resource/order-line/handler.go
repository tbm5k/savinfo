package orderline

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tbm5k/tss/util/formatter"

	o "github.com/tbm5k/tss/api/resource/order"
	p "github.com/tbm5k/tss/api/resource/product"
)

type OrderLineHandler struct {
    repository *OrderLineRepository
	productRepo *p.ProductRepository
	orderRepo *o.OrderRepository
}

func New(db *gorm.DB) *OrderLineHandler {
    return &OrderLineHandler{
        repository: NewOrderLineRepository(db),
		productRepo: p.NewProductRepository(db),
		orderRepo: o.NewOrderRepository(db),
    }
}

func (h *OrderLineHandler) List(w http.ResponseWriter, r *http.Request) {

	orderLines, err := h.repository.List()
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	formatter.DataFormatter(w, orderLines.ToDtos(), http.StatusOK)
}

func (h *OrderLineHandler) Read(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }

	orderLine, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "OrderLine not found", http.StatusNotFound)
		return
	}

	formatter.DataFormatter(w, orderLine.ToDto(), http.StatusCreated)
}

func (h *OrderLineHandler) Create(w http.ResponseWriter, r *http.Request) {
	f := &Form{}
    if err := json.NewDecoder(r.Body).Decode(f); err != nil {
		formatter.ErrorFormatter(w, "Error processing", http.StatusInternalServerError)
		return
    }

	/*
		1. User passes the product and quantity
		2. Retrieve the product to get the price
		3. Look for a pending order
		4. If non, create a new order, else use the existing order
		5. Attach the order & product to the orderline
		6. Save
	*/

    id, err := uuid.Parse(f.ProductID)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }

	product, err := h.productRepo.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid product", http.StatusBadRequest)
		return
	}

	order, err := h.orderRepo.GetPendingOrder()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		order = &o.Order {
			ID:     uuid.New(),
			Status: "pending",
			UserID: uuid.MustParse("d3d9a613-9e82-4381-bb55-b5f31ab26cef"), // TODO: assign the correct user ID
		}
		if _, err := h.orderRepo.Create(order); err != nil {
            formatter.ErrorFormatter(w, "Failed to create order", http.StatusInternalServerError)
            return
        }
	} else if err != nil {
        formatter.ErrorFormatter(w, "Failed to retrieve order", http.StatusInternalServerError)
        return
    }

	orderLine := &OrderLineDto{}
	newOrderLine := orderLine.ToModel()
	newOrderLine.ID = uuid.New()
	newOrderLine.UnitPrice = product.Price
	newOrderLine.OrderID = order.ID
	newOrderLine.ProductID = product.ID
	newOrderLine.Quantity = f.Quantity

	saved, err := h.repository.Create(newOrderLine)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	order.Total = order.Total + (product.Price * newOrderLine.Quantity)
	err = h.orderRepo.Update(order.ID, order)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	formatter.DataFormatter(w, saved, http.StatusCreated)
}

