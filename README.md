# goBeyond

**goBeyond** is a command line utility providing information about your computer hardware using **WMI**. It is easy to use, lightweight and user-friendly.

![goBeyond](https://github.com/user-attachments/assets/76dbaad4-9f50-4812-a54d-1b9ea59a8cbe)

## Features

- **Hidden Information about Hardware**: Detect info like the serial number, manufacturer and driver version of your **CPU**, **Memory** and **GPU**.
- **Easy setup**: Easy to setup via **git clone** or **go install**. ([Installation](#installation))
- **User-Experience**: Simple UI & easy to use Syntax.

## Platforms

| Platform       | Developed | Tested |     Version     |
|----------------|:---------:|:------:|:---------------:|
| Windows        |     ✅     |   ✅    | Windows 11 24H2 |
| Linux          |     ❌     |   ❌    |        ❌        |
| macOS          |     ❌     |   ❌    |        ❌        |

## Installation

### Via `go install`

1. **Install the package**:
```bash
go install github.com/hustender/goBeyond@latest
```

2. **Run the program**:
```bash 
goBeyond <component>
```

### Via `git clone`:

1. **Clone the repository**
```bash
git clone https://github.com/hustender/goBeyond.git
cd goBeyond/
```

2. **Install the program**:
```bash
go install
```

3. **Run the program**:
```bash
goBeyond <args>
```

## Available Components

- **CPU**: `cpu`
- **RAM**: `memory`
- **Graphics card**: `gpu`

## Contributing

To contribute, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
