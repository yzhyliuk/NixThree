package comment

type Comment struct {
	PostID	int		`json:"postId"`
	ID 		int		`json:"id"`
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
	body 	string	`json:"body"`
}
