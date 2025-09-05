// Package config はアプリケーション内の設定に関する機能を提供する
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

type Environment string

const (
	EnvLocal = Environment("local")
	EnvTest  = Environment("test")
	EnvProd  = Environment("prod")
)

var runtimeEnv Environment

var envMap = map[string]string{
	"app.env": "APP_ENV",
	"app.log": "APOP_LOG",
}

var defaultMap = map[string]any{
	"app.log": "info",
}

func IsLocal() bool {
	return runtimeEnv == EnvLocal
}

func LoadConfig() error {
	env := Environment(os.Getenv("APP_ENV"))
	if env == "" {
		env = EnvLocal
	}

	return LoadForEnv(env)
}

func LoadForEnv(env Environment) error {
	r, err := envConfigs.Open(fmt.Sprintf("env/%s.yaml", string(env)))
	if err != nil {
		return errors.Wrap(err, "ファイルを開けませんでした")
	}
	defer r.Close()

	return LoadFile(r)
}

func LoadFile(r io.Reader) error {
	for k, v := range envMap {
		if err := viper.BindEnv(k, v); err != nil {
			return errors.Wrap(err, "環境変数と関連付けできませんでした")
		}
	}

	for k, v := range defaultMap {
		viper.SetDefault(k, v)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		return errors.Wrap(err, "設定ファイルを読み込めませんでした")
	}

	switch env := Environment(viper.GetString("app.env")); env {
	case "":
		viper.Set("app.env", EnvLocal)
		env = EnvLocal

		fallthrough

	case EnvLocal, EnvTest, EnvProd:
		runtimeEnv = env

	default:
		return errors.New("不明な実行環境が指定されました")
	}

	return nil
}
