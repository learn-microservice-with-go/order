package model

type Office struct {
	OfficeCode   string `gorm:"column:officeCode" json:"office_code"`
	City         string `gorm:"column:city" json:"city"`
	Phone        string `gorm:"column:phone" json:"phone"`
	AddressLine1 string `gorm:"column:addressLine1" json:"address_line1"`
	AddressLine2 string `gorm:"column:addressLine2" json:"address_line2"`
	State        string `gorm:"column:state" json:"state"`
	Country      string `gorm:"column:country" json:"country"`
	PostalCode   string `gorm:"column:postalCode" json:"postal_code"`
	Territory    string `gorm:"column:territory" json:"territory"`
}
