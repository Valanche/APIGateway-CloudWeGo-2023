package main

import (
	"context"
	"day3/kxS/dao"
	serverz "day3/kxS/kitex_gen/kitex/serverZ"
	"encoding/json"
	"fmt"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	studentDao dao.StudentDao
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *serverz.Student) (resp *serverz.RegisterResp, err error) {

	err = s.studentDao.AddStudent(student)

	respN := serverz.RegisterResp{
		Success: true,
		Message: "z",
	}

	respN.Success = err == nil

	if respN.Success {
		respN.Message = "yes " + student.Name
	} else {
		respN.Message = "no " + student.Name
	}

	resp = &respN
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *serverz.QueryReq) (resp *serverz.Student, err error) {

	var stud serverz.Student

	stud, err = s.studentDao.GetStudentById(req.Id)

	resp = &stud
	return
}

func (s *StudentServiceImpl) InitDB() {
	s.studentDao.InitDB()
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
