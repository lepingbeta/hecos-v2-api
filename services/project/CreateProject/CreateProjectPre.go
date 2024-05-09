/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-09 19:30:21
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-09 19:30:31
 * @FilePath     : /hecos-v2-api/services/project/CreateProject/CreateProjectPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package CreateProject

import (
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func CreateProjectPre(user t.CreateProjectParams) (t.CreateProjectParams, string, string, error) {

	return user, "", "", nil
}
