package main

import (
	"fmt"
)

func main() {
	fmt.Printf(`
	=======欢迎使用Arun管理系统=======
	请先登录再进行如下的操作!
	1.add/a(新增用户)
	2.delete/d(删除用户)
	3.modify/m(修改用户)
	4.search/s(查询用户信息)
	`)
	//最近比较忙,先写死了,后续完善
	count := 0
	for count < 3 {
		flag := login()
		if flag {
			break
		} else {
			count = count + 1
			if count == 3 {
				fmt.Printf("输入次数已经达%v,将退出程序!", count)
				return
			} else {
				fmt.Println("密码不正确,请重新输入!")
				continue
			}
		}
	}

	var inputText string
	var err error
	_users := users
	for {
		fmt.Printf("请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):")
		fmt.Scan(&inputText)

		if "add" == inputText || "a" == inputText {
			if _users, err = add(users); err == nil {
				usersPrint(_users)
			} else {
				fmt.Println("添加失败", err)
			}
		} else if "delete" == inputText || "d" == inputText {
			deleteUser(_users)
		} else if "modify" == inputText || "m" == inputText {
			modifyUser(_users)
		} else if "search" == inputText || "s" == inputText {
			findUser(_users)
		}
	}
}

/*
E:\002golang\01NoteAll\20210410day04\02practice\Homework>go run .

        =======欢迎使用Arun管理系统=======
        请先登录再进行如下的操作!
        1.add/a(新增用户)
        2.delete/d(删除用户)
        3.modify/m(修改用户)
        4.search/s(查询用户信息)
        登录密码:
密码不正确,请重新输入!
登录密码:
恭喜您登录成功,登录密码为: 123
请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):a
请输入用户名(1~10非空字符,不能数字开头):arun888
请输入地址(2~20非空字符,不能数字开头):北京市海淀区中关村
请输入联系方式(11位数字):16666886688
|id:1 |name:arun1     |tel:16899886688|addr:北京市                  |
|id:2 |name:josh      |tel:15299886688|addr:America              |
|id:3 |name:arun888   |tel:16666886688|addr:北京市海淀区中关村            |
请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):m
请输入要修改的用户ID:3
id:3,name:arun888,tel:16666886688,addr:北京市海淀区中关村;在数据库中找到用户,您确认要修改用户信息吗?(y/yes/Y/YES)
y
请输入用户名(1~10非空字符,不能数字开头):哈哈哈
请输入联系方式(11位数字):16666886666
请输入地址(2~20非空字符,不能数字开头):地球中国北京
|id:1 |name:arun1     |tel:16899886688|addr:北京市                  |
|id:2 |name:josh      |tel:15299886688|addr:America              |
|id:3 |name:哈哈哈       |tel:16666886666|addr:地球中国北京               |
请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):s
请输入要模糊查询的字符串:哈
id:3,name:哈哈哈,tel:16666886666,addr:地球中国北京;
请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):d
请输入要删除的用户ID:1
id:1,name:arun1,addr:北京市,tel:16899886688;您确认要删除吗?(y/yes/Y/YES)
YES
|id:2 |name:josh      |tel:15299886688|addr:America              |
|id:3 |name:哈哈哈       |tel:16666886666|addr:地球中国北京               |
请选择想要的操作add/a(新增用户),delete/d(删除用户),modify/m(修改用户),search/s(查询用户信息):
*/