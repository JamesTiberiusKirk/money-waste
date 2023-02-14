package site

import (
	"encoding/json"
)

func stringyfyJSON(object any) string {
	stringObject, _ := json.Marshal(object)
	return string(stringObject)
}
