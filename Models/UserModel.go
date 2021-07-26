package Models



type Student struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	DOB          string `json:"dob"`
	Address      string `json:"address"`
	Subject      string `json:"subject"`
	Marks        int    `json:"marks"`
}
func (b *Student) TableName() string {
	return "student"
}