type {{.Name}} struct {
    {{- range $i,$v := .Fields}}
     {{$v.Name}} {{ $v.Type}} `json:"{{$v.Tag}},omitempty"`
    {{- end}}
}