// +build js

package imports

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Imports() error {
	res, err := http.Get("pkg/imports.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("imports: %s", res.Status)
	}
	decode := json.NewDecoder(res.Body).Decode
	m := map[string]string{}
	if err = decode(&m); err != nil {
		return err
	}
	AddImports(m)
	return nil
}
