// idl/hello.thrift
namespace go idlmanager

struct AddServiceReq {
    1: string Name (api.body="svcname");
    2: string FileName (api.body="idlfile");
    3: binary FileContent (api.body="idlfile");
}

struct ChangeServiceReq {
    1: string Name (api.body="svcname");
    2: binary FileContent (api.body="idlfile")
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
    ManageServiceResp AddIDL(1:AddServiceReq request) (api.post="/manage/Add")
    ManageServiceResp ChangeIDL(1:ChangeServiceReq request) (api.post="/manage/Change", api.serializer = 'form')
    ManageServiceResp DeleteIDL(1:DeleteServiceReq request) (api.post="/manage/Delete")
    ManageServiceResp GetIDL(1:GetServiceReq request) (api.get="/manage/Get")
    ManageServiceResp ListIDL() (api.get="/manage/List")
}
