BinaryName=apigateway

etcd:
	etcd --log-level debug

kitex:
	rm tests/svcs/kitex.demo/foo.db
	cd tests/svcs/kitex.demo && chmod +x build.sh && ./build.sh && ./output/bootstrap.sh

test1:
	cd tests/svcs/kitex.test1 && chmod +x build.sh && ./build.sh && ./output/bootstrap.sh

hertz:
	cd cmd && go build -o ./${BinaryName} && exec ./${BinaryName}

idlManager:
	cd idl_manager && go build -o ./idl_manager && exec ./idl_manager