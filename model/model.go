package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"echo/helpers"
)

type M map[string]interface{}

type Item struct {
	Name       string `json:"name" form:"name"`
	EmployeeId int
	Employee   *Employee
}

// Employee represents the model for an employee
type Employee struct {
	ID        int    `json:"id" form:"id" swagger:"description(id)"`
	Full_Name string `json:"full_name" form:"full_name" swagger:"description(Full Name)" valid:"required"`
	Email     string `json:"email" form:"email" swagger:"description(Email)" valid:"required,email"`
	Password  string `json:"password" form:"password" swagger:"description(Password)" valid:"required"`
	Age       int    `json:"age" form:"age" swagger:"description(Age)" valid:"required"`
	Division  string `json:"division" form:"division" swagger:"description(Division)" valid:"required"`
	Item      []Item
}

type DeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(e)

	if errCreate != nil {
		err = errCreate
		return
	}

	e.Password = helpers.HassPass(e.Password)
	err = nil
	return
}
