package model

type News struct {
	Id        uint   `gorm:"primary_key" json:"id" form:"id"`
	Hash      string `json:"hash" gorm:"type:varchar(32);unique_index"`
	Url       string `json:"url" gorm:"type:varchar(100)"`
	Title     string `json:"title" gorm:"type:varchar(150)"`
	Path      string `json:"Path" gorm:"type:varchar(80)"`
	Parts     uint   `json:"port" gorm:"type:int(6)"`
	Readed    string `json:"readed" gorm:"type:int(6)"`
	CreatedAt string `json:"created_at" gorm:"type:varchar(10)"`
}
