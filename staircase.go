// HackerRank
// https://www.hackerrank.com/challenges/staircase/problem?isFullScreen=true

func staircase(n int32) {
    // Write your code here
    var baris int32
    var kolom int32
    
    for baris = 1; baris <= n; baris++ {
        for kolom = 1; kolom <= n - baris; kolom++ {
            fmt.Print(" ")
        }
        for kolom = 1; kolom <= baris; kolom++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}
