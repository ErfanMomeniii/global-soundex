# global-soundex

`global-soundex` is a Go package that provides Soundex phonetic encoding for global languages. Soundex is a phonetic algorithm used to index words by their sound when pronounced. This package is ideal for approximate string matching, name matching, and spelling correction across different languages.

## ✨ Features

- 🗣️ **Arabic and English** Soundex support
- 🔄 Normalize inputs (diacritics, punctuation, and casing)
- 🔠 Language-specific encoding rules
- 🧠 Intelligent suggestion using Hamming and levenshtein distance

## 📦 Installation

```bash
go get -u github.com/erfanmomeniii/global-soundex
```

## 🚀 Usage
### Arabic Example
```go
package main

import (
	"fmt"
	gs "github.com/erfanmomeniii/global-soundex"
)

func main() {
	ar := gs.New("ar")
	ar.AddEntities([]string{"محمد", "محمود", "عبدالله", "أحمد"})

	fmt.Println("Code:", ar.Encode("احمد")) // output: a530
	fmt.Println("Match:", ar.Correspond("احمد")) // output: احمد
	fmt.Println("Suggest:", ar.Suggest("احمذ")) // output: احمد
}
```

### English Example
```go
package main

import (
	"fmt"
	gs "github.com/erfanmomeniii/global-soundex"
)

func main() {
	en := gs.New("en")
	en.AddEntities([]string{"Robert", "Rubin", "Ashcraft"})

	fmt.Println("Code:", en.Encode("Robert")) // output: r163
	fmt.Println("Match:", en.Correspond("Rupert")) // output: Rupert
	fmt.Println("Suggest:", en.Suggest("Ribert")) // output: Robert
}
```

## 🔍 Difference Between `Correspond` and `Suggest`

| Method       | Description                                                                                                                                                                                                             |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Correspond` | Returns a known entity if the Soundex code **exactly matches** an entry in the added list of entities. If no match is found, it returns the original input.                                                             |
| `Suggest`    | Returns the **closest matching entity** based on **Hamming or Levenshtein distance** between the Soundex codes of the input and the known entities. Useful when the input is slightly misspelled or phonetically close. |

### 📘 Example

```go
en := globalsoundex.New("en")
en.AddEntities([]string{"Robert", "Rupert", "Rubin"})

fmt.Println(en.Correspond("Rupert")) // Output: Rupert (exact match)
fmt.Println(en.Correspond("Ribert")) // Output: Ribert (no exact match)
fmt.Println(en.Correspond("Ruperz")) // Output: Rupert (same encoding as "Rupert")

fmt.Println(en.Suggest("Ribert"))    // Output: Robert (closest match based on Soundex code)
```

## Contributing

Pull requests are welcome! For any changes, please open an issue first to discuss the proposed modification. Ensure tests are updated accordingly.
