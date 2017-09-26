func isValid(s string) bool {
    if len(s)%2 !=0 {
        return false
    }
    left := map[byte]struct{}{
        '(': struct{}{},
        '[': struct{}{},
        '{': struct{}{},
    }
    right := map[byte]struct{}{
        ')': struct{}{},
        ']': struct{}{},
        '}': struct{}{},
    }
    
    bytes := []byte(s)
    stack := []byte{}
    for  i := range bytes {
        if _, ok := left[bytes[i]]; ok {
            stack = append(stack,bytes[i])
        }else if _, ok = right[bytes[i]]; ok {
            if len(stack) == 0 {
                return false
            }
            switch stack[len(stack)-1]  {
                case '(':
                if bytes[i] != ')'{
                    return false
                }
                case '[':
                if bytes[i] !=']' {
                    return false
                }
                case '{':
                if bytes[i] !='}' {
                    return false
                }
            }
            stack = stack[:len(stack) -1]
            
        } else {
            return false
        }
    }
    return len(stack) == 0
}
//---------------------------------------
func isValid(s string) bool {
    var str []string
	//m := map[string]string{"{": "}", "(": ")", "[": "]"}
    m := map[string]string{"}": "{", ")": "(", "]": "["}
	for _, x := range s {
		if len(str) == 0 {
			str = append(str, string(x))
			continue
		}
		if str[len(str)-1] == m[string(x)] {
			str = str[:len(str)-1]
		} else {
			str = append(str, string(x))
		}
	}
	if len(str) == 0 {
		return true
	}
	return false
}
