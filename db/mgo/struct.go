package mgo

// DBConfig -
type DBConfig struct {
	Host       string `json:"host" yaml:"host"`             // DB host Ex: 127.0.0.1:27017
	User       string `json:"user" yaml:"user"`             // username Ex: root
	Password   string `json:"password" yaml:"password"`     // password
	Database   string `json:"database" yaml:"database"`     // DB
	SSL        bool   `json:"ssl" yaml:"ssl"`               // 是否設定 mongo ssl(加密)
	AppName    string `json:"appname" yaml:"appname"`       // set Appname
	PoolSize   uint64 `json:"poolsize" yaml:"poolsize"`     // 連接池個數(分散DB請求) default = 5
	ReplicaSet string `json:"replicaset" yaml:"replicaset"` // mongoDB叢集名稱 Ex: rs0
}
