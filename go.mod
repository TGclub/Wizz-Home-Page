module Wizz-homepage-go

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/appleboy/gin-jwt/v2 v2.6.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/spf13/viper v1.6.2
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
)

replace (
    cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go latest
    golang.org/x/crypto => github.com/golang/crypto latest
    golang.org/x/net => github.com/golang/net latest
    golang.org/x/sync => github.com/golang/sync latest
    golang.org/x/sys => github.com/golang/sys latest
    golang.org/x/text => github.com/golang/text latest
    google.golang.org/appengine => github.com/golang/appengine latest
)