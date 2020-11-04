package domain

type StudentCredentials struct {
	Sid 	int 	`json:"sid" form:"sid"`
	Did 	int 	`json:"did" form:"did"` 
	Marks	int		`json:"marks" form:"marks"`
}