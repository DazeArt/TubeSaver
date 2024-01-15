package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeValid2ecodeValid(t *testing.T) {
	token := "eyJhbGciOiAiSFMyNTYiLCAidHlwIjogIkpXVCJ9.eyJsb2dpbiI6ICJpIHdhbnQiLCAicGFzc3dvcmQiOiAiMTAwIiwgImV4cCI6IDYwMH0.e9f116b95558c4b8e4eba993bcb059dbc622753f519b6e64edc744fe1f6cb373"
	expected := make(jwt.MapClaims)
	expected["login"] = "i want"
	expected["password"] = "100"
	result := DecodeValid(token)
	if expected["login"] != result["login"] || expected["password"] != result["password"] {
		t.Errorf("Incorrect result. Expected %s, got %s", expected, result)
	}
}

func TestHistoryhandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Создаем хэндлер с помощью функции http.HandlerFunc
	handler := http.HandlerFunc(historyhandle)

	// Вызываем метод ServeHTTP хэндлера, передавая тестовый ответ и тестовый запрос
	handler.ServeHTTP(rr, req)

	// Проверяем, что код статуса тестового ответа равен 200
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestAuthhandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Создаем хэндлер с помощью функции http.HandlerFunc
	handler := http.HandlerFunc(authhandle)

	// Вызываем метод ServeHTTP хэндлера, передавая тестовый ответ и тестовый запрос
	handler.ServeHTTP(rr, req)

	// Проверяем, что код статуса тестового ответа равен 200
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestEnterhandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Создаем хэндлер с помощью функции http.HandlerFunc
	handler := http.HandlerFunc(enterhandle)

	// Вызываем метод ServeHTTP хэндлера, передавая тестовый ответ и тестовый запрос
	handler.ServeHTTP(rr, req)

	// Проверяем, что код статуса тестового ответа равен 200
	assert.Equal(t, http.StatusOK, rr.Code)
}
