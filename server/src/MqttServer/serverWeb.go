package MqttServer

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func RunWebServer() {

	echoService := echo.New()
	echoService.Static("/", "./view/")
	// echoService.Static("/assets", "./view/assets/")
	echoService.GET("/", func(c echo.Context) error {
		fb, err := ioutil.ReadFile("./view/index.html")
		if err != nil {
			return err
		}
		c.HTML(200, string(fb))
		return nil
	})

	// 登录
	echoService.POST("/server/login", func(c echo.Context) error {
		_user := c.FormValue("user")
		rows, err := SqlUtil.Query("select '1' from user where user=? and passwd=? limit 1;", _user, c.FormValue("passwd"))
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "sql err",
			})
			return nil
		}
		defer rows.Close()
		isLogin := ""
		for rows.Next() {
			rows.Scan(&isLogin)
		}

		if isLogin == "" {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "login err",
			})
			return nil
		}
		md5Token := randStr(64)
		jwtToken, err := JwtNew(_user, md5Token)
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "jwt err",
			})
			return nil
		}

		rows2, err := SqlUtil.Query("update user set token=? where user=? limit 1;", md5Token, _user)
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "sql2 err",
			})
			return nil
		}
		rows2.Close()

		c.JSON(200, map[string]string{
			"status": "ok",
			"token":  jwtToken,
		})

		return nil
	})

	// 登录检测
	echoService.POST("/server/checkLogin", func(c echo.Context) error {

		jwtToken, err := JwtParse(c.FormValue("token"))
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "token err",
			})
			return nil
		}

		rows, err := SqlUtil.Query("select '1' from user where token=? limit 1;", jwtToken.Token)
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "sql err",
			})
			return nil
		}
		defer rows.Close()

		isOk := ""
		for rows.Next() {
			rows.Scan(&isOk)
		}

		if isOk != "1" {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "no login err",
			})
			return nil
		}

		c.JSON(200, map[string]string{
			"status": "ok",
		})

		return nil
	})

	// 获取 账户 剩余资源
	echoService.POST("/server/getUserAssets", func(c echo.Context) error {
		jwtToken, err := JwtParse(c.FormValue("token"))
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "token err",
			})
			return nil
		}

		rows, err := SqlUtil.Query("call getAssets(?);", jwtToken.User)
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "token err",
			})
			return nil
		}
		defer rows.Close()

		nodeAssetsMax := ""
		topicAssetsMax := ""
		nodeAssetsNum := ""
		topicAssetsNum := ""

		for rows.Next() {
			rows.Scan(&nodeAssetsMax, &topicAssetsMax, &nodeAssetsNum, &topicAssetsNum)
		}

		c.JSON(200, map[string]string{
			"status":         "ok",
			"nodeAssetsMax":  nodeAssetsMax,
			"topicAssetsMax": topicAssetsMax,
			"nodeAssetsNum":  nodeAssetsNum,
			"topicAssetsNum": topicAssetsNum,
		})
		return nil
	})

	// 获取 mqtt 节点账号列表
	echoService.POST("/server/getMqttUserList", func(c echo.Context) error {
		jwtToken, err := JwtParse(c.FormValue("token"))
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "token err",
			})
			return nil
		}
		rows, err := SqlUtil.Query("select user,passwd,date from mqtt_user where master=? limit 50;", jwtToken.User)
		if err != nil {
			c.JSON(200, map[string]string{
				"status": "err",
				"msg":    "sql err",
			})
			return nil
		}
		defer rows.Close()

		lists := []map[string]string{}
		_u := ""
		_d := ""
		_p := ""
		for rows.Next() {
			rows.Scan(&_u, &_d, &_p)
			lists = append(lists, map[string]string{
				"user":   _u,
				"date":   _d,
				"passwd": _p,
			})
		}

		c.JSON(200, map[string]interface{}{
			"status": "ok",
			"users":  lists,
		})

		return nil
	})

	echoService.Logger.Fatal(echoService.Start(":" + ConfigGet("web.port")))
}
