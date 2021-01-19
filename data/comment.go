package data

var (
	dbname = "blogbase"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

//Merge : merge current post with another, filling empty fields
func (p *Comment) Merge(merge *Comment) {
	if p.Body == "" {
		p.Body = merge.Body
	}
	if p.Name == "" {
		p.Name = merge.Name
	}
	if p.Email == "" {
		p.Email = merge.Email
	}
	if p.PostID == 0 {
		p.PostID = merge.PostID
	}
}
