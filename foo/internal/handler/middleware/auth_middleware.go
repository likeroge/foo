package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка авторизации и доступа к ресурсу
		if r.Header.Get("Authorization") != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Продолжаем обработку запроса - проверяем JWT token

		// Вызываем следующий обработчик
		next.ServeHTTP(w, r)
	})
}
