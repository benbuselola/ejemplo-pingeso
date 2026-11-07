package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"tarea-api/config"
	"tarea-api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTareas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var tareas []models.Tarea
	collection := config.DB.Collection("tareas")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Obtener todos los documentos
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var tarea models.Tarea
		cursor.Decode(&tarea)
		tareas = append(tareas, tarea)
	}

	// Si no hay tareas, devolver un arreglo vacío en lugar de null
	if tareas == nil {
		tareas = []models.Tarea{}
	}

	json.NewEncoder(w).Encode(tareas)
}