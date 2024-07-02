package controllers

import (
    "net/http"
    "teste-enube/services"
	"strconv"
    "github.com/gin-gonic/gin"
)

// GetDataByField godoc
// @Summary Get data by field
// @Description Retrieve paginated data from a specified field in the Excel sheet.
// @Tags data
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param field path string true "Name of the field to retrieve" Enums(PartnerId, PartnerName, CustomerId, CustomerName, CustomerDomainName, CustomerCountry, MpnId, Tier2MpnId, InvoiceNumber, ProductId, SkuId, AvailabilityId, SkuName, ProductName, PublisherName, PublisherId, SubscriptionDescription, SubscriptionId, ChargeStartDate, ChargeEndDate, UsageDate, MeterType, MeterCategory, MeterId, MeterSubCategory, MeterName, MeterRegion, Unit, ResourceLocation, ConsumedService, ResourceGroup, ResourceURI, ChargeType, UnitPrice, Quantity, UnitType, BillingPreTaxTotal, BillingCurrency, PricingPreTaxTotal, PricingCurrency, ServiceInfo1, ServiceInfo2, Tags, AdditionalInfo, EffectiveUnitPrice, PCToBCExchangeRate, PCToBCExchangeRateDate, EntitlementId, EntitlementDescription, PartnerEarnedCreditPercentage, CreditPercentage, CreditType, BenefitOrderId, BenefitId, BenefitType)
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Results per page" default(10)
// @Success 200 {object} map[string]interface{} "Successful retrieval of field data"
// @Failure 400 {object} map[string]string "Invalid page or limit number"
// @Failure 404 {object} map[string]string "Field not found or error in processing request"
// @Router /data/{field} [get]
func GetDataByField(c *gin.Context) {
    field := c.Param("field")
    if field == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No field provided"})
        return
    }

    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }

    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
        return
    }

    data, total, err := services.GetDataForField(field, page, limit)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Field not found or error in processing request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  data,
        "page":  page,
        "limit": limit,
        "total": total,
    })
}
