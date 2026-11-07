package main

import (
	"fmt"
	"log"
	"net/http"

	"tarea-api/config"
	"tarea-api/controllers"
)

func main() {
	// 1. Conectar a PostgreSQL
	config.ConnectDB()

	// 2. Definir rutas (compatibles con /api/tareas y /api/v1/tareas)
	http.HandleFunc("GET /api/tareas", controllers.GetTareas)
	http.HandleFunc("POST /api/tareas", controllers.CreateTarea)
	http.HandleFunc("GET /api/v1/tareas", controllers.GetTareas)
	http.HandleFunc("POST /api/v1/tareas", controllers.CreateTarea)

	// Healthcheck según estándar de la asignatura
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// 3. Levantar servidor
	fmt.Println("Servidor corriendo en el puerto 8080 (PostgreSQL)")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
