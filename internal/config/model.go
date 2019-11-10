package config

type Config struct {
	Server   serverConfig   `yaml:"server"`
	Database databaseConfig `yaml:"database"`
}

type serverConfig struct {
	Debug       bool              `yaml:"debug"`
	Port        int               `yaml:"port"`
	Recognition recognitionConfig `yaml:"recognition"`
}

type recognitionConfig struct {
	Threshold     float32 `yaml:"threshold"`
	DlibModelsDir string  `yaml:"dlib_models_dir"`
}

type databaseConfig struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	DBName           string `yaml:"db_name"`
	FacesCollection  string `yaml:"faces_collection"`
	PeopleCollection string `yaml:"people_collection"`
}
