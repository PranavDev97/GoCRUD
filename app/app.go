package app

import(
	"github.com/labstack/echo"
	"net/http"
	"GoCRUD/conf"
	"fmt"
)

func Run(){
	e := echo.New()
	Routes(e)
	s := &http.Server{
		Addr:         conf.Conf.Port,
	}
	err := e.StartServer(s)
	if err!=nil {
		fmt.Println("aa")
		e.Logger.Info("shutting down the server")
	}
}