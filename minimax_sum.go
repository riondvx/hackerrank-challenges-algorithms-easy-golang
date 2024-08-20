// HackerRank
// https://www.hackerrank.com/challenges/mini-max-sum/problem?isFullScreen=true

func miniMaxSum(arr []int32) {
    // Write your code here
    var sum int64 = 0
    max_number := arr[0]
    min_number := arr[0]
    
    for _, number := range arr {
        sum += int64(number)
        
        if number > max_number {
            max_number = number
        }
        if number < min_number {
            min_number = number
        }
    }
    
    max_sum := sum - int64(min_number)
    min_sum := sum - int64(max_number)
    
    fmt.Println(min_sum, max_sum)
}
