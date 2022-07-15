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
}

// Removes all structs names included by xml encoder
var DefinitonRemoveParams = []AttrMod{
	{
		Name: "APIKeyRequest",
	},
}
