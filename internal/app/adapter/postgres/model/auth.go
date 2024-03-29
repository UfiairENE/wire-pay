package model

//TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}


//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

//AuthModel ...
type AuthModel struct{}
