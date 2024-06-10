/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 21:08:13
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-07 17:44:22
 * @FilePath     : /hecos-v2-api/config/jwt.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package config

// 创建一个密钥，用于对JWT进行签名
var JwtSecret = "BUVAjvYns51HuBB"
var JwtExpSec int64 = 86400 * 7

var JwtRefreshSecret = "zsHSFdqBTut8JN4"
var JwtRefreshSec int64 = 86400 * 30
