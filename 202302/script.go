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
    total_ids := 0

    for scanner.Scan() {
        game := scanner.Text()

        max_blue := count_cube_by_color("blue", game)
        max_green := count_cube_by_color("green", game)
        max_red := count_cube_by_color("red", game)

        subtotal := max_red * max_green * max_blue
        total_ids += subtotal
    }
    fmt.Println("Total : ", total_ids)
}

func count_cube_by_color(color string, game string) (max_kube int) {
    re := regexp.MustCompile("[0-9]+\\s" + color)
    count := re.FindAllStringSubmatch(game, -1)

    max_kube = 0

    if len(count) > 0 {
        for _, color_kube := range count {
            kube := strings.Replace(color_kube[0], " " + color, "", 1)
            count_kube := string_number_to_int(kube)

            if max_kube < count_kube {
                max_kube = count_kube
            }
        }
    }
    return
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
