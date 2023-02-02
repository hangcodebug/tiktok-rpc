service:
	kitex --thrift-plugin validator -module github.com/BlueGopher/tiktok-rpc idl/usercore.thrift # execute in the project root directory
	kitex --thrift-plugin validator -module github.com/BlueGopher/tiktok-rpc -service userCore -use github.com/BlueGopher/tiktok-rpc/kitex_gen ../../idl/usercore.thrift # execute in cmd/user