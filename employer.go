package checkr

// Employer ...
// https://docs.checkr.com/#employer
type Employer struct {
	ID           string `json:"id"`
	Object       string `json:"object"`
	URI          string `json:"uri"`
	CandidateID  string `json:"candidate_id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Salary       int    `json:"salary"`
	ContractType string `json:"contract_type"`
	DoNotContact bool   `json:"do_not_contact"`
	StartDate    string `json:"start_date"`
	EmployerURL  string `json:"employer_url"`
	EndDate      string `json:"end_date"`
	Address      struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		State   string `json:"state"`
		Zipcode string `json:"zipcode"`
		Country string `json:"country"`
	} `json:"address"`
	Manager struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Title string `json:"title"`
	} `json:"manager"`
}
