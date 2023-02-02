package consts

const (
	UserTableName = "user"

	UserServiceName = "usercore"
	UserServiceAddr = ":10000"

	TCP            = "tcp"
	ExprotEndpoint = ":4317"

	ETCDAddress = "127.0.0.1:2379"

	MysqlDefaultDSB = "root:123456@tcp(localhost:3306)/blue_tiktok?charset=utf8&parseTime=True&loc=Local"

	//
	SecretKey   = "Blue_gopher_secret_key"
	IdentityKey = "id"
)
