module github.com/qiaocco/go-gin-example

go 1.16

require (
	github.com/beego/beego/v2 v2.0.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace (
	github.com/qiaocco/go-gin-example/conf => ./conf
	github.com/qiaocco/go-gin-example/middleware => ./middleware
	github.com/qiaocco/go-gin-example/models => ./models
	github.com/qiaocco/go-gin-example/routers => ./routers
)
