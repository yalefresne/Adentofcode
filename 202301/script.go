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

func word_to_int(line string) (string_with_int string) {
    string_with_int = strings.Replace(strings.ToLower(line), "nine", "n9e", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "eight", "e8t", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "seven", "s7n", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "six", "s6x", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "five", "f5e", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "four", "f4r", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "three", "t3e", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "two", "t2o", -1)
    string_with_int = strings.Replace(strings.ToLower(string_with_int), "one", "o1e", -1)
    return
}

func main() {
    //re := regexp.MustCompile("[0-9]+")
    // last := array_numbers[len(array_numbers)-1]
    //inputs_from_calendar, err := os.ReadFile("./inputs.txt")
    inputs_from_calendar, err := os.Open("./inputs.txt")
    check(err)

    defer inputs_from_calendar.Close()

    scanner := bufio.NewScanner(inputs_from_calendar)
    var total int
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
        line = word_to_int(line)
        fmt.Println(line)
        re := regexp.MustCompile("[0-9]+")
        numbers_in_line := re.FindAllString(line, -1)
        first := numbers_in_line[0]
        last := numbers_in_line[len(numbers_in_line)-1]
        //fmt.Println(line)
        first = first[:1]
        last = last[len(last)-1:]
        fmt.Println("first:", first)
        fmt.Println("last:", last)
        fmt.Println("together:", first+last)
        line_int, err := strconv.Atoi(first+last)
        check(err)
        total += line_int
    }

    fmt.Println("total:", total)
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
