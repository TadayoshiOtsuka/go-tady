package config

var Config *ProjectConfig

type ProjectConfig struct {
	Name string
}

func init() {
	Config = &ProjectConfig{}
}
