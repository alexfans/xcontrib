package xcontrib

import (
	"entgo.io/ent/entc/gen"
	"strings"
	"text/template"
)

var (
	TemplateFuncs = template.FuncMap{
		"firstLower":          firstLower,
		"concat":              concat,
		"canDelete":           canDelete,
		"NewConditionWrapper": NewConditionWrapper,
	}
)

// FirstLower 字符串首字母小写
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func concat(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

func canDelete(genType *gen.Type) bool {
	for _, field := range genType.Fields {
		fieldName := field.Name
		if fieldName == "is_deleted" || fieldName == "deleted_at" {
			return true
		}
	}
	return false
}

func NewConditionWrapper(condition Condition, genName string) ConditionWrapper {
	return ConditionWrapper{
		Condition: condition,
		GenName:   genName,
	}
}
