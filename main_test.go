package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status code esperado %d, recebeu %d", http.StatusOK, res.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("erro ao decodificar JSON: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("status esperado 'ok', recebeu '%s'", body["status"])
	}
}

func TestHelloEndpointDefault(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status code esperado %d, recebeu %d", http.StatusOK, res.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("erro ao decodificar JSON: %v", err)
	}

	expected := "Olá, mundo!"
	if body["message"] != expected {
		t.Fatalf("mensagem esperada '%s', recebeu '%s'", expected, body["message"])
	}
}

func TestHelloEndpointWithName(t *testing.T) {
	r := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ana", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status code esperado %d, recebeu %d", http.StatusOK, res.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("erro ao decodificar JSON: %v", err)
	}

	expected := "Olá, Ana!"
	if body["message"] != expected {
		t.Fatalf("mensagem esperada '%s', recebeu '%s'", expected, body["message"])
	}
}
