package middleware

import (
    "context"
	"log"
    "net/http"
    "strings"
    "time"

    userpb "product-ms/internal/infrastructure/client/pb"
)

func AuthMiddleware(userClient userpb.UserServiceClient) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            // if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            //     http.Error(w, "Unauthorized" + authHeader, http.StatusUnauthorized)
            //     return
            // }

            token := strings.TrimPrefix(authHeader, "Bearer ")

            ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
            defer cancel()

			log.Println("fdafdas")
            res, err := userClient.VerifyToken(ctx, &userpb.VerifyTokenRequest{Token: token})
            if err != nil || !res.Valid {
				log.Println(err)
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }
			log.Println("fdafdas2")

            ctx = context.WithValue(r.Context(), "userID", res.UserId)

            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
