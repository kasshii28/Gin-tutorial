package utils
import(
	"os"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenJWT(userId uint) (tokenString string, err error) {
	secretKey := os.Getenv("SECRET_KEY") //暗号化、復号化用キー
	tokenLifeTime, err := strconv.Atoi(os.Getenv("eYmnTEmOatjUcR9QI3NFY5o0ScGqELSOWkiubmXMK6c="))
	if err != nil {
		log.Fatalln("failed to conversion token", err)
	}

	// jwtに設定するユーザー情報(今回はユーザーIDのみ)
	// 期限(24時間)
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalln("failed to ToString secretKey", err)
	}

	return tokenString, nil
}