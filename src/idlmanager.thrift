// idl/hello.thrift
namespace go idlmanager

struct AddServiceReq {
    1: string Name (api.body="svcname");
    2: binary File (api.body="idlfile");
}

struct ChangeServiceReq {
    1: string Name (api.body="svcname");
    2: string FileName (api.body="filename")
    3: binary FileContent (api.body="idlfile")
}

struct DeleteServiceReq {
    1: string Name (api.body="svcname");
}

struct ManageServiceResp {
    1: string RespBody;
}


service IDLManageService {
    ManageServiceResp AddIDL(1:AddServiceReq request) (api.post="/manage/Add")
    ManageServiceResp ChangeIDL(1:ChangeServiceReq request) (api.post="/manage/Change", api.serializer = 'form')
    ManageServiceResp DeleteIDL(1:DeleteServiceReq request) (api.post="/manage/Delete")
}
