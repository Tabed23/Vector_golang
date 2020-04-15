package main

import (
	. "Vector_class/vector"
	"fmt"
)

func main(){
	v := Vector{}
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)
	v.PushBack(5)
	v.PushBack(6)
	v.PushBack(7)
	fmt.Printf("vector size := %d\n",v.Size())
	v.Print()
	fmt.Printf("vector at index 2:= %d \n", v.At(2))
	fmt.Println("pop vector")
	v.Pop()
	v.Print()
	fmt.Printf("front := %d\n", v.Front())
	fmt.Printf("back := %d\n", v.Back())
	b,e := 2,4
	v.Erase(&b, &e)
	v.Print()
	fmt.Println("reversing the vector ")
	v.Reverse()
	v.Print()
	fmt.Printf("capicity =%d \n", v.Capacity())
	v.Sort()
	v.Print()
}