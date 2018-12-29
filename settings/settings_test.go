package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSettings(t *testing.T){

	t.Run("test init", func(t *testing.T){
		Init()
		got := settings
		assert.NotEmpty(t, got.PrivateKeyPath)
		assert.NotEmpty(t, got.PublicKeyPath)
		assert.NotEmpty(t, got.JWTExpirationDelta)
	})
	t.Run("test load from environment", func (t *testing.T){
		expected := Settings{
			"settings/local/keys/private_key",
			"settings/local/keys/public_key.pub",
			72,
		}
		LoadSettingsByEnv("local")
		got := settings
		assert.Equal(t, expected, got)
	})
	t.Run("get environment", func(t *testing.T){
		assert.NotEmpty(t, GetEnvironment())
	})
}
