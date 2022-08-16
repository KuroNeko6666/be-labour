package models

type FileModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	FileName string `json:"file" gorm:"unique"`
	FilePath string `json:"path" gorm:"unique"`
}
