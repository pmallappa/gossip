package main

import (
	"fmt"
	"os"
	//"strings"
)

import (
//"plat"
)

func printBanner() {
	fmt.Printf(`
        $$$$$$\             $$$$$$\  $$$$$$$\  $$$$$$$$\ $$\
       $$  __$$\           $$  __$$\ $$  __$$\ $$  _____|$$ |
       $$ /  \__| $$$$$$\  $$ /  \__|$$ |  $$ |$$ |      $$ |
       $$ |$$$$\ $$  __$$\ \$$$$$$\  $$$$$$$  |$$$$$\    $$ |
       $$ |\_$$ |$$ /  $$ | \____$$\ $$  ____/ $$  __|   $$ |
       $$ |  $$ |$$ |  $$ |$$\   $$ |$$ |      $$ |      $$ |
       \$$$$$$  |\$$$$$$  |\$$$$$$  |$$ |      $$$$$$$$\ $$$$$$$$\
        \______/  \______/  \______/ \__|      \________|\________|

                  Copyright(c) 2009-2013 by Prem Mallappa
                        <prem.mallappa@gmail.com>
`)
}

func main() {

	printBanner()

	parseFlags()

	if err := platform.Init(); err != nil {
		println(os.Args[0])
		os.Exit(-123)
	}
	//  Initialize debugConsole
	//
	//
	//  memoryInit()
	//
	//  load files/bytes specified using -ld
	//
	//  Print on debugConsole Welcome message
	//
	//  Print on Platform Uart Welcome message.
	//
	//  Print on STDOUT Commandline used to invoke
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	platform.Start()
}
