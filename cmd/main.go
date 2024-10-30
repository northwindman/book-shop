package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"

	"github.com/northwindman/book-shop/internal/app/config"
	"github.com/northwindman/book-shop/internal/app/repository/pgrepo"
	"github.com/northwindman/book-shop/internal/app/services"
	"github.com/northwindman/book-shop/internal/app/transport/httpserver"
	"github.com/northwindman/book-shop/internal/pkg/pg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

const tokenTTL = time.Minute * 5

func run() error {
	// read config from env
	cfg := config.Read()

	pgDB, err := pg.Dial(cfg.DSN)
	if err != nil {
		return fmt.Errorf("pg.Dial failed: %w", err)
	}

	// run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations")
		if err := runPgMigrations(cfg.DSN, cfg.MigrationsPath); err != nil {
			return fmt.Errorf("runPgMigrations failed: %w", err)
		}
	}

	// create repositories
	userRepo := pgrepo.NewUserRepo(pgDB)
	bookRepo := pgrepo.NewBookRepo(pgDB)
	categoryRepo := pgrepo.NewCategoryRepo(pgDB)
	cartRepo := pgrepo.NewCartRepo(pgDB)

	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	tokenService := services.NewTokenService(tokenTTL)
	cartService := services.NewCartService(cartRepo)

	// create http server with application injected
	httpServer := httpserver.NewHttpServer(userService, tokenService, bookService, categoryService, cartService)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("book-shop API v0.1"))
	}).Methods("GET")

	router.HandleFunc("/signup", httpServer.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/signin", httpServer.SignIn).Methods(http.MethodPost)

	router.HandleFunc("/books", httpServer.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{book_id}", httpServer.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book", httpServer.CheckAdmin(httpServer.CreateBook)).Methods(http.MethodPost)
	router.HandleFunc("/book/{book_id}", httpServer.CheckAdmin(httpServer.UpdateBook)).Methods(http.MethodPatch)
	router.HandleFunc("/book/{book_id}", httpServer.CheckAdmin(httpServer.DeleteBook)).Methods(http.MethodDelete)

	router.HandleFunc("/categories", httpServer.GetCategories).Methods(http.MethodGet)
	router.HandleFunc("/category/{category_id}", httpServer.GetCategory).Methods(http.MethodGet)
	router.HandleFunc("/category", httpServer.CheckAdmin(httpServer.CreateCategory)).Methods(http.MethodPost)
	router.HandleFunc("/category/{category_id}", httpServer.CheckAdmin(httpServer.UpdateCategory)).Methods(http.MethodPatch)
	router.HandleFunc("/category/{category_id}", httpServer.CheckAdmin(httpServer.DeleteCategory)).Methods(http.MethodDelete)

	router.HandleFunc("/cart", httpServer.CheckAuthorizedUser(httpServer.UpdateCart)).Methods(http.MethodPost)
	router.HandleFunc("/checkout", httpServer.CheckAuthorizedUser(httpServer.Checkout)).Methods(http.MethodPost)

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("Cleaning expired carts")
				err := cartRepo.CleanExpiredCarts(ctx, time.Minute)
				if err != nil {
					log.Printf("cartRepo.CleanExpiredCarts failed: %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}(context.TODO())

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}

// runPgMigrations runs Postgres migrations
func runPgMigrations(dsn, path string) error {
	if path == "" {
		return errors.New("no migrations path provided")
	}
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	m, err := migrate.New(
		path,
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
