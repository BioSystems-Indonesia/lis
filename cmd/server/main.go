package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BioSystems-Indonesia/lis/internal/config"
	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
	"github.com/BioSystems-Indonesia/lis/internal/handler"
	"github.com/BioSystems-Indonesia/lis/internal/repository"
	"github.com/BioSystems-Indonesia/lis/internal/usecase"
)

func main() {
	dbConfig := config.GetDatabaseConfig()

	db, err := config.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established")

	if err := config.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	patientRepo := repository.NewPatientRepository(db)
	workOrderRepo := repository.NewWorkOrderRepository(db)

	patientUC := usecase.NewPatientUsecase(db, patientRepo)
	workOrderUC := usecase.NewWorkOrderUsecase(db, workOrderRepo, patientRepo)

	patientHandler := handler.NewPatientHandler(patientUC)
	workOrderHandler := handler.NewWorkOrderHandler(workOrderUC)

	mux := http.NewServeMux()

	mux.HandleFunc("/patients", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("q") != "" {
				patientHandler.Search(w, r)
			} else if r.URL.Query().Get("id") != "" {
				patientHandler.GetByID(w, r)
			} else {
				patientHandler.GetAll(w, r)
			}
		case http.MethodPost:
			patientHandler.Create(w, r)
		case http.MethodPut:
			patientHandler.Update(w, r)
		case http.MethodDelete:
			patientHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/work-orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("no_order") != "" {
				workOrderHandler.GetByNoOrder(w, r)
			} else if r.URL.Query().Get("doctor") != "" {
				workOrderHandler.GetByDoctor(w, r)
			} else if r.URL.Query().Get("analyst") != "" {
				workOrderHandler.GetByAnalyst(w, r)
			} else {
				workOrderHandler.GetAll(w, r)
			}
		case http.MethodPost:
			workOrderHandler.Create(w, r)
		case http.MethodPut:
			workOrderHandler.Update(w, r)
		case http.MethodDelete:
			workOrderHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC recovered: %v", err)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				response := dto.ResponseError{
					Code:    http.StatusInternalServerError,
					Status:  "Inrernal Server Error",
					Message: fmt.Sprintf("Internal Server Error: %v", err),
				}

				json.NewEncoder(w).Encode(response)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
