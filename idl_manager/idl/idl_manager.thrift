// idl/hello.thrift
namespace go idlmanager


struct SetIdlPathReq{
    1: string IdlPath(api.body="idlpath")
    2: string GateWayPath(api.body="gatewaypath")

}

struct IdlService {
    1: string Name (api.body="svcname");
    2: string FileName (api.body="idlfile");
    3: string FileContent (api.body="idlcontent");
}

struct ChangeServiceReq {
    1: string Name (api.body="svcname");
    2: string FileContent (api.body="idlcontent")
}

struct DeleteServiceReq {
    1: string Name (api.body="svcname");
}

struct GetServiceReq {
    1: string Name (api.query="svcname");
}

struct ManageServiceResp {
    1: string RespBody;
    2: optional list<string> Items
}


service IDLManageService {
    ManageServiceResp SetIdlPath(1:SetIdlPathReq request) (api.post="manage/Path")
    ManageServiceResp AddService(1:IdlService request) (api.post="/manage/Add")
    ManageServiceResp ChangeService(1:ChangeServiceReq request) (api.post="/manage/Change", api.serializer = 'form')
    ManageServiceResp DeleteService(1:DeleteServiceReq request) (api.post="/manage/Delete")
    ManageServiceResp GetService(1:GetServiceReq request) (api.get="/manage/Get")
    ManageServiceResp ListServices() (api.get="/manage/List")
}
