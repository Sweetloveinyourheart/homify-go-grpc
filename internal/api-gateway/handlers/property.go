package handlers

import (
	"context"
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/api-gateway/dtos"
	"homify-go-grpc/internal/api-gateway/helpers"
	"homify-go-grpc/internal/api-gateway/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PropertyHandler struct {
	grpcClient proto.PropertyClient
	validator  *validator.Validate
}

func NewPropertyHandler(c proto.PropertyClient, validate *validator.Validate) *PropertyHandler {
	return &PropertyHandler{
		grpcClient: c,
		validator:  validate,
	}
}

// @Summary Add a new property
// @Tags Property
// @Description Add a new property to the property listings.
// @ID add-new-property
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param input body dtos.NewPropertyDTO true "New Property object to be added"
// @Success 200 {object} interface{} "Successfully added the new property"
// @Router /property [post]
func (h *PropertyHandler) AddNewProperty(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(403, gin.H{"error": "Forbidden resource"})
		return
	}

	authenticatedUser := user.(middlewares.AuthenticatedUser)

	newPropertyData := dtos.NewPropertyDTO{}
	if bindError := ctx.ShouldBindJSON(&newPropertyData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(newPropertyData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	newProperty := &proto.NewPropertyRequest{
		HostId:      uint32(authenticatedUser.Id),
		Title:       newPropertyData.Title,
		Description: newPropertyData.Description,
		Price:       newPropertyData.Price,
		AmenityId:   uint32(newPropertyData.AmenityId),
		CategoryId:  uint32(newPropertyData.CategoryId),
		Destination: &proto.NewDestinationRequest{
			Country:   newPropertyData.Destination.Country,
			City:      newPropertyData.Destination.City,
			Latitude:  newPropertyData.Destination.Latitude,
			Longitude: newPropertyData.Destination.Longitude,
		},
	}

	propertyCtx := context.Background()

	h.grpcClient.AddProperty(propertyCtx, newProperty)

	ctx.JSON(200, gin.H{
		"Success": true,
	})
}

// @Summary Edit a property
// @Tags Property
// @Description Edit a property to the property listings.
// @ID edit-property
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param input body dtos.EditPropertyDTO true "Property object to be updated"
// @Success 200 {object} interface{} "Successfully updated the new property"
// @Router /property [put]
func (h *PropertyHandler) EditProperty(ctx *gin.Context) {
	propertyData := dtos.EditPropertyDTO{}
	if bindError := ctx.ShouldBindJSON(&propertyData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(propertyData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	property := &proto.EditPropertyRequest{
		Title:       propertyData.Title,
		Description: propertyData.Description,
		Price:       propertyData.Price,
		AmenityId:   uint32(propertyData.AmenityId),
		CategoryId:  uint32(propertyData.CategoryId),
		Destination: &proto.EditDestinationRequest{
			Country:   propertyData.Destination.Country,
			City:      propertyData.Destination.City,
			Latitude:  propertyData.Destination.Latitude,
			Longitude: propertyData.Destination.Longitude,
		},
	}

	propertyCtx := context.Background()

	h.grpcClient.EditProperty(propertyCtx, property)

	ctx.JSON(200, gin.H{
		"Success": true,
	})
}

// @Summary Sync all the property to ES
// @Tags Property
// @Description Sync all the property to ES
// @ID sync-properties
// @Accept  json
// @Produce  json
// @Security Authorization
// @Success 200 {object} interface{} "Successfully"
// @Router /property/sync [put]
func (h *PropertyHandler) SyncProperties(ctx *gin.Context) {
	propertyCtx := context.Background()

	h.grpcClient.SyncProperties(propertyCtx, &proto.SyncPropertiesRequest{})

	ctx.JSON(200, gin.H{
		"Success": true,
	})
}
