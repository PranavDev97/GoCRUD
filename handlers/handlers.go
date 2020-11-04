package handlers

import(
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
	"GoCRUD/domain"
	"GoCRUD/backend"
	"fmt"
	"strconv"
)

func Create(c echo.Context) error {
	studentDetails := domain.StudentCredentials{}
	err := json.Unmarshal([]byte(c.FormValue("studentdetails")), &studentDetails);
	//err:=c.Bind(&studentDetails)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusForbidden,"Error while binding data")
	}
	response,err := backend.Create(studentDetails)
	if err!=nil{
		return echo.NewHTTPError(http.StatusForbidden, "Error while inserting data into DB")
	}
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	var sid int
	err:=json.Unmarshal([]byte(c.FormValue("studentid")), &sid);
	response,err := backend.Delete(sid)
	if err!=nil{
		return echo.NewHTTPError(http.StatusForbidden, "Error while inserting data into DB")
	}
	return c.JSON(http.StatusOK, response)
}

func Read(c echo.Context) error {
	sid:=c.QueryParam("sid")
	id,err:=strconv.Atoi(sid)
	if err!=nil{
		fmt.Println("string conversion error:",err)
	}
	response,err:=backend.Read(id)
	if err!=nil{
		return echo.NewHTTPError(http.StatusForbidden, "Error while Reading data from DB")
	}
	return c.JSON(http.StatusOK,response)
}
func Update(c echo.Context) error {
	sid:=c.QueryParam("sid")
	id,err:=strconv.Atoi(sid)
	if err!=nil{
		fmt.Println("string conversion error:",err)
	}
	marksString:=c.FormValue("marks")
	marks,err := strconv.Atoi(marksString)
	if err!=nil{
		fmt.Println("string conversion error for marks:",err)
	}
	response,err:=backend.Update(id,marks)
	if err!=nil {
		return echo.NewHTTPError(http.StatusForbidden, "Error while deleting data from DB")
	}
	return c.JSON(http.StatusOK,response)
}