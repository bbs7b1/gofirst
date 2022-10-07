package main

import (
	"testing"
)

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

func TestMatchProfiles(t *testing.T) {
	if false == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
		Profile{{groupId: 22, interestIds: []int{11, 21}}, {groupId: 3, interestIds: []int{3, 45, 32}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
		Profile{{groupId: 1, interestIds: []int{23}}, {groupId: 2, interestIds: []int{3}}},
	) {
		t.Fail()
	}

	if true == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{4, 5}}},
	) {
		t.Fail()
	}
	if true == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}, {groupId: 3, interestIds: []int{3}}},
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
	) {
		t.Fail()
	}

	if true == matchProfiles(
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}},
		Profile{{groupId: 1, interestIds: []int{1, 2}}, {groupId: 2, interestIds: []int{3}}, {groupId: 3, interestIds: []int{3}}},
	) {
		t.Fail()
	}
}
