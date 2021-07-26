package Models
import (
	"fmt"
	"github.com/Exercise2/Config"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all student data
func GetAllUsers(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	//Config.DB.Where()
	return nil
}
//CreateUser ... Insert New data
func CreateUser(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one student by Id
func GetUserByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update student
func UpdateUser(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}
//DeleteUser ... Delete student
func DeleteUser(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}