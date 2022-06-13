package domain

type CityLocation struct {
	CoordX float32 `xml:"x"`
	CoordY float32 `xml:"y"`
}

type Cities struct {
	Cities []string `xml:"miejscowosci_listaResponse"`
}
