package main

import (
	"chat/router"
	"chat/utils"
	"fmt"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		fmt.Println("项目启动失败")
	}
}
