package main

import (
	"context"
	"day3/kxS/dbdata"
	serverz "day3/kxS/kitex_gen/kitex/serverZ"
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	dbStu  *gorm.DB
	dbColl *gorm.DB
}

var memS = make(map[int32]*serverz.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *serverz.Student) (resp *serverz.RegisterResp, err error) {

	studData := dbdata.NewStudent(student)
	collData := dbdata.NewCollege(student)
	if _, ok := memS[student.Id]; ok {
		respN := serverz.RegisterResp{
			Success: true,
			Message: "no " + studData.Name,
		}
		return &respN, nil
	}

	memS[student.Id] = student

	result := s.dbStu.Create(&studData)
	if result.RowsAffected != 0 {
		result = s.dbColl.First(&collData, "name = ?", studData.CollegeName)
		if result.RowsAffected == 0 {
			result = s.dbColl.Create(&collData)
		}

	} else {
		fmt.Println(result.Error.Error())
	}

	respN := serverz.RegisterResp{
		Success: true,
		Message: "no " + studData.Name,
	}

	respN.Success = result.RowsAffected > 0

	if respN.Success {
		respN.Message = "yes " + studData.Name
	} else {
		respN.Message = "no " + studData.Name
	}

	resp = &respN

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *serverz.QueryReq) (resp *serverz.Student, err error) {

	if _, ok := memS[req.Id]; ok {
		resp = memS[req.Id]
		return
	}

	var studData dbdata.Student
	var collData dbdata.College

	var stud serverz.Student

	result := s.dbStu.First(&studData, req.Id)
	if result.RowsAffected != 0 {
		s.dbColl.First(&collData, "name = ?", studData.CollegeName)
	} else {
		result.Error = nil
	}

	stud.Id = studData.Id
	stud.Name = studData.Name
	stud.College = &serverz.College{
		Name:    collData.Name,
		Address: collData.Address,
	}
	stud.Email = strings.Split(studData.Emails, ",")
	stud.Sex = studData.Sex

	resp = &stud
	return
}

func (s *StudentServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
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

	s.dbStu = db.Table("students").Session(&gorm.Session{})
	s.dbColl = db.Table("colleges").Session(&gorm.Session{})
}

func (s *StudentServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	fmt.Println("Here!")
	reqS := request.(string)
	switch method {
	case "Query":
		var queryReq serverz.QueryReq
		err = json.Unmarshal([]byte(reqS), &queryReq)
		if err != nil {
			panic(err)
		}
		response, err = s.Query(ctx, &queryReq)

	case "Register":
		var registerReq serverz.Student
		err = json.Unmarshal([]byte(reqS), &registerReq)
		if err != nil {
			panic(err)
		}
		response, err = s.Register(ctx, &registerReq)
	}

	if err != nil {
		panic(err)
	}

	byteResp, err := json.Marshal(response)
	response = string(byteResp)

	return
}
