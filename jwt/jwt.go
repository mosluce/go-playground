package main

import "github.com/dgrijalva/jwt-go"
import "fmt"
import "time"

type HelloClaim struct {
	jwt.StandardClaims // 繼承標準 Claims

	Name string // 任意自訂欄位
}

func main() {
	secret := []byte("12345678")

	claim := HelloClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 10, // UnixTimestamp 單位秒, ex: 10秒後
		},
		"mosluce",
	}

	// 建立 token 物件
	// jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 建立 token 物件並取出字串
	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(secret)
	// 輸出
	fmt.Println("token >>", tokenString)

	// 從 token string 轉換成 token (需指定 Claim 型別)
	token, err := jwt.ParseWithClaims(tokenString, &HelloClaim{}, func(token *jwt.Token) (interface{}, error) {
		// TODO 可以在這裡檢查 method 之類的
		return secret, nil
	})

	// 過期的話會產生 error
	if err != nil {
		fmt.Println(err)
		return
	}

	// 轉型(?) 並輸出
	if claim, ok := token.Claims.(*HelloClaim); ok && token.Valid {
		fmt.Println("ok >>", claim.Name)
	} else {
		fmt.Println("failed!!!")
	}

}
