package structx

type Address struct {
	Village  string
	Province string
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Address
}

func StructInStructExample() string {
	ubaid := User{
		ID:        1,
		FirstName: "Ubaidillah",
		LastName:  "Hakim",
		Address: Address{
			Village:  "kalisampurno",
			Province: "east java",
		},
	}

	return ubaid.Address.Village
}
