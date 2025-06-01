package repository

import (
	"encoding/json"
	"io"
	"path/filepath"

	"github.com/priscila-albertini-da-silva/item-detail-ml/internal/domain"
	"github.com/priscila-albertini-da-silva/item-detail-ml/pkg"
	"github.com/sirupsen/logrus"
)

type ProductItemRepository interface {
	GetAll() ([]domain.ProductItem, error)
}

type productItemRepository struct {
	dataPath   string
	fileOpener pkg.FileOpener
}

func NewProductRepository(dataPath string) ProductItemRepository {
	return &productItemRepository{
		dataPath:   dataPath,
		fileOpener: pkg.OsFileOpener{},
	}
}

func NewProductRepositoryWithOpener(dataPath string, opener pkg.FileOpener) ProductItemRepository {
	return &productItemRepository{
		dataPath:   dataPath,
		fileOpener: opener,
	}
}

func (r *productItemRepository) GetAll() ([]domain.ProductItem, error) {
	file := filepath.Join(r.dataPath, "products.json")
	f, err := r.fileOpener.Open(file)
	if err != nil {
		logrus.WithError(err).WithField("file", file).Error("Error opening products file")
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
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
