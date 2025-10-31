package middleware

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/web"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Handler http.Handler
}

type contextKey string

const userIDKey = contextKey("userId")

func ContextWithUserID(request *http.Request, userId int) *http.Request {
	ctx := context.WithValue(request.Context(), userIDKey, userId)
	return request.WithContext(ctx)
}

func GetUserIDFromContext(ctx context.Context) int {
	if val, ok := ctx.Value(userIDKey).(int); ok {
		return val
	}
	return 0
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	// Public paths that don't require authorization
	publicPaths := []string{
		"/api/users/login",
		"/api/users/register",
	}

	// Allow public paths to bypass authentication
	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			middleware.Handler.ServeHTTP(writer, request)
			return
		}
	}

	// Get Authorization header
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Missing or invalid Authorization header",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Extract token from header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Define your JWT secret key
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	// Parse and validate JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Invalid or expired token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Extract user_id from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Invalid token claims",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Missing user_id in token",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userID := int(userIDFloat)

	// Add userId to request context
	newRequest := ContextWithUserID(request, userID)

	// Proceed to the next handler
	middleware.Handler.ServeHTTP(writer, newRequest)
}
