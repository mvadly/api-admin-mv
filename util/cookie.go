package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveCookie(value string, r *gin.Context) {
	cookieName := "CookieData"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != "" {
		c = &http.Cookie{
			Name:       cookieName,
			Value:      value,
			Path:       "",
			Domain:     "api.com",
			Expires:    time.Now().Add(5 * time.Minute),
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   []string{},
		}
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = value
		c.Expires = time.Now().Add(5 * time.Minute)
		r.SetCookie(c.Name, c.Value, 0, "/", "adminmv.com", true, true)
	}

}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = "token"
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}
