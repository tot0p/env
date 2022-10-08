package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// LoadEnvMap Function to load the environment into a map
func LoadEnvMap(path string) (map[string]string, error) {
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

// LoadEnv Function to load the environment in os Env
func LoadEnv(path string) error {
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
