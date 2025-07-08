package order

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joeshaw/envdecode"
	"github.com/tbm5k/tss/util/formatter"
	"gorm.io/gorm"

	m "github.com/tbm5k/tss/util/mailer"
	s "github.com/tbm5k/tss/util/sms"
)


type OrderHandler struct {
    repository *OrderRepository
	mailer *m.Mailer
	sms *s.Sms
}

func New(db *gorm.DB) *OrderHandler {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatal("Cannot decode env")
	}
    return &OrderHandler{
        repository: NewOrderRepository(db),
		sms: s.NewSmser(),
		mailer: m.NewMailer(
			cfg.User,
			cfg.Pass,
			cfg.User,
		),
    }
}

type Config struct {
	User string `env:"SMTP_USER"`
	Pass string `env:"SMTP_PASS"`
}

func (h *OrderHandler) Process(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.repository.Update(id, &Order{Status: "processing"})
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	o, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Could not finish executing order processing request", http.StatusInternalServerError)
		return
	}

	// email to admin
	msg := fmt.Sprintf("Order by %v\nContains: %v\nTotal: %v", o.UserID, 0, o.Total)

	em := &m.Email{
		To: "tedburg5@gmail.com",
		Subject: "Order processing brief",
		Body: msg,
	}

	if err := h.mailer.Send(em); err != nil {
		log.Println("Order is processing but email not sent")
		formatter.MessageFormatter(w, "Order completed", http.StatusOK)
		return
	}

	// sms to customer
	if err = h.sms.Send(fmt.Sprintf("You have placed an order"), ""); err != nil {
		log.Println("Order is processing but sms not sent")
		formatter.MessageFormatter(w, "Order completed", http.StatusOK)
		return
	}

	formatter.MessageFormatter(w, "Order completed", http.StatusOK)
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {

	orders, err := h.repository.List()
	if err != nil {
		formatter.ErrorFormatter(w, "Server error", http.StatusInternalServerError)
		return
	}

	formatter.DataFormatter(w, orders.ToDtos(), http.StatusOK)
}

func (h *OrderHandler) Read(w http.ResponseWriter, r *http.Request) {

    idStr := chi.URLParam(r, "id")
    id, err := uuid.Parse(idStr)
    if err != nil {
		formatter.ErrorFormatter(w, "Invalid id", http.StatusBadRequest)
		return
    }


	order, err := h.repository.Read(id)
	if err != nil {
		formatter.ErrorFormatter(w, "Order not found", http.StatusNotFound)
		return
	}

	formatter.DataFormatter(w, order.ToDto(), http.StatusCreated)
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    // decode req body to json
    order := &OrderDto{}
    if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		formatter.ErrorFormatter(w, "Error processing", http.StatusInternalServerError)
		return
    }

	newOrder := order.ToModel()
	newOrder.ID = uuid.New()

	saved, err := h.repository.Create(newOrder)
	if err != nil {
		formatter.ErrorFormatter(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	formatter.DataFormatter(w, saved, http.StatusCreated)
}

