package repositories

import (
	"time"

	"shwetaik-sql-acc-backend-api/models"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (p *PaymentRepo) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := p.db.Find(&payments).Error
	return payments, err
}

func (p *PaymentRepo) GetByDOCKEY(docKey uint) (*models.Payment, error) {
	var payment models.Payment
	err := p.db.First(&payment, docKey).Error
	return &payment, err
}

func (p *PaymentRepo) Create(payment *models.Payment) error {
	tx := p.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	firstPayment := models.Payment{}
	if err := tx.First(&firstPayment).Error; err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}

	firstPaymentDetail := models.PaymentDetail{}
	if err := tx.First(&firstPaymentDetail).Error; err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}

	paymentMethodRepo := NewPaymentMethodRepo(p.db)
	paymentMethod, err := paymentMethodRepo.GetByCode(payment.PAYMENTMETHOD)
	if err != nil {
		tx.Rollback()
		return err
	}

	var currencyRate float64
	if err := tx.Raw("SELECT BUYINGRATE FROM CURRENCY WHERE CODE=?", paymentMethod.CURRENCYCODE).Scan(&currencyRate).Error; err != nil {
		tx.Rollback()
		return err
	}

	if firstPayment.DOCKEY > 0 {
		payment.DOCKEY = PaymentDetailID
	} else {
		payment.DOCKEY = firstPaymentDetail.DTLKEY - 1
	}

	payment.CURRENCYRATE = currencyRate
	payment.JOURNAL = paymentMethod.JOURNAL
	payment.CURRENCYCODE = paymentMethod.CURRENCYCODE

	if firstPayment.GLTRANSID > 0 {
		payment.GLTRANSID = GLTransID
	} else {
		payment.GLTRANSID = firstPayment.GLTRANSID - 1
	}

	payment.LASTMODIFIED = uint(time.Now().Unix())

	total := 0.0
	for i := range payment.DETAILS {
		if firstPaymentDetail.DTLKEY > 0 {
			payment.DETAILS[i].DTLKEY = PaymentDetailID - 1
		} else {
			payment.DETAILS[i].DTLKEY = firstPaymentDetail.DTLKEY - int(i+1)
		}
		payment.DETAILS[i].SEQ = uint(i+1) * 1000
		total += payment.DETAILS[i].AMOUNT
		payment.DETAILS[i].DOCKEY = payment.DOCKEY
		payment.DETAILS[i].CURRENCYCODE = paymentMethod.CURRENCYCODE
		payment.DETAILS[i].CURRENCYRATE = currencyRate
		if err := tx.Create(&payment.DETAILS[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := p.createGLTrans(tx, payment.DETAILS[i].CODE, payment.DETAILS[i].DESCRIPTION, payment, payment.DETAILS[i].AMOUNT, 0, payment.DETAILS[i].CURRENCYAMOUNT, 0, "S", payment.DETAILS[i].DTLKEY); err != nil {
			tx.Rollback()
			return err
		}
	}
	payment.DOCAMT = total
	payment.LOCALDOCAMT = total

	if err := p.createGLTrans(tx, payment.PAYMENTMETHOD, payment.DESCRIPTION, payment, 0, payment.DOCAMT, 0, payment.LOCALDOCAMT, "M", payment.DOCKEY); err != nil {
		tx.Rollback()
		return err
	}

	payment.DETAILS = nil
	if err := tx.Save(payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (p *PaymentRepo) createGLTrans(
	tx *gorm.DB,
	code string,
	description *string,
	payment *models.Payment,
	dr float64,
	cr float64,
	localDr float64,
	localCr float64,
	tableType string,
	fromKey int,
) error {
	firstGLTrans := models.GLTrans{}
	if err := tx.First(&firstGLTrans).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	GLTrans := models.GLTrans{
		GLTRANSID:    int64(payment.GLTRANSID),
		CODE:         code,
		AREA:         payment.AREA,
		AGENT:        payment.AGENT,
		PROJECT:      payment.PROJECT,
		JOURNAL:      payment.JOURNAL,
		CURRENCYCODE: payment.CURRENCYCODE,
		CURRENCYRATE: payment.CURRENCYRATE,
		DESCRIPTION:  *description,
		DR:           dr,
		CR:           cr,
		LOCALDR:      localDr,
		LOCALCR:      localCr,
		REF1:         payment.DOCNO,
		FROMDOCTYPE:  payment.DOCTYPE,
		FROMKEY:      fromKey,
		TABLETYPE:    tableType,
		CANCELLED:    payment.CANCELLED,
	}

	if firstGLTrans.DOCKEY > 0 {
		GLTrans.DOCKEY = GLTransID
	} else {
		GLTrans.DOCKEY = firstGLTrans.DOCKEY - 1
	}

	return tx.Create(&GLTrans).Error
}
