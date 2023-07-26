package dao

import (
	"day3/kxS/cache"
	"day3/kxS/dbdata"
	serverz "day3/kxS/kitex_gen/kitex/serverZ"
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StudentDao struct {
	dbStu  *gorm.DB
	dbColl *gorm.DB

	studentCache *cache.Cache
}

func (dao *StudentDao) AddStudent(student *serverz.Student) (err error) {
	//dao.studentCache.Put(int(student.Id), student)

	studData := dbdata.NewStudent(student)

	result := dao.dbStu.Create(&studData)
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert student failed: " + result.Error.Error())
	}

	//fmt.Printf("\"create student done------------------\": %v\n", "create student done------------------")

	var collData dbdata.College
	result = dao.dbColl.First(&collData, "name = ?", student.College.Name)
	if result.RowsAffected == 0 {
		collData = dbdata.NewCollege(student)
		result = dao.dbColl.Create(&collData)
		if result.RowsAffected == 0 {
			return fmt.Errorf("insert college failed: " + result.Error.Error())
		}
	}

	return nil
}

func (dao *StudentDao) GetStudentById(id int32) (serverz.Student, error) {

	student, flag := dao.studentCache.Get(int(id))
	if flag {
		return *student, nil
	}

	var studData dbdata.Student
	var collData dbdata.College

	result := dao.dbStu.First(&studData, id)

	if result.RowsAffected == 0 {
		return *student, fmt.Errorf("no such student whose id = " + string(id))
	}

	result = dao.dbColl.First(&collData, "name = ?", studData.CollegeName)
	if result.RowsAffected == 0 {
		return *student, fmt.Errorf("no such college whose name =  " + studData.CollegeName)
	}

	student = serverz.NewStudent()

	student.Id = studData.Id
	student.Name = "studData.Name"
	student.College = &serverz.College{
		Name:    collData.Name,
		Address: collData.Address,
	}
	student.Email = strings.Split(studData.Emails, ",")
	student.Sex = studData.Sex

	return *student, nil
}

func (dao *StudentDao) InitDB() {
	dsn := "host=124.221.127.200 user=backend password=backend dbname=l23o6 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	// drop table
	db.Migrator().DropTable(dbdata.Student{})
	db.Migrator().DropTable(dbdata.College{})
	// create table
	err = db.Migrator().CreateTable(dbdata.Student{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().CreateTable(dbdata.College{})
	if err != nil {
		panic(err)
	}

	dao.dbStu = db.Table("students").Session(&gorm.Session{})
	dao.dbColl = db.Table("colleges").Session(&gorm.Session{})
	dao.studentCache = cache.NewCache(1000)
}
