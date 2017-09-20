func repeatedSubstringPattern(s string) bool {
    lenght := len(s)
    var pos []int
    for i := lenght/2 ; i >= 1; i-- {
        if  lenght % i == 0 {
            pos = append(pos, i)
        }
    }
    if len(pos) <1 {
        return false
    }
    ret  := false
    runes := []rune(s)
    for  _,v := range pos {
        fmt.Println(v)
        ret = isSame(v, runes)
        if ret {
            return ret
        }
        
    }
    return  ret
}

func  isSame(l int, runes []rune) bool {
    for i :=0 ; i<len(runes)-l; i= i+l {
        if strings.EqualFold(string(runes[i:i+l]), string(runes[i+l:i+l+l]))  {
            continue 
        } else {
            return false
        }
        fmt.Println(i)
        
    }
    return  true
}
//--------------------------------
func repeatedSubstringPattern(s string) bool {
	length := len(s)
	if length == 1 {
		return false
	}
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}
		
        newStr := strings.Repeat(s[0:i], length/i)//use Repeat , so cool 
		if newStr == s {
			return true
		}
	}
	return false
}
