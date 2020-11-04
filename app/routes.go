package app

import(
	"github.com/labstack/echo"
	"GoCRUD/handlers"
)

func Routes(e *echo.Echo){
	e.GET("/create", handlers.Create)
	e.GET("/delete",handlers.Delete)
	e.GET("/read",handlers.Read)
	e.GET("/update",handlers.Update)
}