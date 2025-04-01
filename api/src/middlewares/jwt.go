package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/jwtauth/v5"
	configs "github.com/redrum15/prueba/src/config"
	"github.com/rs/zerolog/log"
)

type ContextKey string

const (
	UserIDKey   ContextKey = "userID"
	UserRoleKey ContextKey = "userRole"
)

var ja *jwtauth.JWTAuth = jwtauth.New("HS256", []byte(configs.Envs.JWTSecret), nil)

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

func JWTVerifier() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractTokenFromHeader(r)
			if token == "" {
				http.Error(w, "Authorization token required", http.StatusUnauthorized)
				return
			}

			jwtToken, err := ja.Decode(token)
			if err != nil {
				log.Error().Err(err).Msg("Failed to extract claims from JWT token")
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			fmt.Println(token)
			if _, err := jwtauth.VerifyToken(ja, token); err != nil {
				log.Error().Err(err).Msg("Failed to extract claims from JWT token")
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			claims, err := jwtToken.AsMap(r.Context())
			if err != nil {
				log.Error().Err(err).Msg("Failed to extract claims from JWT token")
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			if userID, ok := claims["user_id"].(string); ok {
				ctx = context.WithValue(ctx, UserIDKey, userID)
			}
			if userRole, ok := claims["role"].(string); ok {
				ctx = context.WithValue(ctx, UserRoleKey, userRole)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(UserRoleKey).(string)
			if !ok {
				http.Error(w, "Role information not available", http.StatusUnauthorized)
				return
			}

			hasRole := false
			for _, requiredRole := range roles {
				if role == requiredRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func GenerateJWT(userID, role string) (string, error) {
	claims := map[string]interface{}{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	_, tokenString, err := ja.Encode(claims)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate JWT token")
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

func GetUserRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(UserRoleKey).(string)
	return role, ok
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
