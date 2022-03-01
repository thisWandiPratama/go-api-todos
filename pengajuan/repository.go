package pengajuan

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Pengajuan, error)
	FindBuktiAll(ID int) ([]BuktiJaminan, error)
	FindByID(ID int) (Pengajuan, error)
	UpdateJaminanPengajuan(jaminan Pengajuan) (Pengajuan, error)
	AddBuktiJaminan(bukti BuktiJaminan) (BuktiJaminan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Pengajuan, error) {
	var user []Pengajuan

	err := r.db.Order("no_pengajuan desc").Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindBuktiAll(ID int) ([]BuktiJaminan, error) {
	var user []BuktiJaminan

	err := r.db.Where("id_pengajuan = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (Pengajuan, error) {
	var user Pengajuan

	err := r.db.Where("no_pengajuan = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateJaminanPengajuan(jaminan Pengajuan) (Pengajuan, error) {

	err := r.db.Where("no_pengajuan = ?", jaminan.NoPengajuan).Save(&jaminan).Error
	if err != nil {
		return jaminan, err
	}
	return jaminan, nil
}

func (r *repository) AddBuktiJaminan(bukti BuktiJaminan) (BuktiJaminan, error) {
	err := r.db.Create(&bukti).Error
	if err != nil {
		return bukti, err
	}
	return bukti, nil
}
