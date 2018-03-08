package main

import (
	"api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (mdb *mockDB) AllProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	products = append(products, &models.Product{1, "Emma", 9.44})
	products = append(products, &models.Product{2, "James", 5.99})
	return products, nil
}

func TestProducts(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error(err)
	}

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.getProducts).ServeHTTP(w, req)

	products, _ := env.db.AllProducts()
	result, _ := json.Marshal(products)
	expected := string(result)
	received := w.Body.String()

	if expected != received {
		t.Errorf("\n...expected = %v\n...received = %v", expected, received)
	}
}
