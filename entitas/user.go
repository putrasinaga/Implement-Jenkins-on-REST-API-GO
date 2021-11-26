package entitas

type User struct {
	Id   int
	Nama string
	Umur int
}

func (u User) CekUmur() string {
	var ket string

	if u.Umur < 15 {
		ket = "anak-anak"
	} else if u.Umur < 17 {
		ket = "Remaja"
	} else {
		ket = "Dewasa"
	}

	return ket
}
