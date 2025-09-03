package uud

import "testing"

func TestDump_MathExpression(t *testing.T) {
    in := "𝑓(𝑥) = 3𝑥³ − 2𝑥² + 5𝑥 − 7\n"
    want := "" +
        "𝑓  mathematical italic small f\n" +
        "(  left parenthesis\n" +
        "𝑥  mathematical italic small x\n" +
        ")  right parenthesis\n" +
        "   space\n" +
        "=  equals sign\n" +
        "   space\n" +
        "3  three\n" +
        "𝑥  mathematical italic small x\n" +
        "³  superscript three\n" +
        "   space\n" +
        "−  minus sign\n" +
        "   space\n" +
        "2  two\n" +
        "𝑥  mathematical italic small x\n" +
        "²  superscript two\n" +
        "   space\n" +
        "+  plus sign\n" +
        "   space\n" +
        "5  five\n" +
        "𝑥  mathematical italic small x\n" +
        "   space\n" +
        "−  minus sign\n" +
        "   space\n" +
        "7  seven\n"

    got := Dump(in, false, false)
    if got != want {
        t.Fatalf("unexpected output.\n--- got ---\n%q\n--- want ---\n%q\n", got, want)
    }
}

func TestDump_SpanishAccents(t *testing.T) {
    in := "creó\n"
    want := "" +
        "c  c\n" +
        "r  r\n" +
        "e  e\n" +
        "ó  o with acute\n"
    got := Dump(in, false, false)
    if got != want {
        t.Fatalf("unexpected output.\n--- got ---\n%q\n--- want ---\n%q\n", got, want)
    }
}

func TestDump_HebrewAndNiqqud(t *testing.T) {
    in := "בְּרֵאשִׁית, בָּרָא אֱלֹהִים\n"
    want := "" +
        "ב  bet\n" +
        "ְ  hebrew point sheva\n" +
        "ּ  hebrew point dagesh or mapiq\n" +
        "ר  resh\n" +
        "ֵ  hebrew point tsere\n" +
        "א  alef\n" +
        "ש  shin\n" +
        "ִ  hebrew point hiriq\n" +
        "ׁ  hebrew point shin dot\n" +
        "י  yod\n" +
        "ת  tav\n" +
        ",  comma\n" +
        "   space\n" +
        "ב  bet\n" +
        "ָ  hebrew point qamats\n" +
        "ּ  hebrew point dagesh or mapiq\n" +
        "ר  resh\n" +
        "ָ  hebrew point qamats\n" +
        "א  alef\n" +
        "   space\n" +
        "א  alef\n" +
        "ֱ  hebrew point hataf segol\n" +
        "ל  lamed\n" +
        "ֹ  hebrew point holam\n" +
        "ה  he\n" +
        "ִ  hebrew point hiriq\n" +
        "י  yod\n" +
        "ם  final mem\n"

    got := Dump(in, false, false)
    if got != want {
        t.Fatalf("unexpected output.\n--- got ---\n%q\n--- want ---\n%q\n", got, want)
    }
}

