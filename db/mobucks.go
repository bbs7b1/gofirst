package db

type ProfileRecord struct {
	GroupId     int
	InterestIds []int
}

type Profile []ProfileRecord

var users map[int]Profile

func createUsers() map[int]Profile {
	profile1 := []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}

	m := make(map[int]Profile)
	m[1] = profile1
	return m
}

func GetAllCampaigns() map[int]Profile {
	profile1 := []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{3}}}

	profile2 := []ProfileRecord{
		{GroupId: 22, InterestIds: []int{11, 21}},
		{GroupId: 3, InterestIds: []int{3, 45, 32}}}

	profile3 := []ProfileRecord{
		{GroupId: 1, InterestIds: []int{23}},
		{GroupId: 3, InterestIds: []int{2}}}

	profile4 := []ProfileRecord{
		{GroupId: 1, InterestIds: []int{1, 2}},
		{GroupId: 2, InterestIds: []int{4, 5}}}

	m := make(map[int]Profile)
	m[1] = profile1
	m[2] = profile2
	m[3] = profile3
	m[4] = profile4
	return m
}

func FindUser(userId int) Profile {
	return users[userId]
}

func Init() {
	users = createUsers()
}
