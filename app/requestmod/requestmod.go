package requestmod

import (
	"fmt"
	"strings"
)

type RequestModifier struct {
	head     *AttrMod
	attrMods *[]AttrMod
	toRemove *[]AttrMod
}

func NewRequestModifier() RequestModifier {
	return RequestModifier{
		head:     &DefinitionHead,
		attrMods: &DefinitionAttrs,
		toRemove: &DefinitonRemoveParams,
	}
}

func (r *RequestModifier) ModifyRequest(xmlResultByte []byte) []byte {
	xmlResultString := string(xmlResultByte)
	if r.attrMods != nil && len(*r.attrMods) > 0 {
		for _, v := range *r.attrMods {
			xmlResultString = strings.ReplaceAll(
				xmlResultString,
				buildAttr(AttrMod{
					Name: v.Name,
				}),
				buildAttr(v),
			)
		}
	}

	if r.head != nil {
		xmlResultString = fmt.Sprintf(
			"%s"+xmlResultString+"%s",
			buildAttr(*r.head),
			buildAttr(
				AttrMod{
					Name:   r.head.Name,
					Ending: true,
				}),
		)
	}

	if r.toRemove != nil && len(*r.toRemove) > 0 {
		for _, v := range *r.toRemove {
			xmlResultString = strings.ReplaceAll(
				xmlResultString,
				buildAttr(
					AttrMod{
						Name: v.Name,
					},
				),
				"",
			)
			xmlResultString = strings.ReplaceAll(
				xmlResultString,
				buildAttr(
					AttrMod{
						Name:   v.Name,
						Ending: true,
					},
				),
				"",
			)
		}

	}

	return []byte(xmlResultString)
}

type AttrMod struct {
	Name     string
	AttrName string
	Value    string
	Ending   bool
}

func buildAttr(attr AttrMod) string {
	if attr.AttrName != "" {
		return fmt.Sprintf(
			"<%s %s=\"%s\">", attr.Name, attr.AttrName, attr.Value,
		)
	}
	if attr.Ending {
		return fmt.Sprintf(
			"</%s>", attr.Name,
		)
	}
	return fmt.Sprintf(
		"<%s>", attr.Name,
	)
}
