package repository

import (
	"encoding/json"
	"io"
	"path/filepath"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/pkg"
	"github.com/sirupsen/logrus"
)

type PaymentMethodRepository interface {
	GetAll() ([]domain.PaymentMethod, error)
}

type paymentMethodRepository struct {
	dataPath   string
	fileOpener pkg.FileOpener
}

func NewPaymentMethodRepository(dataPath string) PaymentMethodRepository {
	return &paymentMethodRepository{dataPath: dataPath, fileOpener: pkg.OsFileOpener{}}
}

func NewPaymentMethodRepositoryWithOpener(dataPath string, opener pkg.FileOpener) PaymentMethodRepository {
	return &paymentMethodRepository{
		dataPath:   dataPath,
		fileOpener: opener,
	}
}

func (r *paymentMethodRepository) GetAll() ([]domain.PaymentMethod, error) {
	file := filepath.Join(r.dataPath, "payment_methods.json")
	f, err := r.fileOpener.Open(file)
	if err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error opening payment methods file")
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error reading payment methods file")
		return nil, err
	}

	var paymentMethods []domain.PaymentMethod
	if err := json.Unmarshal(bytes, &paymentMethods); err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error unmarshalling payment methods JSON")
		return nil, err
	}
	return paymentMethods, nil
}
