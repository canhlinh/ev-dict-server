package models

type DictEnVi struct {
	ID       string `gorm:"primary_key" json:"-"`
	Word     string `json:"word"`
	Phonetic string `json:"phonetic"`
	Meanings string `json:"meanings"`
}
