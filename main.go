package main

import (
	"fmt"
	"net/http"
	"shopix/model"
	"shopix/storage"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	register, err := storage.LoadUsers()
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()

	tmplReg := template.Must(template.ParseFiles("template/register_page.html"))

	tmplSignIn := template.Must(template.ParseFiles("template/login_page.html"))

	e.Static("/", "static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Успешно!")
	})

	e.GET("/register", func(c echo.Context) error {
		return tmplReg.Execute(c.Response(), model.ErrReg{})
	})

	e.POST("/", func(c echo.Context) error {
		login := c.FormValue("login")
		pass := c.FormValue("password")
		confirm := c.FormValue("confirm")

		for _, v := range register {
			if login == v.Login {
				return tmplReg.Execute(c.Response(), model.ErrReg{
					Login:    login,
					ErrLogin: "Такой логин уже существует!",
				})
			}
		}

		if pass != confirm {
			return tmplReg.Execute(c.Response(), model.ErrReg{
				Login:   login,
				ErrPass: "Пароли не совпадают",
			})
		}

		register = append(register, model.Registration{
			ID:    len(register) + 1,
			Login: login,
			Pass:  pass,
		})

		fmt.Println(register)
		storage.SavedUsers(register)

		return c.Redirect(http.StatusSeeOther, "/")

	})

	e.GET("/SignIn", func(c echo.Context) error {
		return tmplSignIn.Execute(c.Response(), model.ErrSign{})
	})

	e.POST("/", func(c echo.Context) error {
		login := c.FormValue("login")
		pass := c.FormValue("password")

		for _, v := range register {
			if login == v.Login {
				if pass != v.Pass {
					return tmplSignIn.Execute(c.Response(), model.ErrSign{
						Login:   login,
						ErrPass: "Неверный пароль",
					})
				}

				return c.Redirect(http.StatusSeeOther, "/")
			}
		}

		return tmplSignIn.Execute(c.Response(), model.ErrSign{
			Login:    login,
			ErrLogin: "Аккаунт с таким логином не найден!",
		})

	})

	e.Start(":8080")
}
