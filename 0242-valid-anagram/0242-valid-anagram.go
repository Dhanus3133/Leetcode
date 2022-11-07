func sortString(s string) string {
	temp := strings.Split(s, "")
	sort.Strings(temp)
	return strings.Join(temp, "")
}

func isAnagram(s string, t string) bool {
    if sortString(s) == sortString(t){
        return true
    }
    return false
}