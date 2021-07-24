package repository

type user struct {
	Username       string
	Gamertag       string
	MembershipType string
	MembershipID   string

	Characters [3]character

	AccessToken  string
	RefreshToken string
}

type character struct {
	user *user
	ID   string
}
