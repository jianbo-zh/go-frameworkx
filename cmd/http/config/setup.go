package configs

// NormalizeConfig 配置名称规范化
func NormalizeConfig() {
	appConfig()   // 应用相关配置
	logConfig()   // 日志相关配置
	mysqlConfig() // mysql 配置
	redisConfig() // redis 配置
}
