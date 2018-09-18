package main

import ("fmt")

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	
}