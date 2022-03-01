package pengajuan

type Service interface {
	FindAll() ([]Pengajuan, error)
	FindByID(ID int) (Pengajuan, error)
	FindBuktiAll(ID int) ([]BuktiJaminan, error)
	UpdateJaminanPengajuan(jaminan UpdateJaminanPengajuan) (Pengajuan, error)
	AddBuktiJaminan(bukti AddBuktiJaminan) (BuktiJaminan, error)
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

func (s *service) UpdateJaminanPengajuan(jaminan UpdateJaminanPengajuan) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(jaminan.NoPengajuan)
	if err != nil {
		return pengajuan, err
	}

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

func (s *service) FindByID(ID int) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(ID)
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

func (s *service) AddBuktiJaminan(input AddBuktiJaminan) (BuktiJaminan, error) {
	buktiJaminan := BuktiJaminan{}
	buktiJaminan.IdPengajuan = input.IdPengajuan
	buktiJaminan.Bukti = input.Bukti

	bukti, err := s.repository.AddBuktiJaminan(buktiJaminan)
	if err != nil {
		return bukti, err
	}

	return bukti, nil
}
