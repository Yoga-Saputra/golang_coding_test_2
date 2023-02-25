package members

type MemberFormatter struct {
	MemberID  int    `json:"id_member"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
}

func FormatMember(member Member) MemberFormatter {

	formatter := MemberFormatter{
		MemberID:  member.IDMember,
		Username:  member.Username,
		Gender:    member.Gender,
		Skintype:  member.Skintype,
		Skincolor: member.Skincolor,
	}

	return formatter
}

// mapping  slice members
func FormatMemberSlice(members []Member) []MemberFormatter {
	membersFormatter := []MemberFormatter{}

	for _, member := range members {
		memberFormatter := FormatMember(member)
		membersFormatter = append(membersFormatter, memberFormatter)
	}

	return membersFormatter
}
