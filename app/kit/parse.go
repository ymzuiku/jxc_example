package kit

import (
	"encoding/json"
)

func Parse(a interface{}, b interface{}) error {
	if data, err := json.Marshal(a); err == nil {
		if err := json.Unmarshal(data, &b); err != nil {
			return err
		}
	}
	return nil
}
