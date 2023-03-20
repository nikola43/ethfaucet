package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/fatih/color"
	"github.com/nikola43/ethfaucet/app"
)

func main() {
	// system config
	numCpu := runtime.NumCPU()
	usedCpu := numCpu
	runtime.GOMAXPROCS(usedCpu)
	fmt.Println("")
	fmt.Println(color.YellowString("  ----------------- System Info -----------------"))
	fmt.Println(color.CyanString("\t    Number CPU cores available: "), color.GreenString(strconv.Itoa(numCpu)))
	fmt.Println(color.MagentaString("\t    Used of CPU cores: "), color.YellowString(strconv.Itoa(usedCpu)))
	fmt.Println(color.MagentaString(""))

	a := app.App{}
	a.Initialize(":3001")
}
