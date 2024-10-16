package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/keito-isurugi/auth-demo/helper"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorization ヘッダーからトークンを取得
		// authHeader := r.Header.Get("Authorization")
		// fmt.Println(authHeader)
		// if authHeader == "" {
		// 	http.Error(w,"Missing Authorization Header", http.StatusUnauthorized)
		// 	return
		// }
		
		// "Bearer トークン" 形式なので、"Bearer " を取り除く
		// tokenString := authHeader[len("Bearer "):]
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value

		// トークンの検証
		token, err := helper.ValidateJWT(tokenString)
		if err != nil {
			fmt.Println(token, err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// クレームからユーザーID(sub)を取得
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["sub"].(string)
			fmt.Println(userID)

			// レスポンスにユーザーIDを返す
			// w.Write([]byte(fmt.Sprintf("UserID: %s", userID)))
			next(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	}
}