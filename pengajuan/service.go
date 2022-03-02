package pengajuan

import (
	"fmt"
	"strings"
	"time"
)

type Service interface {
	FindAll() ([]Pengajuan, error)
	FindNotifikasiAll() ([]Notifikasi, error)
	FindPersetujuanAll() ([]Pengajuan, error)
	SearchJamianAll(search SearchJaminan) ([]Pengajuan, error)
	FindByID(ID int) (Pengajuan, error)
	FindJaminanBarangByID(ID int) (JaminanBarang, error)
	FindJaminanTanahByID(ID int) (JaminanTanah, error)
	FindBuktiAll(ID int) ([]BuktiJaminan, error)
	FindAllByStatusDraf() ([]Pengajuan, error)
	UpdateJaminanPengajuan(jaminan UpdateJaminanPengajuan) (JaminanBarang, error)
	AddJaminanTanah(jaminan AddJaminanTanah) (JaminanTanah, error)
	AddBuktiJaminan(bukti AddBuktiJaminan) (BuktiJaminan, error)
	AddNotifikasi(notifikasi AddNotifikasi) (Notifikasi, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Pengajuan, error) {
	user, err := s.repository.FindAll()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) FindNotifikasiAll() ([]Notifikasi, error) {
	user, err := s.repository.FindNotifikasiAll()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) FindPersetujuanAll() ([]Pengajuan, error) {
	user, err := s.repository.FindPersetujuanAll()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) SearchJamianAll(search SearchJaminan) ([]Pengajuan, error) {
	user, err := s.repository.SearchJamianAll(search)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateJaminanPengajuan(jaminan UpdateJaminanPengajuan) (JaminanBarang, error) {
	var pengajuan JaminanBarang

	pengajuan.IdPengajuan = jaminan.IdPengajuan
	pengajuan.JenisBarangJaminan = jaminan.JenisBarangJaminan
	pengajuan.NamaBarangJaminan = jaminan.NamaBarangJaminan
	pengajuan.KondisiBarangJaminan = jaminan.KondisiBarangJaminan
	pengajuan.DeskripsiBarangJaminan = jaminan.DeskripsiBarangJaminan

	updatedPengajuan, err := s.repository.UpdateJaminanPengajuan(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) AddJaminanTanah(jaminan AddJaminanTanah) (JaminanTanah, error) {
	var pengajuan JaminanTanah

	pengajuan.IdPengajuan = jaminan.IdPengajuan
	pengajuan.Latitude = jaminan.Latitude
	pengajuan.Longitude = jaminan.Longitude
	pengajuan.Altitude = jaminan.Altitude
	pengajuan.ArahLokasi = jaminan.ArahLokasi
	pengajuan.NamaBarangJaminan = jaminan.NamaBarangJaminan
	pengajuan.KondisiBarangJaminan = jaminan.KondisiBarangJaminan
	pengajuan.DeskripsiBarangJaminan = jaminan.DeskripsiBarangJaminan

	updatedPengajuan, err := s.repository.AddJaminanTanah(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) FindByID(ID int) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(ID)
	if err != nil {
		return pengajuan, err
	}
	return pengajuan, nil
}

func (s *service) FindJaminanBarangByID(ID int) (JaminanBarang, error) {
	pengajuan, err := s.repository.FindJaminanBarangByID(ID)
	if err != nil {
		return pengajuan, err
	}
	return pengajuan, nil
}

func (s *service) FindJaminanTanahByID(ID int) (JaminanTanah, error) {
	pengajuan, err := s.repository.FindJaminanTanahByID(ID)
	if err != nil {
		return pengajuan, err
	}
	return pengajuan, nil
}

func (s *service) FindBuktiAll(ID int) ([]BuktiJaminan, error) {
	bukti, err := s.repository.FindBuktiAll(ID)
	if err != nil {
		return bukti, err
	}

	return bukti, nil
}

func (s *service) FindAllByStatusDraf() ([]Pengajuan, error) {
	bukti, err := s.repository.FindAllByStatusDraf()
	if err != nil {
		return bukti, err
	}

	return bukti, nil
}

func (s *service) AddBuktiJaminan(input AddBuktiJaminan) (BuktiJaminan, error) {
	buktiJaminan := BuktiJaminan{}
	buktiJaminan.IdPengajuan = input.IdPengajuan
	buktiJaminan.Bukti = input.Bukti

	bukti, err := s.repository.AddBuktiJaminan(buktiJaminan)
	fmt.Println(bukti)
	if err != nil {
		return bukti, err
	}
	return bukti, nil
}

func (s *service) AddNotifikasi(input AddNotifikasi) (Notifikasi, error) {
	Notifikasi := Notifikasi{}
	Notifikasi.IdPengajuan = input.IdPengajuan
	Notifikasi.Petugas = input.Petugas

	dt := time.Now()
	tgl := dt.Format("01-02-2006 15:04:05")
	arrayTgl := strings.Split(tgl, " ")
	Notifikasi.TglSurvey = arrayTgl[0]
	Notifikasi.Pukul = arrayTgl[1]

	notifikasi, err := s.repository.AddNotifikasi(Notifikasi)
	if err != nil {
		return notifikasi, err
	}
	return notifikasi, nil
}
