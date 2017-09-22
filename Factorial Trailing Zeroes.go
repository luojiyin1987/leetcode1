func trailingZeroes(n int) int {
    res := 0
    for n >0  {
        res = res + n /5
        n = n/5
    }
    return res 
}
//思路：求结果又多少个0，就是求在阶乘中 有多少个2 和5的因子， 2的个数肯定比5多，
//只要求有多少个5， 通过不断  n/5求出 5的个数
