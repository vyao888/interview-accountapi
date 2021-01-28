package data

func GetAccounts() Accounts {
	list := Accounts{}

	p := AccountHolder{
		Identification: "13YH458762",
		BirthDate:      "2017-07-23",
		BirthCountry:   "GB",
		Address:        []string{"10 Avenue des Champs"},
		City:           "London",
		Country:        "GB",
	}

	o := AccountHolder{
		Identification: "123654",
		BirthDate:      "1970-01-01",
		BirthCountry:   "GB",
		Address:        []string{"10 Avenue des Champs"},
		City:           "London",
		Country:        "GB",
	}

	a := []Actor{
		{
			Name:      [4]string{"Jeff Page"},
			BirthDate: "1970-01-01",
			Residency: "GB",
		},
	}

	b := AccountHolderAsOranization{AccountHolder: o, Actor: a}

	r := Relationships{
		[]AccountPair{
			{"account_event", "c1023677-70ee-417a-9a6a-e211241f1e9c"},
			{"account_event", "437284fa-62a6-4f1d-893d-2959c9780288"},},
		[]AccountPair{
			{"accounts", "a52d13a4-f435-4c00-cfad-f5e7ac5972df"}},
	}

	list = append(list,
		&Account{
			Country:                    "GB",
			BaseCurrency:               "GBP",
			BankID:                     "400300",
			BankIDCode:                 "GBDSC",
			AccountNumber:              "41426819",
			BIC:                        "NWBKGB22",
			IBAN:                       "GB11NWBK40030041426819",
			CustomerID:                 "",
			Name:                       [4]string{"Samantha Holder"},
			AlternativeNames:           [3]string{"Sam Holder"},
			AccountClassification:      "Personal",
			JointAccount:               true,
			AccountMatchingOptOut:      false,
			SecondaryIdentification:    "A1B2C3D4",
			Switched:                   false,
			Status:                     Status(2),
			PrivateIdentification:      p,
			OrganisationIdentification: b,
			Relationships:              r,
		})
		return list
}
