package models

import (
	"time"
)

type GLTrans struct {
	DOCKEY       int        `gorm:"column:DOCKEY;primary_key"`
	GLTRANSID    int64      `gorm:"column:GLTRANSID;primary_key;auto_increment"`
	CODE         string     `gorm:"column:CODE"`
	DOCDATE      time.Time  `gorm:"column:DOCDATE;autoCreateTime"`
	POSTDATE     time.Time  `gorm:"column:POSTDATE;autoUpdateTime"`
	TAXDATE      time.Time  `gorm:"column:TAXDATE;autoCreateTime"`
	AREA         string     `gorm:"column:AREA;default:----"`
	AGENT        string     `gorm:"column:AGENT;default:----"`
	PROJECT      string     `gorm:"column:PROJECT;default:----"`
	TAX          string     `gorm:"column:TAX"`
	JOURNAL      string     `gorm:"column:JOURNAL"`
	CURRENCYCODE string     `gorm:"column:CURRENCYCODE"`
	CURRENCYRATE float64    `gorm:"column:CURRENCYRATE"`
	DESCRIPTION  string     `gorm:"column:DESCRIPTION"`
	DESCRIPTION2 string     `gorm:"column:DESCRIPTION2"`
	DR           float64    `gorm:"column:DR"`
	CR           float64    `gorm:"column:CR"`
	LOCALDR      float64    `gorm:"column:LOCALDR"`
	LOCALCR      float64    `gorm:"column:LOCALCR"`
	REF1         string     `gorm:"column:REF1"`
	REF2         string     `gorm:"column:REF2"`
	FROMDOCTYPE  string     `gorm:"column:FROMDOCTYPE"`
	FROMKEY      int        `gorm:"column:FROMKEY"`
	TABLETYPE    string     `gorm:"column:TABLETYPE"`
	RECONDATE    *time.Time `gorm:"column:RECONDATE"`
	CANCELLED    bool       `gorm:"column:CANCELLED"`
	AUTOPOST     bool       `gorm:"column:AUTOPOST"`
	NONCE        string     `gorm:"column:NONCE;default:0000"`
	DIGEST       string     `gorm:"column:DIGEST"`
}
