package main

import (
	"fmt"
)

type ProfileRecord struct {
	groupId     int
	interestIds []int
}

type Profile []ProfileRecord

func createTestUsers() map[int]Profile {
	profile1 := []ProfileRecord{
		{groupId: 1, interestIds: []int{1, 2}},
		{groupId: 2, interestIds: []int{3}}}

	m := make(map[int]Profile)
	m[1] = profile1
	return m
}

func createTestCampaigns() map[int]Profile {
	profile1 := []ProfileRecord{
		{groupId: 1, interestIds: []int{1, 2}},
		{groupId: 2, interestIds: []int{3}}}

	profile2 := []ProfileRecord{
		{groupId: 22, interestIds: []int{11, 21}},
		{groupId: 3, interestIds: []int{3, 45, 32}}}

	profile3 := []ProfileRecord{
		{groupId: 1, interestIds: []int{23}},
		{groupId: 3, interestIds: []int{2}}}

	profile4 := []ProfileRecord{
		{groupId: 1, interestIds: []int{1, 2}},
		{groupId: 2, interestIds: []int{4, 5}}}

	m := make(map[int]Profile)
	m[1] = profile1
	m[2] = profile2
	m[3] = profile3
	m[4] = profile4
	return m
}

func intersects(a []int, b []int) bool {
	imap := make(map[int]bool)

	for i := 0; i < len(a); i++ {
		imap[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if imap[b[i]] == true {
			return true
		}
	}
	return false
}

func matchProfiles(user Profile, campaign Profile) bool {
	gmap := make(map[int]int)
	imap := make(map[int][]int)
	for _, uRec := range user {
		gmap[uRec.groupId] = 1
		imap[uRec.groupId] = uRec.interestIds
	}
	for _, cRec := range campaign {
		gmap[cRec.groupId] += 1
		if !intersects(cRec.interestIds, imap[cRec.groupId]) {
			return false
		}
	}
	for _, i := range gmap {
		if i != 2 {
			return false
		}
	}
	return true
}

func getUsersCampaigns(user Profile, campaigns map[int]Profile) []int {
	ids := []int{}
	for key, campaign := range campaigns {
		if matchProfiles(user, campaign) {
			ids = append(ids, key)
		}
	}
	return ids
}

func main() {
	users := createTestUsers()
	campaigns := createTestCampaigns()
	ids := []int{}

	for _, user := range users {
		for key, campaign := range campaigns {
			if matchProfiles(user, campaign) {
				ids = append(ids, key)
			}
		}
	}
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
