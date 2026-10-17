package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TestJWTAuthMiddleware(t *testing.T) {
	secret := "test-secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": 42})
	signed, _ := token.SignedString([]byte(secret))

	r := gin.New()
	r.Use(JWTAuth(secret))
	r.GET("/secure", func(c *gin.Context) {
		uid, _ := c.Get(ContextUserID)
		c.JSON(200, gin.H{"uid": uid})
	})

	req := httptest.NewRequest("GET", "/secure", nil)
	req.Header.Set("Authorization", "Bearer "+signed)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
