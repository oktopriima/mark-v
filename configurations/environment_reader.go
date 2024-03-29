package configurations

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetStrings(key string) []string
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	Init(prefix, ext string)
}

type viperConfig struct{}

func (v *viperConfig) Init(prefix, ext string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(ext)
	viper.SetConfigFile(env + `.` + ext)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetStringSlice(key string) (c []string) {
	c = viper.GetStringSlice(key)
	return
}

func (v *viperConfig) GetStrings(key string) (c []string) {
	val := viper.GetString(key)
	c = strings.Split(val, ",")
	return
}

func (v *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func NewConfig(ext string) Config {
	v := &viperConfig{}
	v.Init("", ext)
	return v
}

func NewWithPrefix(ext string) Config {
	v := &viperConfig{}
	prefix := v.GetString("prefix")
	v.Init(prefix, ext)
	return v
}
