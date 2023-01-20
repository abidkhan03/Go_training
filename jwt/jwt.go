package jwt

import (
    "fmt"
    "github.com/go-chi/jwtauth/v5"
    "net/http"
)


//create /auth/sigin route
//Use JWT Authentication middleware to limit access for users to /parse route.
//For example, you can create a route that only authenticated users can access.
//You can also create a route that only authenticated users with a specific role can access.

var TokenAuth *jwtauth.JWTAuth

func init() {
    TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func Authenticator(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token, claims, err := jwtauth.FromContext(r.Context())

        t,err:=TokenAuth.Decode(token)
        t.PrivateClaims()
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        userRole := claims["role"].(string)
        if userRole != "admin" {
            http.Error(w, "Forbidden", 403)
            return
        }
        next.ServeHTTP(w, r) // call next handler
    })

}


