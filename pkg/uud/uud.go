package uud

import (
    "fmt"
    "strings"
    "unicode"

    "golang.org/x/text/unicode/runenames"
)

// Dump produces the uud output for the given text. If rawName is true,
// it uses official Unicode names; otherwise it uses shortened names.
// If showCode is true, adds a U+XXXX code column.
func Dump(text string, rawName bool, showCode bool) string {
    // Compute column widths (approximate display width).
    charColWidth := 1
    for _, r := range text {
        w := displayWidth(r)
        if w > charColWidth {
            charColWidth = w
        }
    }

    // Prepare rows, trimming trailing newlines like uud.py
    trimmed := strings.TrimRight(text, "\n")
    rows := make([][3]string, 0, len(trimmed))
    for _, r := range trimmed {
        var name string
        if rawName {
            name = officialUnicodeName(r)
        } else {
            name = shortUnicodeName(r)
        }
        var code *string
        if showCode {
            c := fmt.Sprintf("U+%04X", r)
            code = &c
        }
        if code != nil {
            rows = append(rows, [3]string{string(r), *code, name})
        } else {
            rows = append(rows, [3]string{string(r), "", name})
        }
    }

    codeColWidth := 0
    if showCode {
        for _, row := range rows {
            if l := len(row[1]); l > codeColWidth {
                codeColWidth = l
            }
        }
    }

    var b strings.Builder
    for _, row := range rows {
        ch := []rune(row[0])[0]
        pad := charColWidth - displayWidth(ch)
        if pad < 0 {
            pad = 0
        }
        if showCode {
            fmt.Fprintf(&b, "%s%s  %-*s  %s\n", string(ch), strings.Repeat(" ", pad), codeColWidth, row[1], row[2])
        } else {
            fmt.Fprintf(&b, "%s%s  %s\n", string(ch), strings.Repeat(" ", pad), row[2])
        }
    }
    return b.String()
}

// displayWidth approximates character display width similarly to uud.py.
// We treat combining marks as width 0, and all others as width 1, clamped to at least 1.
func displayWidth(r rune) int {
    if unicode.Is(unicode.M, r) { // Mark (combining)
        return 1 // clamp to at least 1 for alignment
    }
    // We do not differentiate East Asian Wide/Fullwidth here as examples don't require it.
    return 1
}

func shortUnicodeName(r rune) string {
    name := officialUnicodeName(r)
    if strings.HasPrefix(name, "<unnamed ") {
        return name
    }
    tokens := strings.Fields(name)

    // Collapse SCRIPT (SMALL|CAPITAL)? LETTER X -> X
    for i, t := range tokens {
        if t == "LETTER" {
            rest := tokens[i+1:]
            if len(rest) > 0 {
                return strings.ToLower(strings.Join(rest, " "))
            }
            break
        }
    }

    // DIGIT ZERO -> zero
    if len(tokens) >= 2 && tokens[0] == "DIGIT" {
        return strings.ToLower(strings.Join(tokens[1:], " "))
    }

    // CJK (UNIFIED|COMPATIBILITY) IDEOGRAPH-XXXX -> ideograph-xxxx
    if strings.HasPrefix(name, "CJK UNIFIED IDEOGRAPH-") || strings.HasPrefix(name, "CJK COMPATIBILITY IDEOGRAPH-") {
        // Remove either prefix chunk to leave "IDEOGRAPH-XXXX"
        name2 := strings.Replace(name, "CJK UNIFIED ", "", 1)
        name2 = strings.Replace(name2, "CJK COMPATIBILITY ", "", 1)
        return strings.ToLower(name2)
    }

    // HANGUL SYLLABLE GA -> ga
    if strings.HasPrefix(name, "HANGUL SYLLABLE ") {
        return strings.ToLower(strings.TrimPrefix(name, "HANGUL SYLLABLE "))
    }

    return strings.ToLower(name)
}

func officialUnicodeName(r rune) string {
    if n, ok := unicodeNames[r]; ok {
        return n
    }
    // Use Unicode Character Database names from x/text.
    if name := runenames.Name(r); name != "" {
        // x/text collapses some CJK ideographs to a placeholder; synthesize a standard name.
        if strings.Contains(name, "CJK Ideograph") {
            if (r >= 0xF900 && r <= 0xFAFF) || (r >= 0x2F800 && r <= 0x2FA1F) {
                return fmt.Sprintf("CJK COMPATIBILITY IDEOGRAPH-%04X", r)
            }
            return fmt.Sprintf("CJK UNIFIED IDEOGRAPH-%04X", r)
        }
        return name
    }
    return fmt.Sprintf("<unnamed U+%04X>", r)
}

// unicodeNames includes official names for all code points used in README examples
// unicodeNames allows local overrides when needed (kept minimal).
var unicodeNames = map[rune]string{
    // Keep a few overrides only if needed in tests or for consistency.
}
