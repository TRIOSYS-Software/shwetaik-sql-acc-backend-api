package models

import (
	"time"
)

type PaymentDetail struct {
	DTLKEY              int        `gorm:"primary_key;column:DTLKEY"`
	DOCKEY              int        `gorm:"column:DOCKEY"`
	SEQ                 uint       `gorm:"column:SEQ"`
	AREA                string     `gorm:"column:AREA;default:----"`
	AGENT               string     `gorm:"column:AGENT;default:----"`
	PROJECT             string     `gorm:"column:PROJECT;default:----"`
	CODE                string     `gorm:"column:CODE"`
	DESCRIPTION         *string    `gorm:"column:DESCRIPTION"`
	GSTDOCDATE          *time.Time `gorm:"column:GST_DOCDATE"`
	GSTDOCNO            *string    `gorm:"column:GST_DOCNO"`
	COMPANYNAME         *string    `gorm:"column:COMPANYNAME"`
	REGISTERNO          *string    `gorm:"column:REGISTERNO"`
	GSTNO               *string    `gorm:"column:GSTNO"`
	PERMITNO            *string    `gorm:"column:PERMITNO"`
	COUNTRY             *string    `gorm:"column:COUNTRY"`
	TAX                 *string    `gorm:"column:TAX"`
	TARIFF              *string    `gorm:"column:TARIFF"`
	IRBM_CLASSIFICATION *string    `gorm:"column:IRBM_CLASSIFICATION"`
	TAXRATE             *string    `gorm:"column:TAXRATE"`
	TAXAMT              float64    `gorm:"column:TAXAMT;default:0"`
	LOCALTAXAMT         float64    `gorm:"column:LOCALTAXAMT;default:0"`
	TAXINCLUSIVE        bool       `gorm:"column:TAXINCLUSIVE"`
	AMOUNT              float64    `gorm:"column:AMOUNT"`
	LOCALAMOUNT         float64    `gorm:"column:LOCALAMOUNT"`
	CURRENCYCODE        string     `gorm:"column:CURRENCYCODE"`
	CURRENCYRATE        float64    `gorm:"column:CURRENCYRATE"`
	CURRENCYAMOUNT      float64    `gorm:"column:CURRENCYAMOUNT"`
	OCR                 *string    `gorm:"column:OCR;default:NULL"`
}

func (PaymentDetail) TableName() string {
	return "GL_CBDTL"
}
