package mentors

// Mentor holds mentor informantion
type Mentor struct {
	FullName string `json:"full_name"`
	UserName string `json:"username"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Tweet    string `json:"tweet"`
	TweetID  string `json:"tweet_id"`
}
