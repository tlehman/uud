# `uud(1)` - unicode dump

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`uud` is a Unicode dumping tool inspired by `xxd(1)`, the hex dump tool

## Example usage:

Unpacking unicode math characters:
```
echo '𝑓(𝑥) = 3𝑥³ − 2𝑥² + 5𝑥 − 7' | uud
𝑓  mathematical italic small f
(  left parenthesis
𝑥  mathematical italic small x
)  right parenthesis
   space
=  equals sign
   space
3  three
𝑥  mathematical italic small x
³  superscript three
   space
−  minus sign
   space
2  two
𝑥  mathematical italic small x
²  superscript two
   space
+  plus sign
   space
5  five
𝑥  mathematical italic small x
   space
−  minus sign
   space
7  seven
```

Identifying Spanish accents:
```
$ echo 'creó' | uud
c  c
r  r
e  e
ó  o with acute
```

Identifying Hebrew characters:
```
$ echo 'וירא' | uud
ו  vav
י  yod
ר  resh
א  alef
```

Viewing the niqqud on Hebrew characters:

```
$ echo 'בְּרֵאשִׁית, בָּרָא אֱלֹהִים' | uud
ב  bet
ְ  hebrew point sheva
ּ  hebrew point dagesh or mapiq
ר  resh
ֵ  hebrew point tsere
א  alef
ש  shin
ִ  hebrew point hiriq
ׁ  hebrew point shin dot
י  yod
ת  tav
,  comma
   space
ב  bet
ָ  hebrew point qamats
ּ  hebrew point dagesh or mapiq
ר  resh
ָ  hebrew point qamats
א  alef
   space
א  alef
ֱ  hebrew point hataf segol
ל  lamed
ֹ  hebrew point holam
ה  he
ִ  hebrew point hiriq
י  yod
ם  final mem
```

Chinese characters:
```
$ echo '太陽能板張開' | ./bin/uud
太  ideograph-592a
陽  ideograph-967d
能  ideograph-80fd
板  ideograph-677f
張  ideograph-5f35
開  ideograph-958b
```

## License

MIT © 2025 Tobi Lehman <mail@tobilehman.com>. See `LICENSE` for details.
