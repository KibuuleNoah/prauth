package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // 7 days expiry
// const SessionMaxAge = 7 * 24 * 60 * 60 // seconds
//
// var store = sessions.NewCookieStore([]byte("testcookiekey"))
//
// // Initialize store settings
// func InitSessionStore() {
// 	store.Options = &sessions.Options{
// 		Path:     "/",
// 		MaxAge:   SessionMaxAge,
// 		HttpOnly: true,
// 		SameSite: http.SameSiteLaxMode,
// 		Secure:   false, // set true in production with HTTPS
// 	}
// }

// Middleware that checks if user is authenticated
func (r *Middleware) AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := r.AppCtx.CookieStore.Get(ctx.Request, "prauth_session")

		// No user ID found in session -> unauthorized
		userID, ok := session.Values["user_id"]
		fmt.Println()
		if !ok || userID == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/auth?dst=" + ctx.Request.URL.Path)
			return
		}

		// Make userID available in context
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}

