package vendors

type Vendor struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ContactName    string `json:"contact_name"`
	ContactPhone   string `json:"contact_phone"`
	CreditAccepted bool   `json:"credit_accepted"`
}