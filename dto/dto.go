package dto

type DTOUser struct {
	Username     string
	PasswordHash string
	UserRole     string
	Email        string
}

type DTOEmail struct {
	Email     string
	EmailCode int32
}

type DTOTokens struct {
	Username string
	Token    string
}

type DTOCreateUserReq struct {
	DTOUser
	DTOEmail
}

type DTOLoginUserResp struct {
}

type DTOUserMetadata struct {
	Fingerprint  string
	RefreshToken string
	Username     string
}
