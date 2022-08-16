package models

type ScheduleModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Location string `json:"location"`
	Room     string `json:"room"`
	DateTime string `json:"date"`
	Quota    int    `json:"quota"`
}

type ScheduleModelUpdate struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Location string `json:"location"`
	Room     string `json:"room"`
	DateTime string `json:"date"`
	Quota    int    `json:"quota"`
}
