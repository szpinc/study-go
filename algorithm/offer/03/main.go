package main

func replaceSpace(s string) string {
	bytes := []byte{}

	for _, v := range []byte(s) {
		if v == ' ' {
			bytes = append(bytes, []byte("%20")...)
			continue
		}
		bytes = append(bytes, v)
	}
	return string(bytes)
}
