package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// const htmlIndex = `<html><body>
// Logged in with <a href="/login">facebook</a>
// </body></html>
// `

// func handleMain(c echo.Context) error {
// 	return c.HTML(http.StatusOK, htmlIndex)
// }

// func main() {
// 	e := echo.New()

// 	e.GET("/", handleMain)

// 	e.GET("/login/facebook", handleFacebookLogin)
// 	e.GET("/callback/facebook", handleFacebookCallback)

// 	e.Logger.Fatal(e.Start(":8080"))
// }
