package services

import (
	"homify-go-grpc/internal/property-service/models"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/repositories"
	"homify-go-grpc/internal/property-service/types"
	kafka_configs "homify-go-grpc/internal/shared/kafka-configs"

	"gorm.io/gorm"
)

type IPropertyService interface {
	AddNewProperty(
		assetIds types.PropertyAssetIds,
		newProperty models.Property,
		newDestination models.Destination,
	) (bool, error)
}

type PropertyService struct {
	repo     repositories.IPropertyRepository
	producer producers.IPropertyProducer
}

func NewPropertyService(db *gorm.DB, p producers.IPropertyProducer) IPropertyService {
	return &PropertyService{
		repo:     repositories.NewPropertyRepository(db),
		producer: p,
	}
}

func (s *PropertyService) AddNewProperty(
	assetIds types.PropertyAssetIds,
	newProperty models.Property,
	newDestination models.Destination,
) (bool, error) {

	// Publish to kafka
	context := kafka_configs.GetContext()
	s.producer.ProduceMessages(context.SearchTopic, "Hi there !")

	return true, nil
}
