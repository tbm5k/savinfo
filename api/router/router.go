package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tbm5k/tss/api/resource/category"
	"github.com/tbm5k/tss/api/resource/order"
	orderline "github.com/tbm5k/tss/api/resource/order-line"
	"github.com/tbm5k/tss/api/resource/product"
	"github.com/tbm5k/tss/api/resource/user"
	"gorm.io/gorm"
)

func New(port int, db *gorm.DB) {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    })

    userHandler := user.New(db)
	categoryHandler := category.New(db)
	productHandler := product.New(db)
	orderlineHandler := orderline.New(db)
	orderHandler := order.New(db)

    r.Route("/v1", func(r chi.Router) {
        r.Get("/users/{id}", userHandler.Read)
        r.Get("/users", userHandler.List)
        r.Post("/users", userHandler.Create)

        r.Get("/categories/{id}/average", categoryHandler.AveragePrice)
        r.Get("/categories/{id}", categoryHandler.Read)
        r.Get("/categories", categoryHandler.List)
        r.Post("/categories", categoryHandler.Create)

        r.Get("/products/{id}", productHandler.Read)
        r.Get("/products", productHandler.List)
        r.Post("/products", productHandler.Create)

        r.Get("/order-lines/{id}", orderlineHandler.Read)
        r.Get("/order-lines", orderlineHandler.List)
        r.Post("/order-lines", orderlineHandler.Create)

        r.Put("/orders/{id}/process", orderHandler.Process)
        r.Get("/orders/{id}", orderHandler.Read)
        r.Get("/orders", orderHandler.List)
    })

    log.Printf("Server running on port: %v...", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
        log.Fatalf("Cannot start server!: %v", err)
    }
}

