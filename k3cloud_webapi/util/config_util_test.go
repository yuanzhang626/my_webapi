package util

import "testing"

func TestInitConfig(t *testing.T) {
	configPath := "./conf.ini"
	configNodeIn := "config"

	conf, err := InitConfig(configPath, configNodeIn)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("conf:%+v", conf)
}
