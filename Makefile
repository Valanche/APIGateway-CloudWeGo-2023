BinaryName=apigateway

etcd:
	etcd --log-level debug

kitex:
	cd tests/svcs/kitex.demo && chmod +x build.sh && ./build.sh && ./output/bootstrap.sh

hertz:
	cd cmd && go build -o ./${BinaryName} && exec ./${BinaryName}