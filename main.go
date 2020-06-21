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
	// struct {
	//         char *Name;
	// 				char Mail;
	// 				char Password;
	// } nazuna, tomoyuki, ryoma;

	// int main() {
	//         nazuna.Name = "nazuna";
	// 				nazuna.Mail = "nazuna@gmail.com";
	// 				nazuna.Password = "password"
	//         Boobs.Name = "Boobs";
	// 				Boobs.Mail = "b@gmail.com";
	// 				Boobs.Password = "admin"
	// 				Hery.Name = "Hery";
	// 				Hery.Mail = "Hery@gmail.com";
	// 				Hery.Password = "hello"
	// }

	u1 := User{
		Name:     "nazuna",
		Mail:     "me@example.com",
		Password: "password",
	}
	u2 := User{
		Name:     "Boobs",
		Mail:     "b@gmail.com",
		Password: "admin",
	}
	u3 := User{
		Name:     "Hery",
		Mail:     "Hery@example.com",
		Password: "hello",
	}
	// ルーティング設定
	e.GET("/", rootHandler)
	e.GET("/ok", okHandler)
	e.POST("/login", func(c echo.Context) error {
		r := new(User)
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if r == u1 {
			return c.JSON(http.StatusOK, r)
		}

		if r.Name == u1.Name && r.Mail == u1.Mail && r.Password == u1.Password || r.Name == u2.Name && r.Mail == u2.Mail && r.Password == u2.Password || r.Name == u3.Name && r.Mail == u3.Mail && r.Password == u3.Password {
			return c.JSON(http.StatusOK, r)
		} else {
			return c.String(http.StatusUnauthorized, "login fail")
		}
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
