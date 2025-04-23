package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"

	"pustobaseproject/internal/domain/players"
	"pustobaseproject/pkg/utils"
)

// contextKey — это тип, чтобы избежать конфликта ключей в context.Context.
// Мы не используем string напрямую, потому что разные middleware могут
// использовать одинаковые ключи ("player", "user") и перезаписывать данные.
type contextKey string

// playerContextKey — ключ для хранения *player.Player в context.Context запроса.
const playerContextKey = contextKey("player")

// PlayerMiddleware возвращает middleware-функцию.
// Она проверяет JWT-токен, извлекает userID, ищет или создаёт игрока,
// и сохраняет его в context запроса.
func PlayerMiddleware(service *players.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 1. Получаем токен из заголовка Authorization: Bearer <token>
			tokenStr, err := utils.RetrieveTokenFromRequest(r)
			if err != nil {
				http.Error(w, `{"error": "Missing token"}`, http.StatusUnauthorized)
				return
			}

			// 2. Декодируем токен и получаем claims (в них содержится encrypted userID)
			claims, err := utils.DecodeToken(tokenStr)
			if err != nil || claims.Subject == "" {
				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
				return
			}

			// 3. Хешируем userID через SHA-256 (это безопасный и уникальный идентификатор)
			userID := claims.Subject
			hash := sha256.Sum256([]byte(userID))
			hashedID := hex.EncodeToString(hash[:])

			// 4. Пытаемся найти игрока в базе по хешу
			player, err := service.GetPlayerByHashedID(hashedID)
			if err != nil {
				http.Error(w, `{"error": "DB error"}`, http.StatusInternalServerError)
				return
			}

			// 5. Если игрок не найден и мы находимся на save-эндпоинте → создаём игрока
			if player == nil && strings.HasSuffix(r.URL.Path, "/save") {
				encryptedPlayerID, err := utils.EncryptionServiceInstance.Encrypt(userID)
				if err != nil {
					http.Error(w, `{"error": "Failed to encrypt player ID"}`, http.StatusInternalServerError)
					return
				}
				player, err = service.CreatePlayer(hashedID, encryptedPlayerID)
				if err != nil {
					http.Error(w, `{"error": "Failed to create player"}`, http.StatusInternalServerError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), playerContextKey, player)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetPlayerFromContext достаёт игрока из context.Context.
// Используется в хендлерах, чтобы не искать игрока заново в БД.
func GetPlayerFromContext(r *http.Request) *players.Player {
	p, _ := r.Context().Value(playerContextKey).(*players.Player)
	return p
}
