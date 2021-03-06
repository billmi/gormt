package gtools

const (
	_tagGorm = "gorm"
	_tagJSON = "json"
)

//TypeDicMp 精确匹配类型
var TypeDicMp = map[string]string{
	"int":                 "int",
	"bigint":              "int64",
	"varchar":             "string",
	"char":                "string",
	"datetime":            "time.Time",
	"tinyint(1)":          "bool",
	"tinyint(1) unsigned": "bool",
	"json":                "string",
	"text":                "string",
	"timestamp":           "time.Time",
}

//TypeMatchMp 模糊匹配类型
var TypeMatchMp = map[string]string{
	`int[(]\d+[)]`:     "int",
	`bigint[(]\d+[)]`:  "int64",
	`varchar[(]\d+[)]`: "string",
	`char[(]\d+[)]`:    "string",
}
