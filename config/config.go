package config

import (
	"embed"
	"fmt"
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

//go:embed env/*.yaml
var envConfigs embed.FS

var envMap = map[string]string{
	"app.env":       "APP_ENV",
	"app.log.level": "APP_LOG_LEVEL",
}

var defaultMap = map[string]any{
	"app.log.level": "info",
}

type Environment string

var runtimeEnv Environment

const (
	EnvLocal = Environment("local")
	EnvTest  = Environment("test")
	EnvProd  = Environment("prod")
)

func IsLocal() bool {
	return runtimeEnv == EnvLocal
}

func IsTest() bool {
	return runtimeEnv == EnvTest
}

func IsProd() bool {
	return runtimeEnv == EnvProd
}

func LoadConfig() error {
	env := Environment(os.Getenv("APP_ENV"))
	if env == "" {
		env = EnvLocal
	}

	return LoadForEnv(env)
}

func LoadForEnv(env Environment) error {
	f, err := envConfigs.Open(fmt.Sprintf("env/%s.yaml", env))
	if err != nil {
		return errors.Wrap(err, "ファイルを開けませんでした")
	}
	defer f.Close()

	return LoadFromFile(f)
}

func LoadFromFile(r io.Reader) error {
	for k, v := range envMap {
		if err := viper.BindEnv(k, v); err != nil {
			return errors.Wrap(err, "環境変数との紐づけに失敗しました")
		}
	}

	for k, v := range defaultMap {
		viper.SetDefault(k, v)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		return errors.Wrap(err, "設定ファイルを読み込めませんでした")
	}

	env := Environment(viper.GetString("app.env"))
	switch env {
	case "":
		env = EnvLocal
		fallthrough
	case EnvLocal, EnvTest, EnvProd:
		runtimeEnv = env
	default:
		return errors.New("不明な実行環境です")
	}

	return nil
}
