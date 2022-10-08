package main

import (
	"fmt"
	"gofirst/db"
	"gofirst/rules"
	"os"
	"runtime"
	"strconv"
)

func printUserCampaigns(userId int) {
	user := db.FindUser(userId)
	if user == nil {
		fmt.Printf("User %d not found\n", userId)
		return
	}
	campaigns := db.GetAllCampaigns()

	ids := rules.GetUserCampaigns(user, campaigns, runtime.NumCPU())
	fmt.Println(ids)
}

func syntax() {
	fmt.Printf("Usage: gofirst -userid id          Print user compaigns\n")
	fmt.Printf("   or: gofirst -port  port         Start http server\n")
	fmt.Printf("   or: gofirst -help               This screen\n")
	fmt.Printf("\n")
	fmt.Printf("Without arguments start http server on 8080\n")
	os.Exit(1)
}

func main() {
	db.Init()
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-help" {
			syntax()
		} else if os.Args[i] == "-port" {
			if len(os.Args) < i+2 {
				syntax()
			}
			port, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				syntax()
			}
			startServer(port)
			return
		} else if os.Args[i] == "-userid" {
			if len(os.Args) < i+2 {
				syntax()
			}
			userId, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				syntax()
			}
			printUserCampaigns(userId)
			return
		} else {
			syntax()
		}

	}
	startServer(8080)
}
