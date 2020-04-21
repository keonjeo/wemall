module wemall/controller/common

go 1.13

require (
	config v0.0.0-00010101000000-000000000000
	github.com/geekypanda/httpcache v0.0.1 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/kataras/go-errors v0.0.3 // indirect
	github.com/kataras/iris/v12 v12.1.8
	github.com/klauspost/compress v1.10.4 // indirect
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/valyala/fasthttp v1.11.0 // indirect
	gopkg.in/kataras/iris.v6 v6.0.0-20170603191843-09a2066268f9
	gopkg.in/yaml.v2 v2.2.8 // indirect
	model v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace (
	config => ./../../config
	model => ./../../model
	utils => ./../../utils
)
