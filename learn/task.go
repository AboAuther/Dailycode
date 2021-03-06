package main

import (
	"fmt"
	"os"
	"time"
)

type User struct {
	username string
	password string
}
type Logger interface {
	consoleLog()
	fileLog()
}

func newUser(username, password string) User {
	return User{
		username,
		password,
	}
}

func (u User) consoleLog() {
	t := time.Now()
	fmt.Printf("用户创建成功！ 用户名为：%s", u.username)
	fmt.Printf("密码是：%s\n", u.password)
	fmt.Printf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
func (u User) fileLog() {
	t := time.Now()
	file, err := os.OpenFile("./"+u.username+".txt", os.O_RDWR|os.O_CREATE, 0766) // 如果有这个文件就打开没有就新建
	if err != nil {
		fmt.Println(err)
	}
	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) + "密码是：" + u.password + "\n" + fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, _ = file.WriteString(data)
	_ = file.Close()
}
func createUser() {
	var (
		username string
		password string
	)
	fmt.Print("请输入用户名：")
	_, err := fmt.Scan(&username)
	fmt.Print("请输入密码：")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Printf("输入有误！ERROR：%s", err)
	}
	user := newUser(username, password)
	user.consoleLog()
	user.fileLog()
}

func main() {
	createUser()
}
