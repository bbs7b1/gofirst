package rules

import (
	"gofirst/db"
	"sync"
)

// check if two arrays have at least one common element
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

// compare profiles according the rule
func matchProfiles(user db.Profile, campaign db.Profile) bool {
	gmap := make(map[int]int)
	imap := make(map[int][]int)
	for _, uRec := range user {
		gmap[uRec.GroupId] = 1
		imap[uRec.GroupId] = uRec.InterestIds
	}
	for _, cRec := range campaign {
		gmap[cRec.GroupId] += 1
		if !intersects(cRec.InterestIds, imap[cRec.GroupId]) {
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

// share campaigns according maxprocs (CPUs)
func shareCampaignsPerProc(campaigns map[int]db.Profile, maxprocs int) [](map[int]db.Profile) {
	smap := [](map[int]db.Profile){}
	for i := 0; i < maxprocs; i++ {
		smap = append(smap, make(map[int]db.Profile))
	}
	i := 0
	for key, el := range campaigns {
		smap[i][key] = el
		i++
		if i >= len(smap) {
			i = 0
		}
	}
	return smap
}

func asyncMatchProfiles(user db.Profile, m map[int]db.Profile, ids *[]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for key, campaign := range m {
		if matchProfiles(user, campaign) {
			mu.Lock()
			*ids = append(*ids, key)
			mu.Unlock()
		}
	}
}

// return an array of campaigns ids that match user profile
func GetUserCampaigns(user db.Profile, campaigns map[int]db.Profile, maxprocs int) []int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ids := []int{}
	smap := shareCampaignsPerProc(campaigns, maxprocs)
	for _, m := range smap {
		wg.Add(1)
		go asyncMatchProfiles(user, m, &ids, &mu, &wg)
	}
	wg.Wait()
	return ids
}
