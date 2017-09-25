func findAnagrams(s string, p string) []int {
    ls := len(s)
    lp := len(p)
    count  := lp
    var temp []int
    
    if ls == 0 ||  lp == 0 {
        return temp
    } 
    
    pMap := make(map[byte]int)
   
    
    for _, v :=range p {
        pMap[byte(v)]++
    }
    for i :=0 ; i<ls ; i++ {
        
        if pMap[s[i]] >=1 {
            count --
        }
        pMap[s[i]] --
        if i >=lp  {
            if pMap[s[i-lp]] >=0 {
                count ++
            }
            pMap[s[i-lp]] ++
        }
        if count == 0 {
            temp = append(temp, i-lp+1)
        }
    }
    return temp
 
 }
