package models

type Project struct {
	CODE         string   `gorm:"column:CODE;primary_key"`
	DESCRIPTION  *string  `gorm:"column:DESCRIPTION"`
	DESCRIPTION2 *string  `gorm:"column:DESCRIPTION2"`
	PROJECTVALUE *float64 `gorm:"column:PROJECTVALUE;default:0"`
	PROJECTCOST  *float64 `gorm:"column:PROJECTCOST;default:0"`
	ATTACHMENTS  *[]byte  `gorm:"column:ATTACHMENTS"`
	ISACTIVE     *bool    `gorm:"column:ISACTIVE"`
}

func (p *Project) TableName() string {
	return "PROJECT"
}
