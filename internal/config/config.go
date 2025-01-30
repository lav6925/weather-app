package config

// Config holds all application configuration.
type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SslMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Weather struct {
		APIURL      string `mapstructure:"weather_api_url"`
		APIKey      string `mapstructure:"weather_api_key"`
		RefreshTime int    `mapstructure:"weather_refresh_time"`
	} `mapstructure:"weather"`

	Server struct {
		Port     int `mapstructure:"port"`
		HTTPPort int `mapstructure:"httpPort"`
	} `mapstructure:"server"`
}
