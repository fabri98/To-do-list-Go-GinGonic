package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func SetCookie() sessions.Store {
	// Configura el almacenamiento de la sesión
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 600, // 5 minutos
		Path:   "/",
	})
	return store
}
