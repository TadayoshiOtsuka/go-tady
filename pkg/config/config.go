package config

var Config *ProjectConfig

type ProjectConfig struct {
	Name     string
	UserName string
}

func init() {
	Config = &ProjectConfig{}
}
