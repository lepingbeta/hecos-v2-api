package config

// 创建一个密钥，用于对JWT进行签名
var JwtSecret = "BUVAjvYns51HuBB"
var JwtExpSec int64 = 86400 * 7

var JwtRefreshSecret = "zsHSFdqBTut8JN4"
var JwtRefreshSec int64 = 86400 * 30
