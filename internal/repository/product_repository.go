package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/sirupsen/logrus"
)

type ProductRepository interface {
	GetAll() ([]domain.ProductItem, error)
}

type productRepository struct {
	dataPath string
}

func NewProductRepository(dataPath string) ProductRepository {
	return &productRepository{dataPath: dataPath}
}

func (r *productRepository) GetAll() ([]domain.ProductItem, error) {
	file := filepath.Join(r.dataPath, "products.json")
	f, err := os.Open(file)
	if err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error opening products file")
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error reading products file")
		return nil, err
	}

	var products []domain.ProductItem
	if err := json.Unmarshal(bytes, &products); err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error unmarshalling products JSON")
		return nil, err
	}
	return products, nil
}
