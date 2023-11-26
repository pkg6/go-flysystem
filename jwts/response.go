package jwts

type TokenResponse struct {
	Token    string `json:"token" xml:"Token"`
	ExpTime  int64  `json:"exp_time" xml:"ExpTime"`
	ExpireIn int64  `json:"expire_in" xml:"ExpireIn"`
}
