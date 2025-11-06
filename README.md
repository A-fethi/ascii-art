# 🎨 ASCII-Art Generator

**ASCII-Art** is a Go program that takes a string as input and outputs a **graphical representation** of that string using ASCII characters.  
The program handles letters, numbers, spaces, special characters, and line breaks (`\n`) to create visually appealing text banners.

---

## 🚀 Features

- Converts input strings into **ASCII art** using different banner styles:
  - `standard`
  - `shadow`
  - `thinkertoy`
- Supports **multi-line input** with `\n`.
- Handles **letters, numbers, spaces, and special characters**.
- Maintains proper alignment for complex strings and line breaks.

---

## ⚙️ Usage

1. **Run the program with a string argument**:

```bash
go run . "Hello World"
```

2. **Run the program with line breaks:**

```bash
go run . "Hello\nThere"
```

3. **View output:**
```bash
cat result.txt
```

Example:

```txt
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

## 🎨 Banner Format

- Each character has a height of 8 lines.

- Characters are separated by a newline (\n).

- Supports letters, numbers, special characters, and spaces.

- Banner files are preformatted and should not be modified.

## 🧰 Technologies Used

- Language: Go (Golang)

- Packages: Standard library only.

## 🎯 Learning Outcomes

- Working with file input/output using Go’s fs API

- String and data manipulation
 
- Handling multi-line input and text formatting
 
- Implementing ASCII-based visual representation

## 👤 Author
**Abderrahmane Fethi**

Junior Full-Stack Developer | Passionate about clean code and problem-solving