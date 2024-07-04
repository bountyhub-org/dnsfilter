# DNSFilter

DNSFilter is a helper tool that is used to quickly filter out domains found by subdomain discovery tools that are either
invalid, or out of scope. This tool is designed to be used in conjunction with tools like `subfinder`, `assetfinder`, etc.

Invalid domain names are printed to the stderr, while valid domain names are transformed according to IDNA2008 and printed to the stdout.

## Installation

```bash
go install github.com/bountyhub-org/dnsfilter@latest
```

## Usage

### Help

```bash
dnsfilter -h
```

### Example

```bash
subfinder -d example.com | sort -u | dnsfilter -scope .example.com -scope example.net -out-of-scope specific.target.example.com > output.txt
```
