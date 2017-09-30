func isPalindrome(s string) bool {
    s  = strings.ToLower(s)
    ls := len(s)
    runes :=[]rune(s) 
    for i,j :=0, len(s)-1; i <j ; {
        if  !isLowerLetterOrNumber(runes[i])   {
            i++
            ls--
            continue
        }
         if  !isLowerLetterOrNumber(runes[j])  {
            j-- 
             ls --
            continue
        }
        if ls == 0 {
            return true
        }
        if runes[i] != runes[j] {
            return false
        }
        i++ 
        j-- 
        
    }
    return true
}

func isLowerLetterOrNumber(r rune) bool {
    if  r>=97 && r <=122 {
        return true
    } 
    if r>=48 && r <=57 {
        return true
    }
    return false
}

