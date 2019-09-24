package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Decoder struct {
	raw         []byte
	data        yaml.MapSlice
	isHideData  bool
	replaceData string
}

func NewDecoder(isHideData bool, replaceData string) *Decoder {
	return &Decoder{
		isHideData:  isHideData,
		replaceData: replaceData,
	}
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
	for _, item := range d.data {
		if item.Key != "data" {
			continue
		}

		yamlVals, ok := item.Value.(yaml.MapSlice)
		if !ok {
			return fmt.Errorf("Failed to type assertion: yaml.MapSlice")
		}

		for i, kv := range yamlVals {
			if d.isHideData {
				yamlVals[i].Value = d.replaceData
			} else {
				if kv.Value == nil {
					yamlVals[i].Value = nil
					continue
				}
				dec, err := base64.StdEncoding.DecodeString(kv.Value.(string))
				if err != nil {
					return err
				}
				yamlVals[i].Value = string(dec)
			}
		}
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
