package auth

type UserFormatter struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Username   string `json:"username"`
	Alamat     string `json:"alamat"`
	NoTelp     string `json:"no_telp"`
	Email      string `json:"email"`
	Pass       string `json:"pass"`
	StatusAkun string `json:"status_token"`
	Token      string `json:"token"`
}

func FormatUser(user DaftarPetugas) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Nama:       user.Nama,
		Username:   user.Username,
		Alamat:     user.Alamat,
		NoTelp:     user.NoTelp,
		Email:      user.Email,
		Pass:       user.Pass,
		StatusAkun: user.StatusAkun,
		Token:      "sdksk3kdk93ks.skmsodods.343asadcmmooopsd2#$@1",
	}

	return formatter
}