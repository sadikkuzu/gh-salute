# gh-salute

A GitHub CLI extension that greets you with your GitHub profile information.

## Description

`gh-salute` is a simple GitHub CLI extension written in Go that connects to the GitHub API and displays your GitHub username and profile URL. It uses the [go-gh](https://github.com/cli/go-gh) library to interact with the GitHub API.

## Installation

```bash
gh extension install sadikkuzu/gh-salute
```

## Usage

```bash
gh salute
```

Example output:
```
aloha world, this is the gh-salute extension!
running as yourusername
profile: https://github.com/yourusername
```

## Requirements

- [GitHub CLI](https://cli.github.com/) (`gh`)
- Go 1.24.3 or higher

## Development

### Clone the repository

```bash
git clone https://github.com/sadikkuzu/gh-salute.git
cd gh-salute
```

### Build

```bash
go build
```

### Install locally for development

```bash
gh extension install .
```

## License

This project is open source. See the [LICENSE](LICENSE) file for details.

## Author

- [Sadık Kuzu](https://github.com/sadikkuzu)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
