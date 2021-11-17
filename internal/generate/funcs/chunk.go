package funcs

import (
	"github.com/chyroc/godash/internal"
)

var Funcs = map[string]func(string) string{
	"chunk": generateChunk,
}

func generateChunk(typ string) string {
	camelCaseTyp := internal.ToCamelCase(typ)
	var tpl = `func Chunk{{ .CamelCaseTyp }}(array []{{ .Typ }}, size int) [][]{{ .Typ }} {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]{{ .Typ }}{}
	}

	index := 0
	resIndex := 0

	result := make([][]{{ .Typ }}, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}`
	return internal.BuildTemplate(tpl, internal.TemplateData{
		"CamelCaseTyp": camelCaseTyp,
		"Typ":          typ,
	})
}
