package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/ooonkeet/api-dev/e-commApi/internal/adapters/postgresql/sqlc"
	"github.com/ooonkeet/api-dev/e-commApi/internal/orders"
	"github.com/ooonkeet/api-dev/e-commApi/internal/products"
)

// import "honnef.co/go/tools/config"

// run
// mount
func (app *application) mount() http.Handler{
	r := chi.NewRouter()

  // A good base middleware stack
  r.Use(middleware.RequestID) //important for rate limiting
  r.Use(middleware.RealIP) //important for rate limiting, analytics and tracking
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer) //recovery from crashes

  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("all good"))
  })
  productService:=products.NewService(repo.New(app.db))
  productsHandler:=products.NewHandler(productService)
  r.Get("/products",productsHandler.ListProducts)
  orderService:=orders.NewService(repo.New(app.db),app.db)
  ordersHandler:=orders.NewHandler(orderService)
  r.Post("/orders",ordersHandler.PlaceOrder)
	// http.ListenAndServe(":3333",r)
	return r
}

func (app *application) run(h http.Handler) error{
	srv:= &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}
	log.Printf("server has started at addr %s",app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
	db *pgx.Conn
}


type config struct{
	addr string
	db dbConfig
}

type dbConfig struct{
	dsn string
}