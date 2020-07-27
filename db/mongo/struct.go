package mongo

// DBConfig -
type DBConfig struct {
	Host     string
	User     string
	Password string
	Database string
	PoolSize uint64  // connection number
}
