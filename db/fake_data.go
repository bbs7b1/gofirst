package db

func createUsers() map[int]Profile {
	m := make(map[int]Profile)
	m[1] = []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	return m
}

func createCampaigns() map[int]Profile {
	m := make(map[int]Profile)
	m[1] = []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}
	m[2] = []ProfileRecord{
		{GroupId: 22, InterestIds: []int{11, 21}},
		{GroupId: 3, InterestIds: []int{3, 45, 32}}}
	m[3] = []ProfileRecord{
		{GroupId: 1, InterestIds: []int{23}},
		{GroupId: 3, InterestIds: []int{2}}}
	m[4] = []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{4, 5}}}

	return m
}
