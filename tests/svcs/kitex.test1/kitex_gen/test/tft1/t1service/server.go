// Code generated by Kitex v0.6.1. DO NOT EDIT.
package t1service

import (
	server "github.com/cloudwego/kitex/server"
	tft1 "kitex/test1/kitex_gen/test/tft1"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler tft1.T1Service, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
