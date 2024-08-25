package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	Name,Surname,Idea,Email,Password,Saved,Likeclass string
	Likenumber int
}

func (users Users) Migrate() {
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&users)
}

func (users Users) Add(){
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		return
	}

	db.Create(&users)
}

func (users Users) Get(where...interface{}) Users {
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return users
	}
	db.First(&users,where...)
	return users
}

func (users Users) GetAll(where...interface{}) []Users {
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var users1 []Users
	db.Find(&users1, where...)
	return users1
}

func (users Users) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&users).Update(column,value)
}

func (users Users) Updates(data Users){
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&users).Updates(data)
}

func (users Users) Delete()  {
	db,err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&users,users.ID)
}