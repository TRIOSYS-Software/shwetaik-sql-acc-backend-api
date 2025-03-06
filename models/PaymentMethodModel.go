package models

type PaymentMethod struct {
	CODE              string  `gorm:"column:CODE;type:VARCHAR;size:10"`
	JOURNAL           string  `gorm:"column:JOURNAL;type:VARCHAR;size:10"`
	BANKCHARGEACCOUNT string  `gorm:"column:BANKCHARGEACCOUNT;type:VARCHAR;size:10"`
	CURRENCYCODE      string  `gorm:"column:CURRENCYCODE;type:VARCHAR;size:6"`
	OVERDRAFTLIMIT    float64 `gorm:"column:OVERDRAFTLIMIT;type:DECIMAL;precision:18;scale:2"`
	ORDOCNUMBER       int64   `gorm:"column:ORDOCNUMBER;type:BIGINT"`
	PVDOCNUMBER       int64   `gorm:"column:PVDOCNUMBER;type:BIGINT"`
	BANKID            string  `gorm:"column:BANKID;type:VARCHAR;size:10"`
	GIRO              []byte  `gorm:"column:GIRO;type:BLOB SUB_TYPE BINARY"`
	DATA              []byte  `gorm:"column:DATA;type:BLOB SUB_TYPE BINARY"`
	ATTACHMENTS       []byte  `gorm:"column:ATTACHMENTS;type:BLOB SUB_TYPE BINARY"`
	DESCRIPTION       string
}

func (PaymentMethod) TableName() string {
	return "PMMETHOD"
}
