package middleware

import (
    "context"
    "net/http"
    "strings"
    "time"

    userpb "order-ms/internal/infrastructure/client/pb"
)

func AuthMiddleware(userClient userpb.UserServiceClient) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            // if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            //     http.Error(w, "Unauthorized", http.StatusUnauthorized)
            //     return
            // }

            token := strings.TrimPrefix(authHeader, "Bearer ")

            ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
            defer cancel()

            res, err := userClient.VerifyToken(ctx, &userpb.VerifyTokenRequest{Token: token})
            if err != nil || !res.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            ctx = context.WithValue(r.Context(), "userID", res.UserId)

            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
