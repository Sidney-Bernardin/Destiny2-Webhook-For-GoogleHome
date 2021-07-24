package repository

type mock struct {
	gamertag       string
	membershipType string
	membershipID   string
}

func NewMock(gamertag, membershipType, membershipID string) Repository {
	return &mock{gamertag, membershipType, membershipID}
}

func (m *mock) GetUser(uname string) (*user, error) {
	return &user{
		Username:       uname,
		Gamertag:       m.gamertag,
		MembershipType: m.membershipType,
		MembershipID:   m.membershipID,
	}, nil
}
