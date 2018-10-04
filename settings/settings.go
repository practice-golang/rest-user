package settings

const (
	// DbHost : 주소
	DbHost = "localhost"
	// DbPort : 포트
	DbPort = "5432"
	// DbUser : DB 계정
	DbUser = "root"
	// DbPassword : DB 비번
	DbPassword = ""
	// DbName : DB명
	DbName = "postgres"

	// UserTable : users 테이블명
	UserTable = "users"

	// JwtSigningMethod : JWT signing method
	JwtSigningMethod = "HS256"
)

var (
	// JwtSigningKey : JWT Key
	JwtSigningKey = []byte("mySecret")
)
