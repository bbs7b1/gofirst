package main

import (
	"fmt"
	"gofirst/db"
	"gofirst/rules"
	"runtime"
)

func main() {
	users := db.CreateUsers()
	campaigns := db.GetAllCampaigns()

	ids := rules.GetUserCampaigns(users[1], campaigns, runtime.NumCPU())
	fmt.Println(ids)

	//	fmt.Println(users)
	//	fmt.Println(campaigns)

	/*
		s := []Foo{{a: 1, b: "ONE", r: []int{1, 1}}, {a: 2, b: "TWO", r: []int{2, 2}}}

		s = append(s, Foo{3, "THREE", []int{3, 3}})
		s[0].changeFoo()
		for _, el := range s {
			fmt.Println(el)
		}
		fmt.Println(s[0])
	*/
	//createRecord(3);

	/*_, err := as.NewClient("nicetomeetyouava.com", 3000)
	if err != nil {
		fmt.Println("Failed", err)
	} else {
		fmt.Println("Hello, world.")
	}*/
}
