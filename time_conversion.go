// HackerRank
// https://www.hackerrank.com/challenges/time-conversion/problem?isFullScreen=true

func timeConversion(s string) string {
    // Write your code here
    periode := s[len(s)-2:]
    waktu_str := s[:len(s)-2]
    // jam, menit, detik (str)
    waktu_arr := strings.Split(waktu_str, ":")
    
    jam, _ := strconv.Atoi(waktu_arr[0])
    
    if periode == "AM" && jam == 12 {
        jam = 0
    } else if periode == "PM" && jam != 12 {
        jam += 12
    }
    
    return fmt.Sprintf("%02d:%s:%s", jam, waktu_arr[1], waktu_arr[2])
}
