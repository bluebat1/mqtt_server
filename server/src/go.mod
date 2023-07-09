module main

go 1.20

require MqttServer v0.0.0

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/labstack/echo/v4 v4.10.2 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mochi-co/mqtt/v2 v2.2.8 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/rs/zerolog v1.28.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)

replace MqttServer => ./MqttServer
