package data

//Post : struct that represents Post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//Merge : merge current post with another, filling empty fields
func (p *Post) Merge(merge *Post) {
	if p.Body == "" {
		p.Body = merge.Body
	}
	if p.Title == "" {
		p.Title = merge.Title
	}
	if p.UserID == 0 {
		p.UserID = merge.UserID
	}
}
