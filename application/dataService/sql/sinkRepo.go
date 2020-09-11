package sql

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sinkRepo struct {
	db *gorm.DB
}

func NewSinkRepo() *sinkRepo {
	return &sinkRepo{
		db: dbConn,
	}
}

func (sir *sinkRepo) FindsWithTopic() (sl []model.Sink, err error) {
	return sl, sir.db.Preload("Topic").Find(&sl).Error
}

func (sir *sinkRepo) FindByIDWithNodesSensorsValuesTopic(id int) (*model.Sink, error) {
	s := &model.Sink{}
	return s, sir.db.Where("id=?", id).Preload("Nodes.Sensors.SensorValues", orderByASC).Preload("Nodes.Sensors").Preload("Nodes").Preload("Topic").Find(s).Error
}

func (sir *sinkRepo) Create(s *model.Sink) error {
	return sir.db.Omit(clause.Associations).Create(s).Error
}

func (sir *sinkRepo) Delete(s *model.Sink) error {
	return sir.db.Delete(s).Error
}
