package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type Field struct {
	Name string
	Type string
	Tag  string
	Eee  int
}

type StructInfo struct {
	Name   string
	Fields []Field
}

func (this *StructInfo) IsSimilar(m *StructInfo) bool {
	cnt := 0
	temp := make([]Field, 0, 10)
	for _, v := range m.Fields {
		flag := false
		for _, v2 := range this.Fields {
			if v.Name == v2.Name && v.Type == v2.Type {
				cnt++
				flag = true
				break
			}
		}
		if !flag {
			temp = append(temp, v)
		}
	}

	//cnt 相同字段 
	if cnt >= len(this.Fields)*2/3 {
		this.Fields = append(this.Fields, temp...)
		return true
	}
	return false
}

func (this *StructInfo) Parser2File(f *os.File) {
	b, _ := ioutil.ReadFile("template/struct.tpl")
	tmpl, _ := template.New("test").Parse(string(b))
	tmpl.Execute(f, *this)
}

type global map[string]*StructInfo

var globalStruct global

func (this *global) ToFile(file string) {
	f, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	for _, v := range *this {
		sort.Slice(v.Fields, func(i, j int) bool {
			r := strings.Compare(v.Fields[i].Name, v.Fields[j].Name)
			if r > 0 {
				return false
			} else {
				return true
			}
		})
		v.Parser2File(f)
		f.WriteString("\r\n ")
	}
	f.Close()
}

func EnterGlobalStruct(m *StructInfo) string {
	for _, v := range globalStruct {
		if v.IsSimilar(m) {
			return v.Name
		}
	}
	globalStruct[m.Name] = m
	return m.Name
}

func init() {
	globalStruct = make(map[string]*StructInfo)
}

func Json2Struct(b []byte, file string) {
	base := make(map[string]interface{}, 1)
	json.Unmarshal(b, &base)
	HandleMap("key", base)
}

func Test(value map[string]interface{}) {
	for k, v := range value {
		tt := reflect.TypeOf(v)
		fmt.Println(k, tt.Kind().String())
	}
}

func GetValueType(k string, v interface{}) string {
	tt := reflect.TypeOf(v)
	name := ""
	switch tt.Kind() {
	case reflect.String, reflect.Float64, reflect.Int, reflect.Bool:
		name = tt.Kind().String()
	case reflect.Map:
		mapV, _ := v.(map[string]interface{})
		name = HandleMap(k, mapV)
		fmt.Println(k, "map", name)
	case reflect.Array, reflect.Slice:
		vv := reflect.ValueOf(v)
		if vv.Len() == 0 {
			name = "[]interface{}"
		} else {
			zero := (vv.Index(0).Interface())
			name = "[]" + GetValueType(k, zero)
		}
	}
	return name
}

func HandleMap(key string, value map[string]interface{}) string {
	res := StructInfo{
		Name:   UnderscoreToUpperCamelCase(key),
		Fields: make([]Field, 0, 10),
	}
	t := ""
	for k, v := range value {
		name := GetValueType(k, v)

		res.Fields = append(res.Fields, Field{
			Name: UnderscoreToUpperCamelCase(k),
			Type: name,
			Tag:  CamelCaseToUnderscore(k),
		})

		if t == "" {
			t = name
		} else if t != name {
			t = "no"
		}
	}
	if t != "no" && t != "" {
		return fmt.Sprintf("map[string]%s", t)
	} else {
		// fmt.Println(res.Name, res)
		return EnterGlobalStruct(&res)
	}
}

func main() {
	// t := TestS{
	// 	Name:  "aaa",
	// 	Age:   10,
	// 	MM:    1.1,
	// 	TT:    make(map[string]Field, 11),
	// 	Array: make([]Field, 11),
	// 	SS:    []int{1},
	// 	FF:    Field{},
	// }
	// // t := map[string]int{
	// // 	"aaa": 1,
	// // 	"cc":  2,
	// // 	"bb":  21,
	// // }
	// t.TT["m_field"] = Field{Name: "111"}
	// t.TT["m_field2"] = Field{Name: "111"}
	// t.TT["m_field34"] = Field{Name: "2000"}
	// b, _ := json.Marshal(t)
	// Json2Struct(b, "")
	// for k, v := range globalStruct {
	// 	fmt.Println(k, v)
	// }
	// matchid := "5682028409"
	// url := fmt.Sprintf("http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1?match_id=%s&key=E4864EF08820F5E4F2FDEE9A63658E20", matchid)

	// resp, _ := http.Get(url)
	// b, _ := ioutil.ReadAll(resp.Body)

	filepath := path.Join("file", "sample.json")
	b, _ := ioutil.ReadFile(filepath)
	fmt.Println(string(b))
	Json2Struct(b, "")

	globalStruct.ToFile("temp3.go")
}

type TestS struct {
	Name  string           `json:"first_name"`
	Age   int              `json:"my_age"`
	MM    float64          `json:"mm"`
	TT    map[string]Field `json:"good"`
	SS    []int
	FF    Field
	Array []Field `json:"tt_slice"`
}

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output, '_')
		}

		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
