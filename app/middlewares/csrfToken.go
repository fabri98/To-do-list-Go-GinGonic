package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func CsrfToken() gin.HandlerFunc {
	os.Setenv("CSRF_SECRET_KEY", "tu_clave_secreta_aqui")

	csrfKey := []byte(os.Getenv("CSRF_SECRET_KEY"))
	csrfMiddleware := csrf.Protect(csrfKey, csrf.Secure(false))

	return func(c *gin.Context) {
		// Crea un http.Handler que ejecuta la lógica de Gin al final
		csrfHandler := csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Establecemos el Request y ResponseWriter en el contexto de Gin
			c.Request = r
			// Continuamos con el siguiente middleware de Gin
			c.Next()
		}))
		// Llamamos a ServeHTTP del middleware CSRF
		csrfHandler.ServeHTTP(c.Writer, c.Request)

		// Abortamos si CSRF falla
		if c.IsAborted() {
			return
		}
	}
}
