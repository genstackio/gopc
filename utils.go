package gopc

import "sort"

func TestApiCredentials(publicVendorToken string, privateVendorToken string, env string) (bool, error) {
	c := Client{}
	c.Init(publicVendorToken, privateVendorToken, env)
	_, err := c.GetB2CBalance()

	if err != nil {
		return false, err
	}

	return true, nil
}

func buildSignature(m *map[string]string, privateKey string) string {
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := ""
	for _, k := range keys {
		if len(s) > 0 {
			s = s + "&"
		}
		s = s + k + "=" + (*m)[k]
	}

	return s + "&" + privateKey
}
