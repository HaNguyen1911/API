package models

type ModelForm struct {
	Id            int    `json:"id"`
	FullName      string `json:"fullname"`
	CompanyName   string `json:"companyname"`
	BusinessPhone string `json:"businessphone"`
	Email         string `json:"email"`
	Message       string `json:"message"`
}

func (f *ModelForm) TableName() string {
	return "api"
}
