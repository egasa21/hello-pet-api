package repository

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/doctor_model"
)

type DoctorRepository struct {
	db *database.DB
}

func NewDoctorRepository(db *database.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (repo *DoctorRepository) CreateDoctor(doctor *doctor_model.Doctor) (err error) {
	err = repo.db.Database.Create(doctor).Error
	return
}

func (repo *DoctorRepository) GetDoctorById(doctor *doctor_model.Doctor, doctorID string) (err error) {
	err = repo.db.Database.First(doctor, "id=?", doctorID).Error
	return
}

func (repo *DoctorRepository) UpdateDoctor(doctor *doctor_model.Doctor) (err error) {
	err = repo.db.Database.Save(doctor).Error
	return
}

func (repo *DoctorRepository) GetAllDoctors(doctors []doctor_model.Doctor) (err error) {
	err = repo.db.Database.Find(&doctors).Error
	return
}

func (repo *DoctorRepository) DeleteDoctor(doctorId string) (err error) {
	err = repo.db.Database.Delete(&doctor_model.Doctor{}, "id=?", doctorId).Error
	return
}
