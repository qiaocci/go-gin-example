module github.com/qiaocci/go-gin-example

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	//github.com/qiaocci/go-gin-example/pkg/settings => ~/learn/learn-go/go-gin-example/pkg/settings
	github.com/qiaocci/go-gin-example/conf => ./conf
	github.com/qiaocci/go-gin-example/middleware => ./middleware
	github.com/qiaocci/go-gin-example/models => ./models
	github.com/qiaocci/go-gin-example/routers => ./routers
)
