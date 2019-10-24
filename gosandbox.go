package gosandbox

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type UserFunc func(data Data, cancel <-chan struct{}) (Data, error)

type Data struct {
	json string
}

func New(json string) Data {
	return Data{
		json: json,
	}
}

func (d *Data) Unset(path string) error {
	json, err := sjson.Delete(d.json, path)
	if err == nil {
		d.json = json
	}

	return err
}

func (d *Data) Set(path string, v string) error {
	json, err := sjson.Set(d.json, path, v)
	if err == nil {
		d.json = json
	}

	return err
}

func (d *Data) SetInt64(path string, v int64) error {
	json, err := sjson.Set(d.json, path, v)
	if err == nil {
		d.json = json
	}

	return err
}

func (d *Data) GetInt64(path string) int64 {
	return gjson.Get(d.json, path).Int()
}

func (d *Data) Get(path string) string {
	return gjson.Get(d.json, path).String()
}

func (d *Data) String() string {
	return d.json
}


