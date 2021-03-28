package main

import "fmt"

var slice []func()

func main() {
	/*	sl_from := []int{1, 2, 3}
		sl_to := []int{4,5,6,7}
		fmt.Println(InsertStringSlice(sl_to,sl_from,2))*/
	//
	//n := copy(sl_from, sl_to)
	//fmt.Println(sl_from)
	//fmt.Printf("Copied %d elements\n", n) // n == 3

	//写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
	/*	sl3 := []int{1, 2, 3}
		sl3 = append(sl3, 4, 5, 6)
		fmt.Println(sl3)*/

	//写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除
	/*	string:=[]string{"beijing","tianjin","shanghai"}
		fmt.Println(RemoveStringSlice(string))*/

	//使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
	/*	l:=list.New()
		for i:=101;i<103;i++{
			l.PushBack(i)
		}
		for p := l.Front(); p != nil; p = p.Next() {
			fmt.Println("Number", p.Value)
		}*/
	//通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
	/*	var i  int8 =32
		fmt.Printf("i 的类型 %T 占中的字节数是 %d", i, unsafe.Sizeof(i))*/
	//闭包函数满足条件   闭包=函数+引用环境
	/*	a,b :=calc(100)
		r:=a(200)
		v:=b(100)
		fmt.Println(r)
		fmt.Println(v)*/

	m := make(map[string]*student)
	stus := []student{

		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		stu := stu
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(&v.name)
		fmt.Println(k, "=>", v.name)
	}

	//JSON反序列化
	/*c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)*/

	/*	fmt.Println(pack.Add(2,3))
		var peo People = Student{}
		think := "bitch"
		fmt.Println(peo.Speak(think))*/

}
func appendMul(slice []int, factor int) []int {

	newslice := make([]int, len(slice)*factor)
	copy(newslice, slice)
	slice = newslice
	return slice
}
func InsertStringSlice(des, src []int, x int) []int {
	/*	newslice:=make([]int,len(des)+len(src))
		copy(newslice,des[0:x])
		copy(newslice[x:],src[:])
		copy(newslice[x+len(src):],des[x:])
		return newslice*/
	des = append(des[:x], append(src, des[:x]...)...)
	return des
}
func RemoveStringSlice(slice []string) []string {
	return append(slice[:0], slice[len(slice):]...)
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

type student struct {
	name string
	age  int
}

////Student 学生
//type Student struct {
//	ID     int `JSON:"id"`
//	Gender string
//	Name   string
//}
//
////Class 班级
//type Class struct {
//	Title    string
//	Students []*Student
//}
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
