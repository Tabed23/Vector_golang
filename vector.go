package vector

import (
	"fmt"
	"sort"
)

type Vector struct{
	ptr  []int
	 size int
}

func (v * Vector)PushBack(val int) {
	v.ptr = append(v.ptr, val)
	v.size++
}

func (v * Vector)Pop()  {
	v.ptr = v.ptr[:len(v.ptr)-1]
	v.size--
}
func (v * Vector)Size()int  {
	return v.size
}
func (v *Vector)At(idx int)int  {
	if idx > v.size{
		panic("index out of bound")
		return  -1
	}
	return v.ptr[idx]
}
func (v * Vector)Empty()bool  {
	if v.size ==  0{
		return true
	}
	return false
}

func (v * Vector)Front()int {
	if v.size !=  0{
		return v.ptr[0]
	}else {
		panic("vector not yet allocated")
		return -1
	}
}

func (v * Vector)Back()int  {
	if v.size !=  0{
		return v.ptr[v.size-1]
	}else {
		panic("vector not yet allocated")
		return -1
	}
}
func(v *Vector)Print(){
	for i:= 0; i<v.size; i++{
		fmt.Printf("%d ", v.ptr[i])
	}
	fmt.Printf("\n")
}
func (v * Vector)Clear()  {
	v.ptr = nil
	v.size = 0
}
func (v * Vector) Erase(begin *int, end *int)  {
	if *begin <=0 || *begin > v.size && *end <=0 || *end >v.size{
		panic("out of bound error")
		return
	}else{
		v.ptr = append(v.ptr[:*begin], v.ptr[*end+1:]...)
		v.size-= (*begin +*end)/2
	}
}
func (v * Vector)Reverse() {
	for i, j := 0, len(v.ptr)-1; i < j; i, j = i+1, j-1 {
		v.ptr[i], v.ptr[j] = v.ptr[j], v.ptr[i]
	}
}
func(v *Vector) Capacity()int{
	return cap(v.ptr)
}

func (v * Vector)Sort(){
	sort.Ints(v.ptr)
}
