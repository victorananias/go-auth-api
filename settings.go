package main

import (
	"encoding/json"
	"os"
)

type Settings struct {
	Db  DbSettings
	Jwt JwtSettings
}

type JwtSettings struct {
	Secret string
}

type DbSettings struct {
	ConnectionString string
}

const settingsName = "settings.json"

func newSettings() (*Settings, error) {
	bytes, err := os.ReadFile(settingsName)
	if err != nil {
		return nil, err
	}
	var result Settings
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
