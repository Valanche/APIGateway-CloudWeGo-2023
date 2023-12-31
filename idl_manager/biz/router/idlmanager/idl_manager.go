// Code generated by hertz generator. DO NOT EDIT.

package idlmanager

import (
	idlmanager "api/idl_manager/biz/handler/idlmanager"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_manage := root.Group("/manage", _manageMw()...)
		_manage.POST("/Add", append(_addserviceMw(), idlmanager.AddService)...)
		_manage.POST("/Change", append(_changeserviceMw(), idlmanager.ChangeService)...)
		_manage.POST("/Delete", append(_deleteserviceMw(), idlmanager.DeleteService)...)
		_manage.GET("/Get", append(_getserviceMw(), idlmanager.GetService)...)
		_manage.GET("/List", append(_listservicesMw(), idlmanager.ListServices)...)
		_manage.POST("/Path", append(_setidlpathMw(), idlmanager.SetIdlPath)...)
	}
}
