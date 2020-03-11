package meta

// config
type FilePath struct {
	OriginalPath    string `toml:"original_path"`
	PidFileLocation string `toml:"pid_file_location"`
}

type Store struct {
	StoreType      string `toml:"store_type"`
	StoreHost      string `toml:"store_host"`
	StorePort      int    `toml:"store_port"`
	StoreNameSpace string `toml:"store_namespace"`
	StoreUser      string `toml:"store_user"`
	StorePassWord  string `toml:"store_password"`
}

type HttpConf struct {
	Addr string `toml:"addr"`
	Port int    `toml:"port"`
}
