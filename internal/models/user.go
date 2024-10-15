package models

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	StudentID     string `json:"student_id"`
	ProfilePicURL string `json:"profile_pic_url"`
	University    string `json:"university"`
	Major         string `json:"major"`
	Year          string `json:"year"`
	Bio           string `json:"bio"`
}
