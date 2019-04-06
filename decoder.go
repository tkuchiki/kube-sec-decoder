package main

import (
	"encoding/base64"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Decoder struct {
	raw  []byte
	data map[interface{}]interface{}
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) loadFile(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	d.raw = b

	return err
}

func (d *Decoder) unmarshalYaml() error {
	return yaml.Unmarshal(d.raw, &d.data)
}

func (d *Decoder) decode() error {
	for key, val := range d.data["data"].(map[interface{}]interface{}) {
		dec, err := base64.StdEncoding.DecodeString(val.(string))
		if err != nil {
			return err
		}

		d.data["data"].(map[interface{}]interface{})[key] = string(dec)
	}

	return nil
}

func (d *Decoder) marshalYaml() (string, error) {
	decoded, err := yaml.Marshal(d.data)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

func (d *Decoder) Decode(r io.Reader) (string, error) {
	err := d.loadFile(r)
	if err != nil {
		return "", err
	}

	err = d.unmarshalYaml()
	if err != nil {
		return "", err
	}

	err = d.decode()
	if err != nil {
		return "", err
	}

	return d.marshalYaml()
}
