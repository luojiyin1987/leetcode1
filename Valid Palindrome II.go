func validPalindrome(s string) bool {
    ls := len(s)
    if ls == 1 {
        return false
    }
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    reS := string(runes)
    if s == reS {
        return true
    }
    
    for i :=0 ; i<ls ; i++ {
        if s[i] != reS[i] {
            return s[i:ls-i-1] == reS[i+1:ls-i]  || reS[i:ls-i-1] == s[i+1:ls-i]
        }
    }
    return false
    
}
//------------双指针 （滑动窗口） 应该是原来的50%  或者更少
func validPalindrome(s string) bool {
    bytes:=[]byte(s)
    start := 0
    end :=len(bytes)-1
    for start<end{
        if bytes[start]!=bytes[end]{
            return isPalindrome(bytes,start+1,end)||isPalindrome(bytes,start,end-1)
        }
        start++
        end--
    }
    return true
}

func isPalindrome( s []byte,l,r int) bool{
    for l<r{
        if s[l]!=s[r]{
            return false
        }
        l++
        r--
    }
    return true
}
