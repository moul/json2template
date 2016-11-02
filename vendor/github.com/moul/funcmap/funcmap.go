package funcmap

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"json": func(v interface{}) string {
		a, _ := json.Marshal(v)
		return string(a)
	},
	"prettyjson": func(v interface{}) string {
		a, _ := json.MarshalIndent(v, "", "  ")
		return string(a)
	},
	// yaml
	// xml
	// toml
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
	"int": func(v interface{}) string {
		a, err := strconv.Atoi(v.(string))
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return fmt.Sprintf("%d", a)
	},
}
