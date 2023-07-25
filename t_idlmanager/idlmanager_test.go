package unitTests

import (
	idlprovider "apigateway/IDLProvider"
	idlmanager "apigateway/biz/model/idlmanager"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddIDL(t *testing.T) {
	var req idlmanager.AddServiceReq
	req.Name = "test_service"
	req.FileName = "test_file"
	req.FileContent = []byte("test_content")

	err := idlprovider.AddIDL(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal()

}
func TestChangeIDL(t *testing.T) {
	type args struct {
		req idlmanager.ChangeServiceReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{req.Name}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := idlmanagerF.ChangeIDL(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ChangeIDL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteIDL(t *testing.T) {
	type args struct {
		req idlmanager.DeleteServiceReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := idlmanagerF.DeleteIDL(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DeleteIDL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
