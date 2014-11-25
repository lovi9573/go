package main

import (
    "runtime"
)

func pmergesort(ar [] directory){
	//Create parallel mergesort goroutines
	runtime.GOMAXPROCS(4)
	order := make(chan bool)
	size := len(ar)
	go __pmergesort(order, ar[0:size/4])
	go __pmergesort(order, ar[size/4:size/2])
	go __pmergesort(order, ar[size/2:3*size/4])
	go __pmergesort(order, ar[3*size/4:])
	<- order  //Wait here until the routines finish
	<- order
	<- order
	<- order
	merge(ar[:size/4], ar[size/4:size/2], ar[:size/2])
	merge(ar[size/2:3*size/4], ar[3*size/4:], ar[size/2:])
	merge(ar[0:size/2], ar[size/2:], ar)

}

func __pmergesort(order chan bool, ar [] directory) {
	mergesort(ar)
	order <- true
}

func mergesort(ar [] directory) {
	l := len(ar)
	if l <= 1 {
		return
	}
	left := ar[ :l/2]
	right := ar[l/2: ]
	mergesort(left) 
	mergesort(right)
	merge(left,right, ar)
}

func merge(left [] directory, right [] directory, ar [] directory)  {
	var tmp = make([] directory, len(left)+len(right))
	a := 0
	b := 0
	t := 0
	for  a < len(left)  || b < len(right) {
		if b == len(right) || a < len(left) && left[a].size < right[b].size {
			tmp[t] = left[a]
			a +=1
		}else {
			tmp[t] = right[b]
			b+=1
		}
		t+=1	
	}
	for i := range ar {
		ar[i] = tmp[i]
	}
}
