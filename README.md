# thesaurus

[![go-ci](https://github.com/R1nY1x1/thesaurus/actions/workflows/ci.yml/badge.svg)](https://github.com/R1nY1x1/thesaurus/actions/workflows/ci.yml)

Thesaurus by Golang

## Usage

```
>> ./main hello
Input: hello
[hi how-do-you-do howdy hullo] an expression of greeting

>> ./main -xin ありがとう
Input: sorry < ごめんなさい >
[bad regretful] feeling or expressing regret or sorrow or a sense of loss over something done or undone
[good-for-naught good-for-nothing meritless no-account no-count no-good] without merit; of little or no value or use
[blue dark dingy disconsolate dismal drab drear dreary gloomy grim] causing dejection
[deplorable distressing lamentable pitiful sad] bad; unfortunate
```

```
>> ./main -h
Usage of ./main:
  -s string
    	Translate souece (default "ja")
  -t string
    	Translate target (default "en")
  -xin
    	do Translate Input
  -xout
    	do Translate Output
```

## License

MIT

## Author

R1nY1x1
