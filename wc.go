package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {

    content_string string, discard int = ioutil.ReadFile(os.Args[1])
    char_count int = len(content_string)
    words_string string = strings.Fields(content_string)
    words_count int = len(words_string)
    lookForRune(content_string, "\n")

    return fmt.Sprintf("%s\n", words_count)

}

func lookForRune(str string, rn rune) (num int) {

    occurrences := 0

    for _, char := range str {

        if char == rune {
            occurrences++
        }

    }

    return occurrences
}
