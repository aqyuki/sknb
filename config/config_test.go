package config

import (
	"embed"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/test.yaml
var testConfig embed.FS

func TestLoadFile(t *testing.T) {
	t.Parallel()

	r, err := testConfig.Open("testdata/test.yaml")
	if err != nil {
		t.Fatal(err)
	}

	err = LoadFile(r)
	require.NoError(t, err)
	require.Equal(t, "info", viper.GetString("app.log"))
}
