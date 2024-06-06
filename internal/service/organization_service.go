package service

import (
	"go_api/internal/domain"
	"go_api/internal/repository"
)

type OrganizationService interface {
    CreateOrganization(org *domain.Organization) error
    GetOrganizations() ([]domain.Organization, error)
    GetOrganizationByID(id uint) (*domain.Organization, error)
    UpdateOrganization(org *domain.Organization) error
    DeleteOrganization(id uint) error
}

type organizationService struct {
    repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
    return &organizationService{repo}
}

func (s *organizationService) CreateOrganization(org *domain.Organization) error {
    return s.repo.CreateOrganization(org)
}

func (s *organizationService) GetOrganizations() ([]domain.Organization, error) {
	return s.repo.GetOrganizations()
}

func (s *organizationService) GetOrganizationByID(id uint) (*domain.Organization, error) {
	return s.repo.GetOrganizationByID(id)
}

func (s *organizationService) UpdateOrganization(org *domain.Organization) error {
	var orgDB, err = s.repo.GetOrganizationByID(org.ID)

	if err != nil {
		return err
	}

	orgDB.Name = org.Name
	orgDB.Document = org.Document
	
	return s.repo.UpdateOrganization(orgDB)
}

func (s *organizationService) DeleteOrganization(id uint) error {
	var orgDB, err = s.repo.GetOrganizationByID(id)

	if err != nil {
		return err
	}

	if orgDB == nil {
		return nil
	}

	return s.repo.DeleteOrganization(id)
}
