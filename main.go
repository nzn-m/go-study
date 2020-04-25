package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// サーバー用のインスタンスの取得
	e := echo.New()
	// ユーザー
	u := User{
		Email:    "me@example.com",
		Password: "password",
	}
	// ルーティング設定
	e.GET("/helloworld", helloWorld)
	e.POST("/login", func(c echo.Context) error {
		fmt.Printf("u.Email : %v\n", u.Email)
		fmt.Printf("u.Password : %v\n", u.Password)
		r := new(User)
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		//debug
		fmt.Printf("u.Email : %v\n", u.Email)
		fmt.Printf("u.Password : %v\n", u.Password)
		fmt.Printf("r.Email : %v\n", r.Email)
		fmt.Printf("r.Password : %v\n", r.Password)

		// もしうえのEmailとcurlでにゅうりょくしたEmailがいしないか
		//それか,うえのパスワードとcurlで入力したpasswordが一致しなければ,"login fail"としゅつりょく
		if r.Email != u.Email || r.Password != u.Password {
			return c.String(http.StatusUnauthorized, "login fail")
		}
		// 暫定
		token := "success"
		name := "なずな"
		return c.String(http.StatusOK, "{\"token\":\""+token+"\"},{\"name\":\""+name+"\"}")
	})
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!")
}
