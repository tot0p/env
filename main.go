package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadMap Function to load the environment into a map
func LoadMap() (map[string]string, error) {
	return LoadMapPath(".env")
}

// LoadMapPath Function to load the environment into a map
func LoadMapPath(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	datas := string(data)
	sep := "\n"
	strings.ReplaceAll(datas, "\r", "")
	var result = map[string]string{}
	for _, elem := range strings.Split(datas, sep) {
		if strings.Contains(elem, "=") {
			t := strings.SplitN(elem, "=", 2)
			if len(t) > 1 {
				result[t[0]] = t[1]
			} else {
				return nil, errors.New("error : syntax error")
			}
		}
	}
	return result, nil
}

// Load Function to load the environment in os Env
func Load() error {
	return LoadPath(".env")
}

// LoadPath Function to load the environment in os Env
func LoadPath(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	datas := string(data)
	sep := "\n"
	strings.ReplaceAll(datas, "\r", "")
	var result = map[string]string{}
	for _, elem := range strings.Split(datas, sep) {
		if strings.Contains(elem, "=") {
			t := strings.SplitN(elem, "=", 2)
			if len(t) > 1 {
				result[t[0]] = t[1]
			} else {
				return errors.New("error : syntax error")
			}
		}
	}
	for key, value := range result {
		os.Setenv(key, value)
	}
	return nil
}

// Get Function to get the environment
func Get(key string) string {
	return os.Getenv(key)
}

// GetInt Function to get the environment as int
func GetInt(key string) (int, error) {
	return strconv.Atoi(os.Getenv(key))
}

// GetBool Function to get the environment as bool
func GetBool(key string) (bool, error) {
	return strconv.ParseBool(os.Getenv(key))
}

// GetFloat Function to get the environment as float
func GetFloat(key string) (float64, error) {
	return strconv.ParseFloat(os.Getenv(key), 64)
}
