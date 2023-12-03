package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    inputs_from_calendar, err := os.Open("./inputs.txt")
    check(err)

    defer inputs_from_calendar.Close()

    scanner := bufio.NewScanner(inputs_from_calendar)
    lines := make([]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
        // Index all characters in line
        //line_sliced := strings.SplitAfter(line, "")
        // Index line by line
        lines = append(lines, line)
    }

    total := 0

    for i_l, line := range lines {
        digit_rgx := regexp.MustCompile(`[\d]+`)
        numbers := digit_rgx.FindAllStringIndex(line, -1)
        for i := range numbers {
            number := string_number_to_int(string(line[numbers[i][0]:numbers[i][1]]))
            min_i := 0
            max_i := 1
            // same line
            if 1 <= numbers[i][0] {
                min_i = numbers[i][0]-1
            } else {
                min_i = numbers[i][0]
            }
            if len(lines[i_l]) > numbers[i][1] {
                max_i = numbers[i][1]+1
            } else {
                max_i = numbers[i][1]
            }
            // line above
            if i_l > 0 {
                if check_symbol(string(lines[i_l-1][min_i:max_i])) == true {
                    total += number
                    continue
                }
            }
            if check_symbol(string(lines[i_l][min_i:max_i])) == true {
                total += number
                continue
            }
            if len(lines) > i_l+1 {
                if check_symbol(string(lines[i_l+1][min_i:max_i])) == true {
                    total += number
                    continue
                }
            }
        }
    }
    fmt.Println("Total : ", total)
}

func check_symbol(value string) bool {
    fmt.Println(value)
    not_digit_rgx := regexp.MustCompile(`[\!\@\#\$\%\^\&\*\_\-\+\=\?\>\<\\\/]`)
    symbol := not_digit_rgx.FindString(value)
    if symbol != "" && symbol != "." {
        fmt.Println("Symbol spotted: ", symbol)
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
