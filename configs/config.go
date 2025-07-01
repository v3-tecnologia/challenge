package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	IPLimitMaxReq  int    `mapstructure:"IP_LIMIT_MAX_REQUEST"`
	IPBlockTimeSec int    `mapstructure:"IP_BLOCK_TIME_SECONDS"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBAddress      string `mapstructure:"DB_ADDRESS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	WebServerPort  string `mapstructure:"WEBSERVER_PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JWTExperesIn   int    `mapstructure:"JWT_EXPERES_IN"`
	TokenAuth      *jwtauth.JWTAuth
}

func Load(path string) (*conf, error) {
	var conf *conf
	viper.SetConfigType("env")
	viper.AddConfigPath("path")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	conf.TokenAuth = jwtauth.New("HS256", []byte(conf.JWTSecret), nil)
	return conf, nil
}
