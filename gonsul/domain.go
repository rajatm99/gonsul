package gonsul

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type Gonsul struct {
	Host       string
	Path       string
	File       string
	HttpClient http.Client
}

func (g Gonsul) InitFromFile() (err error) {
	fileData, err := os.ReadFile(g.File)
	if err != nil {
		err = errFileNotFound
		return
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		err = errUnableToReadInputFile
		return
	}
	kv, err := g.GetKV()
	if err != nil && !errors.Is(err, errConsulKeyNotFound) {
		return err
	}

	var flag bool
	for k, v := range data {
		if _, ok := kv[k]; !ok {
			flag = true
			kv[k] = v
		}
	}

	if flag == true {
		finalJSON, _err := json.MarshalIndent(data, "", "    ")
		if _err != nil {
			err = _err
			return
		}

		err = g.putValues(finalJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g Gonsul) GetKV() (kv map[string]interface{}, err error) {
	kv = make(map[string]interface{})

	values, err := g.getValues()
	if err != nil {
		return
	}

	if len(values.Value) == 0 {
		err = errValueNotFound
	}

	decodeString, err := base64.StdEncoding.DecodeString(values.Value)
	if err != nil {
		err = errUnableToReadConsulResponse
		return
	}

	err = json.Unmarshal(decodeString, &kv)
	if err != nil {
		err = errUnableToReadConsulResponse
		return
	}

	return
}
