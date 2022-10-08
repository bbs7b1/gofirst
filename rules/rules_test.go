package rules

import (
	"gofirst/db"
	"reflect"
	"sort"
	"testing"
)

/////////////////////////////////////////////////////////////////

func TestIntersects(t *testing.T) {
	t.Log("Testing Intersects")
	if !intersects([]int{1, 2, 3}, []int{3, 4, 5}) {
		t.FailNow()
	}
	if !intersects([]int{1, 2, 3}, []int{3, 2, 5}) {
		t.FailNow()
	}
	if intersects([]int{1}, []int{3, 4, 5}) {
		t.FailNow()
	}
	if intersects([]int{1, 2, 3, 4}, []int{5, 6, 7}) {
		t.FailNow()
	}
}

//////////////////////////////////////////////////////////////

func TestMatchProfiles(t *testing.T) {
	if false == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
		db.Profile{{GroupId: 22, InterestIds: []int{11, 21}}, {GroupId: 3, InterestIds: []int{3, 45, 32}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
		db.Profile{{GroupId: 1, InterestIds: []int{23}}, {GroupId: 2, InterestIds: []int{3}}},
	) {
		t.Fail()
	}

	if true == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{4, 5}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}, {GroupId: 3, InterestIds: []int{3}}},
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
	) {
		t.Fail()
	}

	if true == matchProfiles(
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}},
		db.Profile{{GroupId: 1, InterestIds: []int{1, 2}}, {GroupId: 2, InterestIds: []int{3}}, {GroupId: 3, InterestIds: []int{3}}},
	) {
		t.Fail()
	}
}

///////////////////////////////////////////////////////////////////////////////

func createNCampaigns(n int) map[int]db.Profile {
	ret := make(map[int]db.Profile)
	for i := 0; i < n; i++ {
		ret[i] = db.Profile{{GroupId: i, InterestIds: []int{i}}}
	}
	return ret
}

func TestCreateSharedMap(t *testing.T) {
	smap := createSharedMap(createNCampaigns(10), 1)
	if len(smap) != 1 {
		t.Fail()
		t.Log("len(smap) != 1")
	}
	for i := 0; i < len(smap); i++ {
		if len(smap[i]) != 10 {
			t.Fail()
			t.Log("len != 10")
		}
	}

	smap = createSharedMap(createNCampaigns(4), 8)
	for i := 0; i < len(smap); i++ {
		if len(smap[i]) > 1 {
			t.Fail()
			t.Log("len > 1")
		}
	}

	smap = createSharedMap(createNCampaigns(100), 4)
	for i := 0; i < len(smap); i++ {
		if len(smap[i]) != 25 {
			t.Fail()
			t.Log("len != 25")
		}
	}

	smap = createSharedMap(createNCampaigns(25), 5)
	for i := 0; i < len(smap); i++ {
		if len(smap[i]) != 5 {
			t.Fail()
			t.Log("len != 5")
		}
	}

	smap = createSharedMap(createNCampaigns(25), 8)
	if len(smap[0]) != 4 {
		t.Fail()
		t.Log("len[0] != 4")
	}
	for i := 1; i < len(smap); i++ {
		if len(smap[i]) != 3 {
			t.Fail()
			t.Log("len != 3")
		}
	}
}

////////////////////////////////////////////////////////////////////

func createTestCampaigns4() map[int]db.Profile {
	m := make(map[int]db.Profile)
	m[1] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[2] = []db.ProfileRecord{
		{GroupId: 22, InterestIds: []int{11, 21}},
		{GroupId: 3, InterestIds: []int{3, 45, 32}}}
	m[3] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{23}},
		{GroupId: 3, InterestIds: []int{2}}}
	m[4] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{4, 5}}}
	return m
}

func createTestCampaigns5() map[int]db.Profile {
	m := make(map[int]db.Profile)
	m[1] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[2] = []db.ProfileRecord{
		{GroupId: 22, InterestIds: []int{11, 21}},
		{GroupId: 3, InterestIds: []int{3, 45, 32}}}
	m[3] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{23}},
		{GroupId: 3, InterestIds: []int{2}}}
	m[4] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{4, 5}}}
	m[5] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 6}},
		{GroupId: 2, InterestIds: []int{3, 5}}}
	return m
}

func createTestCampaigns6() map[int]db.Profile {
	m := make(map[int]db.Profile)
	m[1] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[2] = []db.ProfileRecord{
		{GroupId: 22, InterestIds: []int{11, 21}},
		{GroupId: 3, InterestIds: []int{3, 45, 32}}}
	m[3] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{23}},
		{GroupId: 3, InterestIds: []int{2}}}
	m[4] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{4, 5}}}
	m[5] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 6}},
		{GroupId: 2, InterestIds: []int{3, 5}}}
	m[6] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}},
		{GroupId: 3, InterestIds: []int{3}}}
	return m
}

func createTestCampaignsAll() map[int]db.Profile {
	m := make(map[int]db.Profile)
	m[1] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[2] = []db.ProfileRecord{
		{GroupId: 2, InterestIds: []int{3, 21}},
		{GroupId: 1, InterestIds: []int{2, 45, 32}}}
	m[3] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[4] = []db.ProfileRecord{
		{GroupId: 2, InterestIds: []int{3}},
		{GroupId: 1, InterestIds: []int{2, 1, 4}}}
	m[5] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 6}},
		{GroupId: 2, InterestIds: []int{3, 5}}}
	m[6] = []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 6}},
		{GroupId: 2, InterestIds: []int{3, 5}},
		{GroupId: 3, InterestIds: []int{3, 5}}, // this will not returned
	}
	return m
}

func TestGetUserCampaigns(t *testing.T) {
	user := []db.ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	ids := GetUserCampaigns(user, createTestCampaigns4(), 8)
	if !reflect.DeepEqual(ids, []int{1}) {
		t.Fail()
		t.Log("not [1], got ", ids)
	}
	ids = GetUserCampaigns(user, createTestCampaigns5(), 8)
	sort.Ints(ids) // order is no guaranty because of async stuff so sort it to compare
	if !reflect.DeepEqual(ids, []int{1, 5}) {
		t.Fail()
		t.Log("not [1, 5], got ", ids)
	}
	ids = GetUserCampaigns(user, createTestCampaigns6(), 8)
	sort.Ints(ids)
	if !reflect.DeepEqual(ids, []int{1, 5}) {
		t.Fail()
		t.Log("not [1, 5], got", ids)
	}

	ids = GetUserCampaigns(user, createTestCampaignsAll(), 8)
	sort.Ints(ids)
	if !reflect.DeepEqual(ids, []int{1, 2, 3, 4, 5}) {
		t.Fail()
		t.Log("not [1, 5], got", ids)
	}
}
