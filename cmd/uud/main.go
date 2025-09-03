package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "strings"

    uudpkg "github.com/tlehman/uud/pkg/uud"
)

func main() {
    rawName := flag.Bool("raw-name", false, "Use the full official Unicode name (not shortened)")
    showCode := flag.Bool("show-code", false, "Include a U+XXXX codepoint column")
    flag.Parse()

    var text string
    if flag.NArg() > 0 {
        text = strings.Join(flag.Args(), " ")
    } else {
        b, err := io.ReadAll(os.Stdin)
        if err != nil {
            fmt.Fprintln(os.Stderr, "read stdin:", err)
            os.Exit(1)
        }
        text = string(b)
    }

    out := uudpkg.Dump(text, *rawName, *showCode)
    fmt.Print(out)
}

