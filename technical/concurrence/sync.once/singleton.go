package sync_once

import "fmt"

type Singleton interface {
	Say()
}

type defaultInpl struct {
}

func (*defaultInpl) Say() {
	fmt.Println("hello world")
}
