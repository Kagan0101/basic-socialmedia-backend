package savedatabase


import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type Saves struct{
	gorm.Model
	İdea string
}


func (saves Saves) Migrate(){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&saves)
}

func (saves Saves) Add(){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		return
	}

	db.Create(&saves)
}

func (saves Saves) Get(where...interface{}) Saves {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return saves
	}
	db.First(&saves,where...)
	return saves
}

func (saves Saves) GetAll(where...interface{}) []Saves {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var users1 []Saves
	db.Find(&users1, where...)
	return users1
}

func (saves Saves) Update(column string,value interface{})  {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&saves).Update(column,value)
}

func (saves Saves) Updates(data Saves){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&saves).Updates(data)
}

func (saves Saves) Delete()  {
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&saves,saves.ID)
}


