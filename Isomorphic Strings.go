func isIsomorphic(s string, t string) bool {
    for i :=0 ; i <len(s) -1; i++ {
        for j :=i+1; j <len(s); j++ {
            if s[i] == s[j]   {
                if t[i] == t[j] {
                    continue
                } else {
                    return false
                }
            }
            
            if t[i] == t[j] {
                if s[i] == s[j] {
                    continue
                }else {
                    return false
                }
            }
        }
    }
    return true
}
//----------------------------------------------------------
//map is cool 
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
}
