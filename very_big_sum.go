// HackerRank
// https://www.hackerrank.com/challenges/a-very-big-sum/problem?isFullScreen=true

func aVeryBigSum(ar []int64) int64 {
    // Write your code here
    
    var sum int64 = 0
    
    for _, number := range ar {
        sum += number
    }
    
    return sum
}
