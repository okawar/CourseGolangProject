package service

import (
	"encoding/json"
	"errors"
	"golang_pr/entity"
	"golang_pr/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	CRUDS["UserAuth"] = Authorize
}

type JWTClaims struct {
	UserID uint32 `json:"user_id"`
	jwt.StandardClaims
}

var active_users []uint32
var jwtSecret = []byte("your_secret_key") // Секретный ключ для подписи JWT
const jwtTokenLifetime = 2                // Время жизни токена (в часах)

func GenerateToken(userID uint32) (string, error) {
	clms := JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * jwtTokenLifetime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	return token.SignedString(jwtSecret)
}

func Authorize(c *Context) {
	var usr entity.User
	usr, err := GetData[entity.User](c)
	if err != nil {
		c.sendError(400, "Invalid data for `user` entity.")
	}
	storedUsers := sql.GetAllUsers()
	finded := false
	for _, v := range storedUsers {
		if v.Login == usr.Login && v.Password == usr.Password {
			finded = true
			usr = *v
		}
	}
	if !finded {
		c.sendError(401, "Invalid login or password")
	} else {
		active_users = append(active_users, usr.UserId)
		token, err := GenerateToken(usr.UserId)
		if err != nil {
			c.sendError(500, "Failed generating token")
		} else {
			http.SetCookie(c.Response, &http.Cookie{
				Expires:  time.Now().Add(time.Hour * jwtTokenLifetime),
				Value:    token,
				Name:     "jwt",
				HttpOnly: true,
				Path:     "/",
			})
			c.Response.WriteHeader(303)
			c.Response.Write([]byte("Success;ID=" + strconv.FormatUint(uint64(usr.UserId), 10)))
			c.Response.Header().Add("Location", "/index")
		}
	}
}
func VerifyToken(tkn_str string) (*JWTClaims, error) {
	var buff []byte
	json.Unmarshal(buff, &tkn_str)
	token, err := jwt.ParseWithClaims(tkn_str, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func Authentificate(c *Context) (bool, int) {
	tkn, err := c.Request.Cookie("jwt")
	if err != nil {
		log.Println("No token sended.")
		return false, 401
	}
	clms, err := VerifyToken(tkn.Value)
	if err != nil {
		log.Println("Error of verifying token.")
		return false, 500
	} else {
		is_active := false
		for _, u := range active_users {
			is_active = is_active || (u == clms.UserID)
		}
		if is_active && time.Now().Unix() < clms.ExpiresAt {
			log.Printf("Successfully authentificated. User #%d\n", clms.UserID)
			return true, 200
		}
	}
	return false, 0
}
