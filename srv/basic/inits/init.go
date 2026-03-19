package inits

import "github.com/Lilong-maker/consul"

func init() {
	ViperInit()
	consul.NacosInit()
	MysqlInit()
	RedisInit()

}
