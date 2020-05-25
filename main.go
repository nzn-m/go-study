package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type User struct {
	Name     string `json:"Name"`
	Mail     string `json:"Mail"`
	Password string `json:"Password"`
}

func main() {
	// サーバー用のインスタンスの取得
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323/", "http://localhost:1323/login/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t
	// ユーザー
	u := User{
		Name:     "nazuna",
		Mail:     "me@example.com",
		Password: "password",
	}
	// ルーティング設定
	e.GET("/", rootHandler)
	e.GET("/ok", okHandler)
	e.POST("/login", func(c echo.Context) error {
		r := new(User)
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if r.Mail != u.Mail || r.Password != u.Password {
			return c.String(http.StatusUnauthorized, "login fail")
		}
		// 暫定
		return c.JSON(http.StatusOK, r)
	})
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}

func rootHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
func okHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "ok", nil)
}
