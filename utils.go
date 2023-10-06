package gopc

import "sort"

func TestApiCredentials(apiKey string, env string) (bool, error) {
	c := Client{}
	c.Init(apiKey, env)

	_, err := c.GetPaymentDetails("unknown")

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
