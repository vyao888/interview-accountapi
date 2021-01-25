package domain

import (
	"https://github.com/go-playground/validator"
)

// Status of an account
type Status string

const (
	failed
	pending
	confirmed
)


// Actor of an Organization
type Actor struct {
	name      string
	dob       string `validate:"datetime"`
	residency string `validate: "uppercase, len=2"`
}

// AccountHolder as a person
type AccountHolder struct {
	identification string `json:"id" validate:"unique"`
	document_id     string `json:"documentId"`
	birth_date      string `json:"dob" validate:"datetime"`
	birth_country   string `json:"birthCountry" validate:"uppercase, len=2"`
	address        []string `json:"address" validate:"alphanum"`
	city           string `json:"city" validate:"alpha"`
	country        string `json:"country" validate:"uppercase, len=2"`
}

// AccountHolderAsOranization account holder
type AccountHolderAsOranization struct {
	accountHolder AccountHolder 
	actors         Actor         
}

// Relationships of the account
type Relationships struct {
	accountEvents []string
	masterAccount []string
}

// Acccount details definition
type Account struct {
	country                    string                     `validate:"uppercase, len=2"`
	base_currency               string                     `json:"currency" validate:"uppercase, len=3"`
	bank_id                     string                     `json:"bankId" validate:"uppercase, max=11"`
	bank_id_code                 string                     `json:"bankIdCode" validate:"uppercase"`
	account_number              string                     `json:"accountNumber"`
	bic                        string                     `validate:"alphanum, min=8, max=11"`
	iban                       string                     `validate:"alphanum"`
	customer_id                 string                     `json:"cunstomerId"`
	alternative_names           [3]string                  `json:"alternativeNames" validate:"max=140"`
	account_classification      string                     `json:"accountClassification" `
	joint_account               bool                       `json:"jointAccount"`
	account_matching_opt_out      bool                       `json:"accountMatchingOptOut"`
	secondary_ientification     string                     `json:"secondaryIdentification"`
	wwitched                   bool                       
	status                     Status                    
	private_identification      AccountHolder              `json:"privateIdentification"`
	organisation_identification AccountHolderAsOranization `json:"organisationIdentification"`
	relationships              Relationships              
}

