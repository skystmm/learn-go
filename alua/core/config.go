package core

type Config struct {
	maxGogoutine  int
	taskQueneType int
}

//读取配置文件
func (config *Config) LoadConfig(path string) {
	if path == "" {

	}
}
