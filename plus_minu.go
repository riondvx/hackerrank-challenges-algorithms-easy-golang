// HackerRank
// https://www.hackerrank.com/challenges/plus-minus/problem?isFullScreen=true

func plusMinus(arr []int32) {
    // Write your code here
    var n_positive float64 = 0.0
    var n_negative float64 = 0.0
    var n_zero float64 = 0.0
    var len_arr float64 = float64(len(arr))
    
    for _, number := range arr {
        if number == 0 {
            n_zero += 1.0
        } else if number > 0 {
            n_positive += 1.0
        } else {
            n_negative += 1.0
        }
    }
    
    fmt.Println(n_positive / len_arr)
    fmt.Println(n_negative / len_arr)
    fmt.Println(n_zero / len_arr)
}
