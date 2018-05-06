package mentors

// Mentor holds mentor informantion
type Mentor struct {
	UserID       string `json:"user_id"`
	FullName     string `json:"full_name"`
	UserName     string `json:"username"`
	Bio          string `json:"bio"`
	Website      string `json:"website"`
	Tweet        string `json:"tweet"`
	TweetID      string `json:"tweet_id"`
	ProfileImage string `json:"profile_image"`
}
