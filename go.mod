module github.com/dockerian/go-coding

go 1.12

require (
	github.com/aws/aws-sdk-go v1.33.0
	github.com/dockerian/dateparse v0.0.0-20171201195521-0d9501caeeb5
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.2.0
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.1.0 // indirect
	github.com/russross/blackfriday v1.5.2
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.5.1
	github.com/urfave/negroni v1.0.0
	golang.org/x/text v0.3.2
	gopkg.in/yaml.v2 v2.2.4
)

replace (
	github.com/dockerian/go-coding => ../go-coding
	github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2
)
