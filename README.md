# global-soundex
# global-soundex

`global-soundex` is a Go package that provides Soundex phonetic encoding for global languages. Soundex is a phonetic algorithm used to index words by their sound when pronounced. This package is ideal for approximate string matching, name matching, and spelling correction across different languages.

---

## ✨ Features

- 🗣️ **Arabic and English** Soundex support
- 🔄 Normalize inputs (diacritics, punctuation, and casing)
- 🔠 Language-specific encoding rules
- 🧠 Intelligent suggestion using Hamming and levenshtein distance
- 📦 Clean, extensible architecture

---

## 📦 Installation

```bash
go get github.com/erfanmomeniii/global-soundex
```

## 🚀 Usage
Arabic Example
```go
package main

import (
	"fmt"
	gs "github.com/erfanmomeniii/global-soundex"
)

func main() {
	ar := gs.NewArabic()
	ar.AddEntities([]string{"محمد", "محمود", "عبدالله", "أحمد"})

	fmt.Println("Code:", ar.Encode("احمد"))
	fmt.Println("Match:", ar.Correspond("احمد"))
	fmt.Println("Suggest:", ar.Suggest("احمذ"))
}
```

English Example
```go
package main

import (
	"fmt"
	gs "github.com/erfanmomeniii/global-soundex"
)

func main() {
	en := gs.NewEnglish()
	en.AddEntities([]string{"Robert", "Rupert", "Rubin", "Ashcraft"})

	fmt.Println("Code:", en.Encode("Robert"))
	fmt.Println("Match:", en.Correspond("Rupert"))
	fmt.Println("Suggest:", en.Suggest("Ribert"))
}
```

## Contributing

Pull requests are welcome! For any changes, please open an issue first to discuss the proposed modification. Ensure tests are updated accordingly.
