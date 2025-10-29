package config

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Env struct {
	DB DB
}

type DB struct {
	DSN string
}

func New() (*Env, error) {
	file, err := os.Open("config/config.env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// upload all key-value pairs into a map
	vars := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		vars[key] = value
	}

	env := &Env{}

	// fill Env struct from the map
	err = fillStructFromMap(reflect.ValueOf(&env.DB).Elem(), vars)
	if err != nil {
		return nil, err
	}

	return env, scanner.Err()
}

func fillStructFromMap(v reflect.Value, vars map[string]string) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// field name must match the key in the map
		if val, ok := vars[field.Name]; ok {
			value.SetString(val)
		} else {
			return fmt.Errorf("key does not match any field or is missing: %s", field.Name)
		}
	}

	return nil
}
