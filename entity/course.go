package entity

type Course struct {
	Link     string `json:"link,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewCourse() *Course {
	return &Course{
		Password: "apps#12curs-flu2204",
		Link:     "https://dignal.com/academy-flutter/",
	}
}
