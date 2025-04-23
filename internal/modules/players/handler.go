package playermodule

import (
	"encoding/json"
	"net/http"

	"pustobaseproject/internal/middleware"
	"pustobaseproject/pkg/utils"
)

func GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем игрока из контекста
	p := middleware.GetPlayerFromContext(r)
	if p == nil {
		http.Error(w, `{"error": "Player not found in context"}`, http.StatusInternalServerError)
		return
	}

	// 2. Расшифровываем EncryptedPlayerID → получаем оригинальный user_id
	userID, err := utils.EncryptionServiceInstance.Decrypt(p.EncryptedPlayerID)
	if err != nil {
		http.Error(w, `{"error": "Failed to decrypt player ID"}`, http.StatusInternalServerError)
		return
	}

	// 3. Формируем и возвращаем ответ
	resp := struct {
		PlayerID string `json:"player_id"`
	}{
		PlayerID: userID,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		return
	}
}
