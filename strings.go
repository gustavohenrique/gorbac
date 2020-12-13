package gorbac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func FromFile(path string) ([]Role, error) {
	var roles []Role
	var err error
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return roles, err
	}
	if endsWith(path, ".json") {
		err = json.Unmarshal(f, &roles)
	} else if endsWith(path, ".yaml") || endsWith(path, ".yml") {
		err = yaml.Unmarshal(f, &roles)
	} else {
		err = fmt.Errorf("%s is not a JSON or YAML file", path)
	}
	return roles, err
}

func FromJSON(i interface{}) ([]Role, error) {
	var roles []Role
	if i == nil {
		return roles, fmt.Errorf("Argument is nil")
	}
	s := fmt.Sprintf("%s", i)
	err := json.Unmarshal([]byte(s), &roles)
	return roles, err
}

func FromYAML(i interface{}) ([]Role, error) {
	var roles []Role
	if i == nil {
		return roles, fmt.Errorf("Argument is nil")
	}
	s := fmt.Sprintf("%s", i)
	err := yaml.Unmarshal([]byte(s), &roles)
	return roles, err
}

func endsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
