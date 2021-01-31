package data

import (
	"fmt"
	"encoding/json"
)

// Status of an Account
type Status int

const (
	failed = iota
	pending
	confirmed
)

func (s Status) String() string {
	return [...]string{"failed", "pending", "confirmed"}[s]
}

// Classification of an Account
type Classification int

// Account classification
const (
	Personal = iota
	Business
)

func (s Classification) String() string {
	return [...]string{"Personal", "Business"}[s]
}


// Actor of an Organization
type Actor struct {
	Name      [4]string `json:"name" validate:"omitempty,alpha,dive"`
	BirthDate string `json:"birth_date" validate:"date"`
	Residency string `json:"residency" validate:"uppercase,len=2"`
}

// AccountHolder as a person
type AccountHolder struct {
	Identification string   `json:"identification" validate:"unique"`
	BirthDate      string   `json:"birth_date" validate:"date"`
	BirthCountry   string   `json:"birth_country" validate:"uppercase,len=2"`
	Address        []string `json:"address" validate:"omitempty,alphanum,dive"`
	City           string   `json:"city" validate:"alpha"`
	Country        string   `json:"country" validate:"uppercase,len=2"`
}

// AccountHolderAsOranization account holder
type AccountHolderAsOranization struct {
	AccountHolder AccountHolder `json:"account_holder"`
	Actor         []Actor       `json:"actors"`
}

// AccountPair for capture account relationship
type AccountPair struct {
	Type string `json:"type" validate:"alpha"`
	ID string `json:"id" validate:"unique"`
}

// Relationships of the account
type Relationships struct {
	AccountEvents []AccountPair `json:"account_events"`
	MasterAccount []AccountPair `json:"master_account"`
}

// Account details definition
type Account struct {
	Country                    string                     `json:"country" validate:"uppercase,len=2"`
	BaseCurrency               string                     `json:"base_currency" validate:"uppercase,len=3"`
	BankID                     string                     `json:"bank_id" validate:"uppercase,max=11"`
	BankIDCode                 string                     `json:"bank_id_code" validate:"uppercase"`
	AccountNumber              string                     `json:"account_number"`
	BIC                        string                     `json:"bic" validate:"alphanum,len=8|len=11"`
	IBAN                       string                     `json:"iban" validate:"alphanum"`
	CustomerID                 string                     `json:"customer_id"`
	Name                       [4]string                  `json:"name" validate:"omitempty,alpha,dive"`
	AlternativeNames           [3]string                  `json:"alternative_names" validate:"omitempty,max=140,dive"`
	AccountClassification      Classification              `json:"account_classification"`
	JointAccount               bool                       `json:"joint_account"`
	AccountMatchingOptOut      bool                       `json:"account_matching_opt_out"`
	SecondaryIdentification    string                     `json:"secondary_identification"`
	Switched                   bool                       `json:"switched"`
	Status                     Status                     `json:"status"`
	PrivateIdentification      AccountHolder              `json:"private_identification"`
	OrganisationIdentification AccountHolderAsOranization `json:"organisation_identification"`
	Relationships              Relationships              `json:"relationships"`
}

// Address houses a users address information
type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

type User struct {
	FirstName      string     `validate:"required,alpha"`
	LastName       string     `validate:"required,alpha"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
	AddressLines   []string   `json:"address-lines" validate:"omitempty,alpha,dive"`
	Name      [4]string `json:"name" validate:"omitempty,dive,alphanum"`
	BirthDate string `json:"birth_date" validate:"date"`
}

func (a *Account) String() string {
	return fmt.Sprintf("%s", a)
}

func (a Account) Json() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic (err)
	}
	return fmt.Sprintf("%s", b)
}

// Accounts list
type Accounts []*Account
