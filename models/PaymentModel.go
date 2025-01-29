package models

import (
	"time"
)

type Payment struct {
	DOCKEY            int             `gorm:"primary_key;column:DOCKEY;autoincrement"`
	DOCNO             string          `gorm:"column:DOCNO"`
	DOCTYPE           string          `gorm:"column:DOCTYPE"`
	DOCDATE           time.Time       `gorm:"column:DOCDATE;autoCreateTime"`
	POSTDATE          time.Time       `gorm:"column:POSTDATE;autoUpdateTime"`
	TAXDATE           time.Time       `gorm:"column:TAXDATE;autoCreateTime"`
	EIV_UTC           *time.Time      `gorm:"column:EIV_UTC;default:NULL"`
	COMPANYNAME       *string         `gorm:"column:COMPANYNAME;default:NULL"`
	DESCRIPTION       *string         `gorm:"column:DESCRIPTION;default:NULL"`
	DESCRIPTION2      *string         `gorm:"column:DESCRIPTION2;default:NULL"`
	ADDRESS1          *string         `gorm:"column:ADDRESS1;default:NULL"`
	ADDRESS2          *string         `gorm:"column:ADDRESS2;default:NULL"`
	ADDRESS3          *string         `gorm:"column:ADDRESS3;default:NULL"`
	ADDRESS4          *string         `gorm:"column:ADDRESS4;default:NULL"`
	POSTCODE          *string         `gorm:"column:POSTCODE;default:NULL"`
	CITY              *string         `gorm:"column:CITY;default:NULL"`
	STATE             *string         `gorm:"column:STATE;default:NULL"`
	COUNTRY           *string         `gorm:"column:COUNTRY;default:NULL"`
	PHONE1            *string         `gorm:"column:PHONE1;default:NULL"`
	PAYMENTMETHOD     string          `gorm:"column:PAYMENTMETHOD"`
	AREA              string          `gorm:"column:AREA;default:----"`
	AGENT             string          `gorm:"column:AGENT;default:----"`
	PROJECT           string          `gorm:"column:PROJECT;default:----"`
	JOURNAL           string          `gorm:"column:JOURNAL"`
	CHEQUENUMBER      *string         `gorm:"column:CHEQUENUMBER;default:NULL"`
	CURRENCYCODE      string          `gorm:"column:CURRENCYCODE;default:----"`
	CURRENCYRATE      float64         `gorm:"column:CURRENCYRATE;default:1"`
	BANKCHARGE        float64         `gorm:"column:BANKCHARGE;default:0"`
	BANKCHARGEACCOUNT *string         `gorm:"column:BANKCHARGEACCOUNT;default:NULL"`
	DOCAMT            float64         `gorm:"column:DOCAMT"`
	LOCALDOCAMT       float64         `gorm:"column:LOCALDOCAMT"`
	FROMDOCTYPE       *string         `gorm:"column:FROMDOCTYPE;default:NULL"`
	BOUNCEDDATE       *time.Time      `gorm:"column:BOUNCEDDATE;default:NULL"`
	GLTRANSID         int             `gorm:"column:GLTRANSID"`
	CANCELLED         bool            `gorm:"column:CANCELLED;default:0"`
	DEPOSITKEY        *uint           `gorm:"column:DEPOSITKEY;default:NULL"`
	FROMDOC           *string         `gorm:"column:FROMDOC;default:NULL"`
	SALESTAXNO        *string         `gorm:"column:SALESTAXNO;default:NULL"`
	SERVICETAXNO      *string         `gorm:"column:SERVICETAXNO;default:NULL"`
	TIN               *string         `gorm:"column:TIN;default:NULL"`
	IDTYPE            uint            `gorm:"column:IDTYPE;default:0"`
	IDNO              *string         `gorm:"column:IDNO;default:NULL"`
	TOURISMNO         *string         `gorm:"column:TOURISMNO;default:NULL"`
	SIC               *string         `gorm:"column:SIC;default:NULL"`
	SUBMISSIONTYPE    uint            `gorm:"column:SUBMISSIONTYPE;default:0"`
	IRBM_STATUS       uint            `gorm:"column:IRBM_STATUS;default:0"`
	IRBM_INTERNALID   *string         `gorm:"column:IRBM_INTERNALID;default:NULL"`
	IRBM_UUID         *string         `gorm:"column:IRBM_UUID;default:NULL"`
	IRBM_LONGID       *string         `gorm:"column:IRBM_LONGID;default:NULL"`
	UPDATECOUNT       uint            `gorm:"column:UPDATECOUNT;default:0"`
	PRINTCOUNT        uint            `gorm:"column:PRINTCOUNT;default:0"`
	ATTACHMENTS       *[]byte         `gorm:"column:ATTACHMENTS;default:NULL"`
	NOTE              *[]byte         `gorm:"column:NOTE;default:NULL"`
	LASTMODIFIED      uint            `gorm:"column:LASTMODIFIED;default:0"`
	DETAILS           []PaymentDetail `gorm:"foreignKey:DOCKEY;references:DOCKEY"`
}

func (Payment) TableName() string {
	return "GL_CB" // Specify table name
}
