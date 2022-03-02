package auth

import "gorm.io/gorm"

type Repository interface {
	FindByEmail(email string) (DaftarPetugas, error)
	LupaPassword(user DaftarPetugas) (DaftarPetugas, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (DaftarPetugas, error) {
	var user DaftarPetugas

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) LupaPassword(user DaftarPetugas) (DaftarPetugas, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
