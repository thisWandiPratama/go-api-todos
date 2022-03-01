package pengajuan

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Pengajuan, error)
	FindBuktiAll(ID int) ([]BuktiJaminan, error)
	FindByID(ID int) (Pengajuan, error)
	FindJaminanBarangByID(ID int) (JaminanBarang, error)
	FindAllByStatusDraf() ([]Pengajuan, error)
	UpdateJaminanPengajuan(jaminanbarang JaminanBarang) (JaminanBarang, error)
	UpdateStatusDraf(statusdraf Pengajuan) (Pengajuan, error)
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

func (r *repository) FindJaminanBarangByID(ID int) (JaminanBarang, error) {
	var user JaminanBarang

	err := r.db.Order("id desc").Where("id_pengajuan = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAllByStatusDraf() ([]Pengajuan, error) {
	var user []Pengajuan

	err := r.db.Where("status_draf = ?", 1).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateJaminanPengajuan(jaminanbarang JaminanBarang) (JaminanBarang, error) {

	err := r.db.Create(&jaminanbarang).Error
	if err != nil {
		return jaminanbarang, err
	}
	return jaminanbarang, nil
}

func (r *repository) UpdateStatusDraf(statusdraf Pengajuan) (Pengajuan, error) {

	err := r.db.Where("no_pengajuan = ?&no_anggota=?", statusdraf.NoPengajuan, statusdraf.NoAnggota).Save(&statusdraf).Error
	if err != nil {
		return statusdraf, err
	}
	return statusdraf, nil
}

func (r *repository) AddBuktiJaminan(bukti BuktiJaminan) (BuktiJaminan, error) {
	err := r.db.Create(&bukti).Error
	if err != nil {
		return bukti, err
	}
	return bukti, nil
}
