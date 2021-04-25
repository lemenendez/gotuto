package main

import ("fmt")

type Fib struct {
	vals map[int] int
}

func NewFib() * Fib {
	f := &Fib{}
	f.vals = make(map[int] int)
	f.vals[0] = 0
	f.vals[1] = 1
	f.vals[2] = 1
	return f
}

func (f *Fib) Get(v int) int {
	if val, ok := f.vals[v]; ok {
		return val
	}
	r:= f.Get(v-1) + f.Get(v-2)
	f.vals[v] = r
	return r
}

func main() {
	f := NewFib()
	fmt.Println(f.Get(0))
	fmt.Println(f.Get(1))
	fmt.Println(f.Get(2))
	fmt.Println(f.Get(3))
	fmt.Println(f.Get(4))
	fmt.Println(f.Get(5))
	fmt.Println(f.Get(6))
	fmt.Println(f.Get(7))
	fmt.Println(f.Get(8))
	fmt.Println(f.Get(9))
	fmt.Println(f.Get(10))
	fmt.Println(f.Get(11))
	fmt.Println(f.Get(12))

}