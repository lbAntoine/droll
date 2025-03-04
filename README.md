# dRoll - CLI Dice Roller

![Go Version](https://img.shields.io/badge/Go-1.23-blue)
![Build Status](https://img.shields.io/github/actions/workflow/status/lbAntoine/droll/release.yml)
![License](https://img.shields.io/github/license/lbAntoine/droll)

## 🎲 About dRoll

dRoll is a fast, easy-to-use command-line interface (CLI) dice roller built with Go. Perfect for tabletop RPG players, game masters, and anyone who needs quick and versatile dice rolling functionality.

## ✨ Features

- Roll any number of dice with any number of sides
- Multiple output modes (sum, individual rolls, verbose)
- D&D-style critical success/failure messages for d20
- Flexible command-line arguments
- Simple and intuitive interface

## 🚀 Installation

### Homebrew (macOS/Linux)

Coming soon!

<!-- ```bash -->
<!-- brew tap lbAntoine/droll -->
<!-- brew install droll -->
<!-- ``` -->

### Scoop (Windows)

Coming soon!

<!-- ```powershell -->
<!-- scoop bucket add droll https://github.com/lbAntoine/droll.git -->
<!-- scoop install droll -->
<!-- ``` -->

### Go Install

```bash
go install github.com/lbAntoine/droll@latest
```

### Download Binaries

Download the latest release for your platform from the [Releases](https://github.com/lbAntoine/droll/releases) page.

## 🎯 Usage Examples

### Basic Rolling

```bash
# Roll a single six-sided die
droll 6

# Roll 3 six-sided dice
droll 3 6

# Roll a 20-sided die
droll 20
```

### Advanced Options

```bash
# Roll 2 d20 and show the sum
droll -n 2 -d 20 --sum

# Roll 4 d6 and show individual throws
droll -n 4 -d 6 --unit

# Roll with verbose D&D-style messages
droll -n 1 -d 20 -v
```

### Flags

- `-n, --number`: Number of dice to roll
- `-d, --dice`: Type of dice to roll
- `--sum`: Only show the total sum
- `--unit`: Show individual dice throws
- `-v, --verbose`: Enable D&D-style roll messages

## 🤝 Contributing

### Bug Reports & Feature Requests

1. Check existing [Issues](https://github.com/lbAntoine/droll/issues)
2. Open a new issue with detailed description

### Development Setup

1. Clone the repository

```bash
git clone https://github.com/lbAntoine/droll.git
cd droll
```

2. Install dependencies

```bash
go mod tidy
```

3. Run tests

```bash
go test ./...
```

4. Create a new branch for your feature

```bash
git checkout -b feature/amazing-new-thing
```

5. Make your changes and commit

```bash
git commit -m "Add amazing new feature"
```

6. Push and create a Pull Request

### Contribution Guidelines

- Follow Go coding standards
- Write tests for new functionality
- Update documentation
- Ensure all tests pass before submitting

## 🛠 Building from Source

```bash
go build
```

## 📋 Versioning

We use [Semantic Versioning](https://semver.org/).

## 📜 License

GNU AGPLv3.0

<!-- ## 🆘 Support -->
<!---->
<!-- ### Community Support -->
<!---->
<!-- [INSERT COMMUNITY SUPPORT LINKS] -->
<!---->
<!-- ### Sponsorship -->
<!---->
<!-- [INSERT SPONSORSHIP INFORMATION] -->
<!---->
<!-- ### Donations -->
<!---->
<!-- [INSERT DONATION LINKS] -->

---

**Happy Rolling! 🎲**
