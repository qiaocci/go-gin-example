module github.com/qiaocco/go-gin-example

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/beego/beego/v2 v2.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/getsentry/sentry-go v0.11.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	golang.org/x/tools v0.1.2 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace (
	github.com/qiaocco/go-gin-example/conf => ./conf
	github.com/qiaocco/go-gin-example/docs => ./docs
	github.com/qiaocco/go-gin-example/middleware => ./middleware
	github.com/qiaocco/go-gin-example/models => ./models
	github.com/qiaocco/go-gin-example/pkg => ./pkg
	github.com/qiaocco/go-gin-example/routers => ./routers
)
