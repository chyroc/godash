package funcs

import (
	"github.com/chyroc/godash/internal"
)

func generateSet(typ string) string {
	camelCaseTyp := internal.ToCamelCase(typ)
	tpl := `type Set{{ .CamelCaseTyp }} map[{{ .Typ }}]struct{}

func NewSet{{ .CamelCaseTyp }}() Set{{ .CamelCaseTyp }} {
	return make(map[{{ .Typ }}]struct{})
}

// Diff return new set with all items in s but not in o
func (s Set{{ .CamelCaseTyp }}) Diff(o Set{{ .CamelCaseTyp }}) Set{{ .CamelCaseTyp }} {
	result := make(Set{{ .CamelCaseTyp }}, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s Set{{ .CamelCaseTyp }}) Inter(o Set{{ .CamelCaseTyp }}) Set{{ .CamelCaseTyp }} {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(Set{{ .CamelCaseTyp }}, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s Set{{ .CamelCaseTyp }}) Union(o Set{{ .CamelCaseTyp }}) Set{{ .CamelCaseTyp }} {
	result := make(Set{{ .CamelCaseTyp }}, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s Set{{ .CamelCaseTyp }}) Slice() []{{ .Typ }} {
	result := make([]{{ .Typ }}, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s Set{{ .CamelCaseTyp }}) SortedSlice() []{{ .Typ }} {
	slice := list{{ .CamelCaseTyp }}(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s Set{{ .CamelCaseTyp }}) ReversedSortedSlice() []{{ .Typ }} {
	slice := list{{ .CamelCaseTyp }}(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s Set{{ .CamelCaseTyp }}) Add(id {{ .Typ }}) Set{{ .CamelCaseTyp }} {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s Set{{ .CamelCaseTyp }}) Del(id {{ .Typ }}) Set{{ .CamelCaseTyp }} {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s Set{{ .CamelCaseTyp }}) IsExist(id {{ .Typ }}) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s Set{{ .CamelCaseTyp }}) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s Set{{ .CamelCaseTyp }}) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s Set{{ .CamelCaseTyp }}) DelSet(o Set{{ .CamelCaseTyp }}) Set{{ .CamelCaseTyp }} {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s Set{{ .CamelCaseTyp }}) AddSlice(ids []{{ .Typ }}) Set{{ .CamelCaseTyp }} {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s Set{{ .CamelCaseTyp }}) AddSet(o Set{{ .CamelCaseTyp }}) Set{{ .CamelCaseTyp }} {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type list{{ .CamelCaseTyp }} []{{ .Typ }}

func (s list{{ .CamelCaseTyp }}) Len() int           { return len(s) }
func (s list{{ .CamelCaseTyp }}) Less(i, j int) bool { return s[i] < s[j] }
func (s list{{ .CamelCaseTyp }}) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
`
	return internal.BuildTemplate(tpl, internal.TemplateData{
		"CamelCaseTyp": camelCaseTyp,
		"Typ":          typ,
	})
}
