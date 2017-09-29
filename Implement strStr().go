func strStr(haystack string, needle string) int {
    lh := len(haystack)
    ln := len(needle)
    if ln == 0 {
        return 0
    }
    if lh  < ln {
        return -1
    }
    
    ncount := 0
    for i :=0; i<=lh-ln; i++  {
        //fmt.Println(0,i,ncount)
        if haystack[i] == needle[ncount] {
            //fmt.Println(1,i,ncount)
            for ncount <ln {
                if haystack[i+ncount] != needle[ncount] {
                    ncount = 0
                    break
                } 
                ncount++
                //fmt.Println(3,i,ncount)
            } 
            
        } 
        if ncount ==  ln {
            return i 
        } 
    }
    return -1
}
