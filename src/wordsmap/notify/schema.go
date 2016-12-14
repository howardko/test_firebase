package notify

type PairPost struct {
	Token string `form:"token" json:"token" binding:"required"`
	User_ID string `form:"user_id" json:"user_id" binding:"required"`
}

type PairGet struct {
	Token string
	User_ID string
	Pair_ID string
}