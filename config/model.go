package config

type Config struct {
	SwagCfg SwagSyncConfig `yaml:"swag-sync"`
}

type SwagSyncConfig struct {
	SourceDirs       []string `yaml:"srcDirs"`
	IgnoreFileNames  []string `yaml:"ignoreFiles"`
	IncludeFileNames []string `yaml:"includeFiles"`
	PkgName          string   `yaml:"tarPkgName"`
	TarDir           string   `yaml:"tarDir"`
}
