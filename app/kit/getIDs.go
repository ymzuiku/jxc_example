package kit

import (
	"github.com/mitchellh/mapstructure"
)

func GetIDs(list interface{}, key string) ([]int32, error) {
	obj := []map[string]interface{}{}
	if err := mapstructure.Decode(list, &obj); err != nil {
		return nil, err
	}
	ids := make([]int32, 0, len(obj))

	for _, item := range obj {
		v := item[key].(int32)
		ids = append(ids, v)
	}
	return ids, nil
}
