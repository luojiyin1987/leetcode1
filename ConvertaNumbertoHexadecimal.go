func toHex(num int) string { //what about int ? is int32  int64 ???
    if num == 0 {
        return "0"
    }
    if num == -1 {
        return "ffffffff"
    }
    
    var result string
    
    
    
//func getHexDigit(a int) (hex string){
//	if a < 10 {
//		return strconv.Itoa(a)
//	}
//
//	return string([]byte{'a' + byte(a - 10)})
//}
    Hexnum  := [16]string {"0", "1","2","3","4","5","6","7","8","9","a", "b","c","d","e","f"} //which  better
    
    if num < 0 {
        num = 16 * 16 * 16 * 16 * 16 * 16 * 16 * 16  + num 
     }
    for num >0 {
        result = Hexnum[num%16] + result
        num = num / 16
    }
    
    
    return result
}
