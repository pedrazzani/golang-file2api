package entity

type Product struct {
    DepCode string
    LabelDep string
    ProductType string 
    ItemCode string
    Value float32
    DaysToExpiration int
    DescLine1 string
    DescLine2 string
}

//NewProduct create a new Product
func NewProduct(depCode string, labelDep string, prodType string, itemCode string, value float32, days int, descLine1 string, descLine2 string) *Product {
	p := &Product{
		DepCode:        depCode,
		LabelDep:     labelDep,
		ProductType:    prodType,
		ItemCode:     itemCode,
		Value:  value,
		DaysToExpiration: days,
		DescLine1: descLine1,
		DescLine2: descLine2,
	}
	
	return p
}
