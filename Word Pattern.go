func wordPattern(pattern string, str string) bool {
    temp := strings.Fields(str) 
    if len(temp) !=len(pattern) {
        return false
    } 
   
    for i :=0 ; i<len(temp)-1 ; i++ {
        if pattern[i] != pattern[i+1] && temp[i] == temp[i+1] {
            return false
        }//受滑动窗口的启发
        
        for j :=i+1; j<len(temp); j++ {
            if  pattern[i] == pattern[j]  {
                if temp[i] == temp[j] {
                    continue
                } else {
                    return false
                }
            } 
                
        }
    }
    return true 
}
//-----使用map 
func wordPattern(pattern string, str string) bool {
       arr := strings.Split(str, " ")

    mm := make(map[byte]string)
    mm2 := make(map[string]bool)
    if len(arr) != len(pattern) {
        return false
    }
    for i := 0; i < len(pattern); i++ {
        k := pattern[i]
        v := arr[i]

        if mm[k] == "" {
            if mm2[v] {
                return false
            }

            mm[k] = v
            mm2[v] = true
        } else {
            if mm[k] != v {
                return false
            }
        }
    }
    
    return true

}
