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

	lastPayment := models.Payment{}
	if err := tx.Last(&lastPayment).Error; err != nil && err != gorm.ErrRecordNotFound {
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

	payment.DOCKEY = lastPayment.DOCKEY + 1
	payment.CURRENCYRATE = currencyRate
	payment.JOURNAL = paymentMethod.JOURNAL
	payment.CURRENCYCODE = paymentMethod.CURRENCYCODE
	payment.GLTRANSID = lastPayment.GLTRANSID + 1
	payment.LASTMODIFIED = uint(time.Now().Unix())

	total := 0.0
	for i := range payment.DETAILS {
		lastPaymentDetail := models.PaymentDetail{}
		if err := tx.Last(&lastPaymentDetail).Error; err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return err
		}
		payment.DETAILS[i].DTLKEY = lastPaymentDetail.DTLKEY + uint(i+1)
		total += payment.DETAILS[i].AMOUNT
		payment.DETAILS[i].DOCKEY = payment.DOCKEY
		payment.DETAILS[i].CURRENCYCODE = paymentMethod.CURRENCYCODE
		payment.DETAILS[i].CURRENCYRATE = currencyRate
		if err := tx.Create(&payment.DETAILS[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := p.createGLTrans(tx, payment.DETAILS[i].DESCRIPTION, payment, payment.DETAILS[i].AMOUNT, 0, payment.DETAILS[i].CURRENCYAMOUNT, 0, "S", payment.DETAILS[i].DTLKEY); err != nil {
			tx.Rollback()
			return err
		}
	}
	payment.DOCAMT = total
	payment.LOCALDOCAMT = total

	if err := p.createGLTrans(tx, payment.DESCRIPTION, payment, 0, payment.DOCAMT, 0, payment.LOCALDOCAMT, "M", payment.DOCKEY); err != nil {
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
	description *string,
	payment *models.Payment,
	dr float64,
	cr float64,
	localDr float64,
	localCr float64,
	tableType string,
	fromKey uint,
) error {
	lastGLTrans := models.GLTrans{}
	if err := tx.Last(&lastGLTrans).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	GLTrans := models.GLTrans{
		DOCKEY:       lastGLTrans.DOCKEY + 1,
		GLTRANSID:    int64(payment.GLTRANSID),
		CODE:         payment.PAYMENTMETHOD,
		AREA:         payment.AREA,
		AGENT:        payment.AGENT,
		PROJECT:      payment.PROJECT,
		JOURNAL:      payment.JOURNAL,
		CURRENCYCODE: payment.CURRENCYCODE,
		CURRENCYRATE: payment.CURRENCYRATE,
		DESCRIPTION:  description,
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

	return tx.Create(&GLTrans).Error
}
