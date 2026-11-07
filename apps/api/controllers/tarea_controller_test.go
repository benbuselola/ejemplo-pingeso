package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewID(t *testing.T) {
	id1 := newID()
	id2 := newID()
	if len(id1) != 16 {
		t.Errorf("esperaba id de longitud 16, obtuve %d", len(id1))
	}
	if id1 == id2 {
		t.Errorf("se generaron dos IDs idénticos: %s", id1)
	}
}

func TestGetTareas_Options(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/api/tareas", nil)
	rec := httptest.NewRecorder()

	GetTareas(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperaba status 200 para OPTIONS, obtuve %d", rec.Code)
	}
	if rec.Header().Get("Access-Control-Allow-Origin") == "" {
		t.Error("esperaba header Access-Control-Allow-Origin configurado")
	}
}

func TestCreateTarea_Options(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/api/tareas", nil)
	rec := httptest.NewRecorder()

	CreateTarea(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperaba status 200 para OPTIONS, obtuve %d", rec.Code)
	}
}

func TestCreateTarea_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/tareas", bytes.NewBufferString("{invalido}"))
	rec := httptest.NewRecorder()

	CreateTarea(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("esperaba status 400 para JSON inválido, obtuve %d", rec.Code)
	}
}

func TestCreateTarea_EmptyTitulo(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/tareas", bytes.NewBufferString(`{"titulo": ""}`))
	rec := httptest.NewRecorder()

	CreateTarea(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("esperaba status 400 para título vacío, obtuve %d", rec.Code)
	}
}
