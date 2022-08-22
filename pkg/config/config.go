package config

var Config *ProjectConfig

type ProjectConfig struct {
	Name           string
	UserName       string
	TargetPreset   string
	OldPackageName string
}

func init() {
	Config = &ProjectConfig{}
}
