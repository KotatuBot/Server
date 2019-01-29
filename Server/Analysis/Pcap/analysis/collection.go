package analysis

func Count(key []string, data []string) map[string]int {
	m := map[string]int{}
	for _, key_name := range key {
		m[key_name] = 0
	}

	for _, data_tip := range data {
		m[data_tip] += 1
	}

	return m
}

func Unique_info(iparray []string) []string {

	m := make(map[string]bool)
	uniq := []string{}

	for _, ele := range iparray {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	return uniq
}
