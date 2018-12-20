package google

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Parse an import id extracting field values using the given list of regexes.
// They are applied in order. The first in the list is tried first.
//
// e.g:
// - projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/subnetworks/(?P<name>[^/]+) (applied first)
// - (?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+),
// - (?P<name>[^/]+) (applied last)
func parseImportId(idRegexes []string, d TerraformResourceData, config *Config) error {
	for _, idFormat := range idRegexes {
		re, err := regexp.Compile(idFormat)

		if err != nil {
			return fmt.Errorf("Import is not supported. Invalid regex formats.")
		}

		if fieldValues := re.FindStringSubmatch(d.Id()); fieldValues != nil {
			// Starting at index 1, the first match is the full string.
			for i := 1; i < len(fieldValues); i++ {
				fieldName := re.SubexpNames()[i]
				// This part looks confusing.  Because there is no way to know at
				// this point whether 'fieldName' corresponds to a TypeString or a
				// TypeInteger in the resource schema, we need to determine
				// whether to call d.Set() with 'fieldValues[i]', or with an integer
				// parsed from 'fieldValues[i]'.  Normally, we would be able to just
				// use a try/catch pattern - try as a string, and if that doesn't
				// work, try as an integer, and if that doesn't work, return the
				// error.  Unfortunately, this is not possible here - during tests,
				// d.Set(...) will panic if there is an error.  So we try d.Get(),
				// relying on the behavior that an unset field will return the empty
				// value for the appropriate type.  So if the result is a number, we'll
				// set it as a number, but if it's a string, we'll set it as a string.
				val, _ := d.GetOk(fieldName)
				if _, ok := val.(string); ok {
					if err = d.Set(fieldName, fieldValues[i]); err != nil {
						return err
					}
				} else if _, ok := val.(int); ok {
					if atoi, atoiErr := strconv.Atoi(fieldValues[i]); atoiErr == nil {
						// If the value can be parsed as an integer, we try to set the
						// value as an integer.
						if err = d.Set(fieldName, atoi); err != nil {
							return err
						}
					} else {
						return fmt.Errorf("%s appears to be an integer, but %v cannot be parsed as an int.", fieldValues[i])
					}
				} else {
					return fmt.Errorf(
						"We cannot handle %s, which currently has value %v, and should be set to %#v, during import.", fieldName, val, fieldValues[i])
				}
			}

			// The first id format is applied first and contains all the fields.
			err := setDefaultValues(idRegexes[0], d, config)
			if err != nil {
				return err
			}

			return nil
		}
	}
	return fmt.Errorf("Import id %q doesn't match any of the accepted formats: %v", d.Id(), idRegexes)
}

func setDefaultValues(idRegex string, d TerraformResourceData, config *Config) error {
	if _, ok := d.GetOk("project"); !ok && strings.Contains(idRegex, "?P<project>") {
		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		d.Set("project", project)
	}
	if _, ok := d.GetOk("region"); !ok && strings.Contains(idRegex, "?P<region>") {
		region, err := getRegion(d, config)
		if err != nil {
			return err
		}
		d.Set("region", region)
	}
	if _, ok := d.GetOk("zone"); !ok && strings.Contains(idRegex, "?P<zone>") {
		zone, err := getZone(d, config)
		if err != nil {
			return err
		}
		d.Set("zone", zone)
	}
	return nil
}
