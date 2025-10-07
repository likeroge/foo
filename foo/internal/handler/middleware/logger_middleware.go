package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем информацию о запросе
		fmt.Printf("Метод: %s, Путь: %s\n", r.Method, r.URL.Path)

		// Вызываем следующий обработчик
		next.ServeHTTP(w, r)

		// Логируем время обработки
		fmt.Printf("Время обработки: %v\n", time.Since(start))
	})
}
