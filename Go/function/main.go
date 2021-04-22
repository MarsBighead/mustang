package main

import "fmt"

type data struct {
	id      int
	content string
}
type instance func(d *data) int

func level1(ins instance) {
	d := &data{id: 32}
	ins(d)
	fmt.Printf("Level1 data=%#v\n", d)
}
func test2(d *data) int {
	fmt.Println("君不见黄河之水天上来,奔流到海不复回")
	d.id += 1
	fmt.Println("num=", d.id)
	return d.id
}
func test3(d *data) int {
	fmt.Println("君不见高堂明镜悲白发朝如青丝暮成雪")
	d.id += 2
	fmt.Println("num=", d.id)
	return d.id
}
func main() {
	//a := "10"
	level1(test2)
	level1(test3)
}
