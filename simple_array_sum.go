// HackerRank
// https://www.hackerrank.com/challenges/simple-array-sum/problem?isFullScreen=true

func simpleArraySum(ar []int32) int32 {
    // Write your code here
    var hasil int32 = 0
    
    for _, v := range ar {
        hasil += v
    }
    
    return hasil
}
