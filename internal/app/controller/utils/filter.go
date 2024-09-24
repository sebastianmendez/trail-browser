package utils

import (
	"net/url"

	"github.com/sebastianmendez/trail-browser/internal/store/model"
)

func ParseFilters(rawQuery string) map[string]string {
	// allow only access_name and park_spaces
	query, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil
	}

	accessNameFilter := query.Get(model.FaccessName)
	parkSpacesFilter := query.Get(model.FparkSpaces)

	res := map[string]string{}
	res[model.FparkSpaces] = parkSpacesFilter
	res[model.FaccessName] = accessNameFilter
	return res
}
