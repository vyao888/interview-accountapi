package domain

// Status of an Account
type Status int

const (
	failed = iota
	pending
	confirmed
)

// Classification of an Account
type Classification int

// Account classification
const (
	Personal = iota
	Business
)

// Actor of an Organization
type Actor struct {
	Name      string `json:"name" validate:"alpha"`
	BirthDate string `json:"birth_date" validate:"datetime"`
	Residency string `json:"residency" validate:"uppercase,len=2"`
}

// AccountHolder as a person
type AccountHolder struct {
	Identification string   `json:"identification" validate:"unique"`
	BirthDate      string   `json:"birth_date" validate:"datetime"`
	BirthCountry   string   `json:"birth_country" validate:"uppercase, len=2"`
	Address        []string `json:"address" validate:"alphanum"`
	City           string   `json:"city" validate:"alpha"`
	Country        string   `json:"country" validate:"uppercase, len=2"`
}

// AccountHolderAsOranization account holder
type AccountHolderAsOranization struct {
	AccountHolder AccountHolder `json:"account_holder"`
	Actor         []Actor       `json:"actors"`
}

// Relationships of the account
type Relationships struct {
	AccountEvents []string `json:"account_events"`
	MasterAccount []string `json:"master_account"`
}

// Account details definition
type Account struct {
	Country                    string                     `json:"country" validate:"uppercase, len=2"`
	BaseCurrency               string                     `json:"base_currency" validate:"uppercase, len=3"`
	BankID                     string                     `json:"bank_id" validate:"uppercase, max=11"`
	BankIDCode                 string                     `json:"bank_id_code" validate:"uppercase"`
	AccountNumber              string                     `json:"account_number"`
	BIC                        string                     `json:"bic" validate:"alphanum, min=8|11"`
	IBAN                       string                     `json:"iban" validate:"alphanum"`
	CustomerID                 string                     `json:"customer_id"`
	Name                       [4]string                  `json:"name" validate:"alpha"`
	AlternativeNames           [3]string                  `json:"alternative_names" validate:"max=140"`
	AccountClassification      string                     `json:"account_classification" validate:"alpha"`
	JointAccount               bool                       `json:"joint_account"`
	AccountMatchingOptOut      bool                       `json:"account_matching_opt_out"`
	SecondaryIdentification    string                     `json:"secondary_identification"`
	Switched                   bool                       `json:"switched"`
	Status                     Status                     `json:"status"`
	PrivateIdentification      AccountHolder              `json:"private_identification"`
	OrganisationIdentification AccountHolderAsOranization `json:"organisation_identification"`
	Relationships              Relationships              `json:"relationships"`
}

// Accounts list
type Accounts []*Account
