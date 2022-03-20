package mahasiswa

type Service interface {
	AddMahasiswa(mahasiswa AddMahasiswa) (Mahasiswa, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddMahasiswa(mahasiswa AddMahasiswa) (Mahasiswa, error) {
	Mahasiswa := Mahasiswa{}
	Mahasiswa.Npm = mahasiswa.Npm
	Mahasiswa.Nama = mahasiswa.Nama
	Mahasiswa.NmMk = mahasiswa.NmMk
	Mahasiswa.Ntugas = mahasiswa.Ntugas
	Mahasiswa.Nquiz = mahasiswa.Nquiz
	Mahasiswa.Nuts = mahasiswa.Nuts
	Mahasiswa.Nuas = mahasiswa.Nuas
	var nilaiTugas = (10 * mahasiswa.Ntugas) / 100
	var nilaiQuiz = (20 * mahasiswa.Nquiz) / 100
	var nilaiUts = (30 * mahasiswa.Nuts) / 100
	var nilaiUas = (40 * mahasiswa.Nuas) / 100
	Mahasiswa.Total = (nilaiTugas + nilaiQuiz + nilaiUts + nilaiUas)

	notifikasi, err := s.repository.AddMahasiswa(Mahasiswa)
	if err != nil {
		return notifikasi, err
	}
	return notifikasi, nil
}
