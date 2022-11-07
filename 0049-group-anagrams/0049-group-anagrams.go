func groupAnagrams(strs []string) [][]string {
  m := make(map[string][]string)
  for _, s := range strs {
    temp := strings.Split(s, "")
    sort.Strings(temp)
    anagram := strings.Join(temp, "")
    if arr, pres := m[anagram]; pres {
      m[anagram] = append(arr, s)
    } else {
      m[anagram] = []string{s}
    }
  }
  res := [][]string{}
  for _, v := range m {
    res = append(res, v)
  }
  return res
}