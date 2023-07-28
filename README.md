# APIGateway-CloudWeGo-2023

----> [项目说明_zh Click here](/documents/项目说明.md) <----

This is a simple API Gateway based on Hertz and Kitex framework.
Works on **Linux**.

The API Gateway handles incoming **HTTP POST** (**JSON** format) or **HTTP GET** requests and sends corresponding Thrift RPC requests to target services. 
Responses from the services will be returned to the client in the format of **JSON**.

The API Gateway uses **etcd** for service discovery, so you should have etcd installed first.



## Deploy
### 1 - Clone this repo to your deploy path:

```
git clone https://github.com/Valanche/APIGateway-CloudWeGo-2023.git
```
### 2 - Start etcd:

You can start etcd in the same host with the API Gateway.

Start a new terminal, then use command: 
```
etcd
```

### 3 - Start the API Gateway:
Start a new terminal, then use command: 
```
// in dir APIGateway-CloudWeGo-2023
make hertz
```

### (Optional) 4 - Check
You can check whether the API Gateway is ready using the following command:

```
curl localhost:8888/ping
```

The output should be: 
```
{"message":"pong"}
```
### (Optional) 5 - Start the test service
Start a new terminal, then use command: 
```
// in dir APIGateway-CloudWeGo-2023
make kitex
```
This starts a test service. You can use the following commands to access the service through the gateway and see the result:
```
curl --request GET 'http://localhost:8888/api/kitex.demo/Query?id=1'
```
```
curl --request POST 'http://localhost:8888/api/kitex.demo/Register' -d '{
    "sex": "non-binary",
    "college": {
        "address": "nowhere",
        "name": "noname3"
    },
    "email": [
        "noname2@nohost.com"
    ],
    "id": 1,
    "name": "noname1"
}'
```
```
curl --request GET 'http://localhost:8888/api/kitex.demo/Query?id=1'
```
For more info please refer to the idl file.

## Usage
If you want to put this into use, you should do the following things:

### Configure etcd address:
If your etcd center is deployed somewhere else, you should configure the target address in:

```
APIGateway-CloudWeGo-2023/KxCliProvider/provider.go
```

After that you should restart the gateway.

### Add IDL files:
Add the thrift idl files of your services to directory: 
```
APIGateway-CloudWego-2023/cmd/idl/
```

### Match the names and the files
Add the names of your services and the path of its idl file to :
```
APIGateway-CloudWego-2023/cmd/idl/svcPath
```
The format is:
    **serviceName + ", " + idlFilePath ( "./idl/" + idlFileName)**.

For example, I have a service named `kitex.demo`, 

and it's idl file is `./idl/kxServer.thrift`.

so there should be a line `kitex.demo, ./idl/kxServer.thrift` in the file `svcpath`.
### Requests URL
Requests should be sent to :
```
$gatewayAddress:8888/api/$serviceName/$methodName
```
For example, if the API Gateway is deployed on the host whose ip address is `1.1.1.1`, 

and the client want to access the `Register` method of the `kitex.demo` service, 

the request should be sent to:
```
1.1.1.1:8888/api/kitex.demo/Register
```