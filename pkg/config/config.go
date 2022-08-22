package config

var Config *ProjectConfig

type ProjectConfig struct {
	Name         string
	UserName     string
	TargetPreset string
}

func init() {
	Config = &ProjectConfig{}
}
