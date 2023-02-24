package members

type InputMember struct {
	Username  string `json:"username" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Skintype  string `json:"skintype" binding:"required"`
	Skincolor string `json:"skincolor" binding:"required"`
}

type GetMemberDetailInput struct {
	IDMember int `uri:"id" binding:"required"`
}
