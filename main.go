package main

import onu "github.com/rivaldito/onu/ONU"

func main() {

	onu, err := onu.Constructor()
	if err != nil {
		//TODO: Implentar LOG
		panic(err)
	}
	onu.Login()
	onu.Page()

}
