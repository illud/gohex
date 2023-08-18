package main

import (
	"fmt"

	cmd "github.com/illud/gohex/src/cli"
)

func main() {

	gohex := `
	_____       _               
	|  __ \     | |              
	| |  \/ ___ | |__   _____  __
	| | __ / _ \| '_ \ / _ \ \/ /
	| |_\ \ (_) | | | |  __/>  < 
	 \____/\___/|_| |_|\___/_/\_\
				 Created by Illud
`

	fmt.Println(gohex)
	cmd.Command()
}
