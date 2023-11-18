package klaus

const (
	StateUndefined = iota
	StateGreeted
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Alias    string `json:"alias"`
	Lang     string `json:"lang"`
	Messages []int  `json:"messages"`
	State    int    `json:"state"`
}
