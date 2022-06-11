package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CookieName       = "cookie_user"
	CookieTimeLength = 10 * 60
)

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err != nil {
		return nil, err
	}
	context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	return cookie, nil
}
