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
