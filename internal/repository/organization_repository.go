package repository

import (
	"go_api/internal/domain"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
    CreateOrganization(org *domain.Organization) error
    GetOrganizations() ([]domain.Organization, error)
    GetOrganizationByID(id uint) (*domain.Organization, error)
    UpdateOrganization(org *domain.Organization) error
    DeleteOrganization(id uint) error
}

type organizationRepository struct {
    db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
    return &organizationRepository{db}
}

func (r *organizationRepository) CreateOrganization(org *domain.Organization) error {
    return r.db.Create(org).Error
}

func (r *organizationRepository) GetOrganizations() ([]domain.Organization, error) {
    var orgs []domain.Organization
    return orgs, r.db.Find(&orgs).Error
}

func (r *organizationRepository) GetOrganizationByID(id uint) (*domain.Organization, error) {
    var org domain.Organization
    return &org, r.db.Where("id = ?", id).First(&org).Error
}

func (r *organizationRepository) UpdateOrganization(org *domain.Organization) error {
    return r.db.Save(org).Error
}

func (r *organizationRepository) DeleteOrganization(id uint) error {
    return r.db.Delete(&domain.Organization{}, id).Error
}
