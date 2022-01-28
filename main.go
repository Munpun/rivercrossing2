package main

import (
	"fmt"
	"github.com/Munpun/rivercrossing2/state"

)

func main() {
	fmt.Println(state.ViewState())
	state.PutInKylling()
	fmt.Println(state.ViewState())

}
