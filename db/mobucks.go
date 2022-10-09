package db

type ProfileRecord struct {
	GroupId     int
	InterestIds []int
}

type Profile []ProfileRecord

var users map[int]Profile
var campaigns map[int]Profile

func FindUser(userId int) Profile {
	return users[userId]
}

func GetAllCampaigns() map[int]Profile {
	return campaigns
}

func Init() {
	users = createUsers()
	campaigns = createCampaigns()
}
