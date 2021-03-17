package product

type OutputProducts struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func FormatOutputProducts(products []Product) []OutputProducts {

	var allProducts []OutputProducts

	for _, x := range products {
		formatterProduct := OutputProducts{
			Name:  x.Name,
			Price: x.Price,
		}

		allProducts = append(allProducts, formatterProduct)
	}

	return allProducts
}
