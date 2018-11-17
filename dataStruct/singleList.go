package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	Score float32
	next *Student
}

func trans(p *Student)  {
	for p!=nil {
		fmt.Println(p)
		p=p.next
	}
}

// 尾插
func insertTail(tail *Student)  {
	for i:=0;i<10;i++ {
		stu:=&Student{
			Name:fmt.Sprintf("stu%d",i),
			Age:rand.Intn(100),
			Score:rand.Float32() *100,
		}
		tail.next = stu
		tail = stu
	}
}

func insertHead(head1 **Student)  { // 指针的指针
	for i:=0;i<10;i++ {
		stu := Student{
			Name:fmt.Sprintf("stu%d",i),
			Age:rand.Intn(100),
			Score:rand.Float32() * 100,
		}
		stu.next = *head1
		*head1 = &stu
	}
}

func delNode(p *Student,name string) {
	var prev * Student = p //当前节点地址

	for p!= nil {
		if p.Name == name {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

func addNode(p *Student, newNode *Student, name string) {
	for p!=nil {
		if p.Name == name {
			newNode.next = p.next
			p.next = newNode
		}
		p=p.next
	}
}

func main() {
	//var head Student
	var head *Student = new(Student)
	head.Name = "xiaomi"
	head.Age = 18
	head.Score = 89

	insertHead(&head)
	delNode(head,"xiaomi")
	trans(head)

	var newNode *Student = &Student{
		Name:"stu100",
		Age:19,
		Score:90,
	}
	addNode(head,newNode,"stu6")
	trans(head)

}
