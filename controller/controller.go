package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"echo/config"
	_ "echo/docs"
	models "echo/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// CreateEmployee godoc
// @Summary Post
// @Description Post Employee
// @Tags employee
// @Accept  json
// @Produce  json
// @Param model.Employee body model.Employee true "create employee"
// @Success 200 {object} model.Employee
// @Router /employee [post]
func CreateEmployee(ctx echo.Context) error {
	db := config.GetDB()
	employee := models.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Debug().Create(&employee)

	fmt.Println("CreateEmployee")
	return ctx.JSON(http.StatusOK, employee)
}

func CreateItem(ctx echo.Context) error {
	db := config.GetDB()
	item := models.Item{}

	userData, ok := ctx.Get("userData").(jwt.MapClaims)
	if !ok {
		// Handle the case where userData is not of type jwt.MapClaims
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     userData,
			"message": "Failed to get user data",
		})
	}

	userID := uint(userData["id"].(float64))

	if err := ctx.Bind(&item); err != nil {
		return err
	}

	item.EmployeeId = int(userID)

	db.Debug().Create(&item)

	fmt.Println("CreateItem")
	return ctx.JSON(http.StatusOK, item)
}

func HelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func JsonMap(ctx echo.Context) error {
	data := models.M{"message": "Hello", "counter": 2, "statusKode": http.StatusOK}
	return ctx.JSON(http.StatusOK, data)
}

func Page1(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	data := "Hello " + name
	result := fmt.Sprintf("%s", data)
	fmt.Println(result)
	return ctx.String(http.StatusOK, result)
}

// func User(ctx echo.Context) error {
// 	user := model.User{}
// 	// u := new(model.User)
// 	if err := ctx.Bind(&user); err != nil {
// 		return err
// 	}
// 	return ctx.JSON(http.StatusOK, user)
// }

func UpdateEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := models.Employee{}

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	// db.Save(&employee)

	db.Model(&employee).Where("id = ?", employee.ID).
		Updates(models.Employee{
			Full_Name: employee.Full_Name,
			Email:     employee.Email,
			Age:       employee.Age,
			Division:  employee.Division,
		})

	fmt.Println("UpdateEmployee")
	return ctx.JSON(http.StatusOK, employee)
}

func DeleteEmployee(ctx echo.Context) error {
	db := config.GetDB()

	employee := models.Employee{}

	delResp := models.DeleteResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	}

	paramId := ctx.Param("id")

	if err := ctx.Bind(&employee); err != nil {
		return err
	}

	db.Model(&employee).Where("id = ?", paramId).Delete(&employee)

	fmt.Println("DeleteEmployee")
	return ctx.JSON(http.StatusOK, delResp)
}

func Index(ctx echo.Context) error {
	tmpl :=
		template.Must(template.ParseGlob("template/*.html"))

	type M map[string]interface{}

	data := make(M)
	data[config.CSRFKey] = ctx.Get(config.CSRFKey)
	return tmpl.Execute(ctx.Response(), data)
}

func SayHello(ctx echo.Context) error {
	data := make(map[string]interface{})

	if err := ctx.Bind(&data); err != nil {
		return err
	}

	message := 
	fmt.Sprintf("Hello %s , My Gender %s", data["name"], data["gender"])

	return ctx.JSON(http.StatusOK, message)
}


