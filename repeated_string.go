// HackerRank
// https://www.hackerrank.com/challenges/repeated-string/problem?isFullScreen=true


func repeatedString(s string, n int64) int64 {
    // Write your code here
    var n_af int64 = 0
    var n_al int64 = 0
    var n_str int64 = n / int64(len(s))
    var ln_last int64 = n % (n_str * int64(len(s)))
    
    var i int64 = 0
    
    for _, char := range s {
        if char == 'a' {
            n_af += 1
            
            if i < ln_last {
                n_al += 1
            }
        }
        
        i += 1
    }
    
    return (n_af * n_str) + n_al
}
