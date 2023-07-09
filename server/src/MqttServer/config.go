package MqttServer

import (
	"io/ioutil"

	"github.com/stretchr/objx"
)

var (
	config objx.Map
)

func ConfigInit() error {

	fBuf, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	jmap, err := objx.FromJSON(string(fBuf))
	if err != nil {
		return err
	}
	config = jmap
	return nil
}

func ConfigGet(seg string) string {
	return config.Get(seg).Str()
}
