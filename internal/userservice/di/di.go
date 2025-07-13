package dius

import (
	"context"
	"database/sql"
	"fmt"
	handlers "github.com/harshgupta9473/chatapp/internal/userservice/handler"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/harshgupta9473/chatapp/internal/userservice/repositories"
	"github.com/harshgupta9473/chatapp/internal/userservice/services"
)

type UserServiceContainer struct {
	UserRepo    *repositories.UserRepository
	UserService *services.UserService
	AuthHandler *handlers.AuthHandler
}

func initUserServiceContainer() (*UserServiceContainer, error) {
	// Read DSN from env or hardcoded for now
	connStr := "host=localhost port=5432 user=postgres password=supersecurepass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ensure DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	return &UserServiceContainer{
		UserRepo:    userRepo,
		UserService: userService,
		AuthHandler: authHandler,
	}, nil
}

func InitializeUserService() error {
	container, err := initUserServiceContainer()
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	router.HandleFunc("/signup", container.AuthHandler.SignupHandler).Methods("POST")
	router.HandleFunc("/login", container.AuthHandler.LoginHandler).Methods("POST")

	// Start HTTP server
	go func() {
		log.Println("User service started at :8081")
		if err := http.ListenAndServe(":8081", router); err != nil {
			log.Fatalf("User service HTTP server error: %v", err)
		}
	}()

	return nil
}
