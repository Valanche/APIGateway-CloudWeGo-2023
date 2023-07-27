namespace go test.tft1

struct T1 {
    1: string name(api.body="name"),
}
service T1Service {
    T1 Tst(1: T1 t1) (api.get="/hello");
}