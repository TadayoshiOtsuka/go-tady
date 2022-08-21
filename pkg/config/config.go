package config

var Config *ProjectConfig

type ProjectConfig struct {
	Name           string
	UserName       string
	TargetTemplate string
}

func init() {
	Config = &ProjectConfig{}
}
