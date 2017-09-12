func isHappy(n int) bool {
    temp := 0
    for (n !=1 && n !=4) {
        for n >0  {
            temp  += (n%10) * (n %10)
            n = n/10
        }
        n = temp 
        temp = 0
    }
    return  n == 1
    
}
//-----------------------------------------------


func isHappy(n int) bool{
    for n >6 {
        ret := 0
        next := n 
        for next >0  {
            ret += (next%10) * (next %10)
            next = next /10
        }
        n = ret
    }
    return  n==1 
}

// --------The  c code  is cool--------------------------
int digitSquareSum(int n) {
    int sum = 0, tmp;
    while (n) {
        tmp = n % 10;
        sum += tmp * tmp;
        n /= 10;
    }
    return sum;
}

bool isHappy(int n) {
    int slow, fast;
    slow = fast = n;
    do {
        slow = digitSquareSum(slow);
        fast = digitSquareSum(fast);
        fast = digitSquareSum(fast);
    } while(slow != fast);
    if (slow == 1) return 1;
    else return 0;
}
