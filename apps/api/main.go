package main

import (
	"fmt"
	"net/http"

	"tarea-api/config"
	"tarea-api/controllers"
)

func main() {
	// 1. Conectar a la base de datos
	config.ConnectDB()

	// 2. Definir rutas
	http.HandleFunc("GET /api/tareas", controllers.GetTareas)

	// 3. Levantar servidor
	fmt.Println("Servidor corriendo en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}