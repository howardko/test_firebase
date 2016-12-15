package notify

type PairPost struct {
	Token string `form:"token" json:"token" binding:"required"`
	User_ID string `form:"user_id" json:"user_id" binding:"required"`
}

type SendPost struct {
	Pair_ID string `form:"pair_id" json:"pair_id" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
	Data map[string]string `form:"data" json:"data"`
}

type PairGet struct {
	Token string
	User_ID string
	Pair_ID string
}