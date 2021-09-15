package bedtypes

import (
	"context"
	"peduli-covid/businesses/bedtypes"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(conn *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: conn,
	}
}

func (nr *PostgresRepository) FindAll(ctx context.Context) ([]bedtypes.Domain, error) {
	rec := []Bedtypes{}

	err := nr.Conn.Find(&rec).Error
	if err != nil {
		return []bedtypes.Domain{}, err
	}

	var domainBedTypes []bedtypes.Domain
	for _, value := range rec {
		domainBedTypes = append(domainBedTypes, value.toDomain())
	}
	return domainBedTypes, nil
}

func (nr *PostgresRepository) FindByHospitalID(ctx context.Context, hospitalID int) ([]bedtypes.Domain, error) {
	rec := []Bedtypes{}

	err := nr.Conn.Where("hospital_id = ?", hospitalID).Find(&rec).Error
	if err != nil {
		return []bedtypes.Domain{}, err
	}

	var domainBedtypes []bedtypes.Domain
	for _, value := range rec {
		domainBedtypes = append(domainBedtypes, value.toDomain())
	}
	return domainBedtypes, nil
}

func (nr *PostgresRepository) FindByTitleAndHospitalID(ctx context.Context, title string, hospitalID int) (bedtypes.Domain, error) {
	rec := Bedtypes{}
	err := nr.Conn.Where("name = ? AND hospital_id = ?", title, hospitalID).Find(&rec).Error
	if err != nil {
		return bedtypes.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (bedtypes.Domain, error) {
	rec := Bedtypes{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return bedtypes.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) UpdateAvailableBed(ctx context.Context, bedtypeDomain *bedtypes.Domain) error {
	rec := fromDomain(*bedtypeDomain)

	result := nr.Conn.Table("bedtypes").Where("id = ?", rec.ID).Update("bed_available", &rec.BedAvailable)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *PostgresRepository) UpdateBedEmpty(ctx context.Context, bedtypeDomain *bedtypes.Domain) error {
	rec := fromDomain(*bedtypeDomain)

	result := nr.Conn.Table("bedtypes").Where("id = ?", rec.ID).Update("bed_empty", &rec.BedEmpty)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *PostgresRepository) Store(ctx context.Context, bedtypeDomain *bedtypes.Domain) error {
	rec := fromDomain(*bedtypeDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
