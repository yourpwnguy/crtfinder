<div align="center">

![GoStyle LOGO](https://i.imgur.com/69aPTml.png)

</div>
<h4 align="center">Fast and efficient subdomain enumeration tool utilizing crt.sh to identify subdomains recursively.</h4>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/iaakanshff/crtfinder">
<!-- <a href="https://github.com/iaakanshff/crtfinder/releases"><img src="https://img.shields.io/github/downloads/iaakanshff/crtfinder/total"> -->
<a href="https://github.com/iaakanshff/crtfinder/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/iaakanshff/crtfinder">
<!-- <a href="https://github.com/iaakanshff/crtfinder/releases/"><img src="https://img.shields.io/github/release/iaakanshff/crtfinder"> -->
<a href="https://github.com/iaakanshff/crtfinder/issues"><img src="https://img.shields.io/github/issues-raw/iaakanshff/crtfinder">
<a href="https://github.com/iaakanshff/crtfinder/stars"><img src="https://img.shields.io/github/stars/iaakanshff/crtfinder">
<!-- <a href="https://github.com/iaakanshff/crtfinder/discussions"><img src="https://img.shields.io/github/discussions/iaakanshff/crtfinder"> -->
</p>

---

## Features 💡

- Fast subdomain enumeration using crt.sh.
- Recursive subdomain discovery.
- Supports input via command-line flags or input files.

## Installation 🛠️ 

To install the crtfinder tool, you can simply run the following command.

```bash
go install -v "github.com/yourpwnguy/crtfinder/cmd/crtfinder@latest"
cp ~/go/bin/crtfinder /usr/local/bin/
```

## Usage 📝

```yaml
Usage: crtfinder [options]

Options: [flag] [argument] [Description]

INPUT:
  -d string[]   Domains to find subdomains for (comma separated)
  -dL FILE      Input file containing a list of domains

FEATURES:
  -r int        For recursively finding subdomains with time gap between requests (default: 5s)

OUTPUT:
  -o string     Output file to store the subdomains

DEBUG:
  -v none       Check current version
```

## But Why Use Our Tool❓ 

We understand and appreciate there are other tools out there for the same task, but many aren't updated or maintained. Crtfinder gives you more control, speed, and a better user interface. It's regularly updated to ensure reliable performance and modern features for easy subdomain discovery.

## Contributions 🤝

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request.
