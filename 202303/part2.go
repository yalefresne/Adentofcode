package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
    "strconv"
)

func main() {
    inputs_from_calendar, err := os.Open("./inputs.txt")
    check(err)

    defer inputs_from_calendar.Close()

    scanner := bufio.NewScanner(inputs_from_calendar)
    lines := make([]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    total := 0

    for i, line := range lines {
        star_rgx := regexp.MustCompile(`[\*]`)
        stars := star_rgx.FindAllStringIndex(line, -1)

        digit_rgx := regexp.MustCompile(`[\d]+`)
        for _, star := range stars {
            fmt.Println(star)
            nums_spotted := make([]int, 0)
            star_i := star[0]
            numbers := digit_rgx.FindAllStringIndex(lines[i], -1)
            if i > 0 {
                above_line := i-1
                numbers := digit_rgx.FindAllStringIndex(lines[above_line], -1)
                for _, num := range numbers {
                    if num[0] == star_i-1 || num[0] == star_i || num[0] == star_i+1 || num[1]-1 == star_i-1 || num[1]-1 == star_i || num[1]-1 == star_i+1 {
                        nums_spotted = append(nums_spotted, string_number_to_int(string(lines[above_line][num[0]:num[1]])))
                    }
                }
            }
            for _, num := range numbers {
                if num[0] == star_i+1 || num[1]-1 == star_i-1 {
                    nums_spotted = append(nums_spotted, string_number_to_int(string(lines[i][num[0]:num[1]])))
                }
            }
            if i < len(lines) {
                bellow_line := i+1
                numbers := digit_rgx.FindAllStringIndex(lines[bellow_line], -1)
                for _, num := range numbers {
                    if num[0] == star_i-1 || num[0] == star_i || num[0] == star_i+1 || num[1]-1 == star_i-1 || num[1]-1 == star_i || num[1]-1 == star_i+1 {
                        nums_spotted = append(nums_spotted, string_number_to_int(string(lines[bellow_line][num[0]:num[1]])))
                    }
                }
            }
            fmt.Println(nums_spotted)
            if 2 <= len(nums_spotted) {
                total += nums_spotted[0]*nums_spotted[1]
            }
        }
    }
    fmt.Println("Total : ", total)
}

func check_digit(value string) bool {
    digit_rgx := regexp.MustCompile(`[\d]`)
    number := digit_rgx.FindString(value)
    if number != "" {
        fmt.Println("Num spotted: ", number)
        return true
    }
    return false
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func string_number_to_int(num string) (num_int int) {
    num_int, err := strconv.Atoi(num)
    check(err)
    return
}
