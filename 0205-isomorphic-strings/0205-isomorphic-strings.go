func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	n := make(map[byte]byte)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if ctr, pres := m[s[i]]; pres {
			if ctr != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}

		if ctr, pres := n[t[i]]; pres {
			if ctr != s[i] {
				return false
			}
		} else {
			n[t[i]] = s[i]
		}
	}
	return true
}
