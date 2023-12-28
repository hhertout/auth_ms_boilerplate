package e2e_test

import (
	"auth_ms/src/router"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPingRoute(t *testing.T) {
	err := os.Setenv("APP_ENV", "dev")
	if err != nil {
		t.Errorf("Failed to set env")
	}
	r := router.Serve()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r.ServeHTTP(w, req)

	t.Run("Expect status code is ok", func(t *testing.T) {
		if w.Code != 200 {
			t.Errorf("Invalid request")
		}
	})
}
