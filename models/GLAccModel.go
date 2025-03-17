package models

type GLAcc struct {
	DOCKEY         int
	PARENT         int
	CODE           string
	DESCRIPTION    string
	DESCRIPTION2   string
	ACCTYPE        string
	SPECIALACCTYPE string
	TAX            string
	CASHFLOWTYPE   int
	SIC            string
}

func (GLAcc) TableName() string {
	return "GL_ACC"
}
