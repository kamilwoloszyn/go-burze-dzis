package requestmod

const SoapURL = "https://burze.dzis.net/soap.php"

var DefinitionHead = AttrMod{
	Name:     "Envelope",
	AttrName: "xmlns",
	Value:    "http://schemas.xmlsoap.org/soap/envelope/",
}

var DefinitionAttrs = []AttrMod{
	{
		Name:     "KeyAPI",
		AttrName: "xmlns",
		Value:    SoapURL,
	},
	{
		Name:     "miejscowosc",
		AttrName: "xmlns",
		Value:    SoapURL,
	},
	{
		Name:     "miejscowosci_lista",
		AttrName: "xmlns",
		Value:    SoapURL,
	},
	{
		Name:     "szukaj_burzy",
		AttrName: "xmlns",
		Value:    SoapURL,
	},
	{
		Name:     "ostrzezenia_pogodowe",
		AttrName: "xmlns",
		Value:    SoapURL,
	},
}

// Removes all structs names included by xml encoder
var DefinitonRemoveParams = []AttrMod{
	{
		Name: "APIKeyRequest",
	},
	{
		Name: "CityLocationRequest",
	},
	{
		Name: "CitiesRequest",
	},
	{
		Name: "StormSearchRequest",
	},
	{
		Name: "WeatherAlertRequest",
	},
}
