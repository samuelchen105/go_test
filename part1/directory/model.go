package directory

type Directory struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" form:"name" binding:"required"`
	Is_hide bool   `json:"is_hide" form:"is_hide"`
}

func (Directory) TableName() string {
	return "directory"
}
