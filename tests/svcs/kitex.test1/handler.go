package main

import (
	"context"
	"encoding/json"
	"fmt"
	tft1 "kitex/test1/kitex_gen/test/tft1"
)

// T1ServiceImpl implements the last service interface defined in the IDL.
type T1ServiceImpl struct{}

// Tst implements the T1ServiceImpl interface.
func (s *T1ServiceImpl) Tst(ctx context.Context, t1 *tft1.T1) (resp *tft1.T1, err error) {
	resp = &tft1.T1{
		Name: "hello",
	}
	return
}

func (s *T1ServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {

	fmt.Printf("\"hello\": %v\n", "hello")

	reqS := request.(string)
	switch method {
	case "Tst":
		var t1 tft1.T1
		err = json.Unmarshal([]byte(reqS), &t1)
		if err != nil {
			panic(err)
		}
		response, err = s.Tst(ctx, &t1)
		if err != nil {
			panic(err)
		}
	}
	byteResp, err := json.Marshal(response)
	response = string(byteResp)

	return
}
