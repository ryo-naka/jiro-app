package response

type CreateUser struct {
	ResultCode string `json:"resultCode"`
}

type ShowUser struct {
	ID               int        `json:"id"`
	Nickname         string     `json:"nickname"`
	Email            string     `json:"email"`
	Followingcount   int        `json:"followingCount"`
	Followercount    int        `json:"followerCount"`
	FavoriteContents []*Content `json:"favoriteContents"`
}

type Users struct {
	Users []*User
}

type User struct {
	ID       int
	Nickname string
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
