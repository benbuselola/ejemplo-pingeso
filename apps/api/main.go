package main

import (
	"fmt"
	"log"
	"net/http"

	"tarea-api/config"
	"tarea-api/controllers"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/tareas", controllers.GetTareas)
	mux.HandleFunc("POST /api/tareas", controllers.CreateTarea)
	mux.HandleFunc("OPTIONS /api/tareas", controllers.GetTareas)
	mux.HandleFunc("GET /api/v1/tareas", controllers.GetTareas)
	mux.HandleFunc("POST /api/v1/tareas", controllers.CreateTarea)
	mux.HandleFunc("OPTIONS /api/v1/tareas", controllers.GetTareas)

	// Healthcheck según estándar de la asignatura
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return mux
}

func main() {
	// 1. Conectar a PostgreSQL
	config.ConnectDB()

	// 2. Definir rutas
	mux := setupRouter()

	// 3. Levantar servidor
	fmt.Println("Servidor corriendo en el puerto 8080 (PostgreSQL)")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
