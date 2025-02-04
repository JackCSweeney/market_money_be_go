package markets

type Market struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	VendorCount int `json:"vendor_count"`
}