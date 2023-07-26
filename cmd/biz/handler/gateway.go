package handler

import (
	"context"
	"encoding/json"

	kxcliprovider "apigateway/KxCliProvider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/tidwall/gjson"
)

func NewGenericClient(destServiceName string) genericclient.Client {
	genericCli, _ := genericclient.NewClient(destServiceName, generic.BinaryThriftGeneric())
	return genericCli
}

func ForwardPOST(ctx context.Context, c *app.RequestContext) {
	var err error
	// var respStruct = make(map[string]interface{})

	serviceName := c.Param("svc")
	methodName := c.Param("method")

	httpReq, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		panic(err)
	}
	tReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic(err)
	}

	cli := kxcliprovider.GetGenericCli(serviceName)

	reqJson, err := json.Marshal(tReq.Body)
	if err != nil {
		panic(err)
	}

	resp, err := cli.GenericCall(ctx, methodName, string(reqJson))
	if err != nil {
		panic(err)
	}

	// reqS := string(c.Request.Body())
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := cli.GenericCall(ctx, methodName, reqS)
	// if err != nil {
	// 	panic(err)
	// }

	m, _ := gjson.Parse(resp.(string)).Value().(map[string]interface{})

	// err = json.Unmarshal([]byte(resp.(string)), &respStruct)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("m: %v\n", m)

	c.JSON(consts.StatusOK, m)
}

func ForwardGET(ctx context.Context, c *app.RequestContext) {
	var err error
	var respStruct = make(map[string]interface{})

	serviceName := c.Param("svc")
	methodName := c.Param("method")

	httpReq, err := adaptor.GetCompatRequest(&c.Request)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	cli := kxcliprovider.GetGenericCli(serviceName)

	//TODO: fit multi value for one param
	for k, v := range httpReq.URL.Query() {
		respStruct[k] = v[0]
	}

	reqJson, err := json.Marshal(respStruct)
	jsonS := string(reqJson)
	if err != nil {
		panic(err)
	}

	resp, err := cli.GenericCall(ctx, methodName, string(jsonS))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(resp.(string)), &respStruct)
	if err != nil {
		panic(err)
	}

	c.JSON(consts.StatusOK, respStruct)
}
