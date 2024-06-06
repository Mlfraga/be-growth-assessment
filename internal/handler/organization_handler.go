package handler

import (
	"go_api/internal/domain"
	"go_api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
    service service.OrganizationService
}

func NewOrganizationHandler(service service.OrganizationService) *OrganizationHandler {
    return &OrganizationHandler{service}
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
    var org domain.Organization
    if err := c.ShouldBindJSON(&org); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateOrganization(&org); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
    var orgs []domain.Organization
    
	if err := c.ShouldBindJSON(&orgs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
	orgs, err := h.service.GetOrganizations()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusCreated, orgs)
}

func (h *OrganizationHandler) GetOrganizationByID(c *gin.Context) {
	idStr := c.Params.ByName("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	
	if err != nil {
		// Handle error if the conversion fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	idUint := uint(idUint64)

	org, err := h.service.GetOrganizationByID(idUint)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	var org domain.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateOrganization(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	idStr := c.Params.ByName("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	
	if err != nil {
		// Handle error if the conversion fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	idUint := uint(idUint64)

	err = h.service.DeleteOrganization(idUint)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Organization deleted successfully"})
}


