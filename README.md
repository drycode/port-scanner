# Golang PortScanner

![GitHub top language](https://img.shields.io/github/languages/top/drypycode/port-scanner)
![GitHub](https://img.shields.io/github/license/drypycode/port-scanner)
![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/drypycode/port-scanner/master)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/drypycode/port-scanner)

> This portscanner is an alternative to common portscanners like [nmap](https://nmap.org/). It's use of go routines and channels allows it to scan tens of thousands of arbitrary ports on a remote network in seconds.

> ![alt](https://github.com/drypycode/images/blob/master/portscanner.png)

##

## Installation and Usage

#### Install the binary

```
go get github.com/drypycode/port-scanner
```

#### Clone the Repo

```
git clone github.com/drypycode/port-scanner
...
go run main.go
```

#### Command line arguments

```sh
Usage of ./main:
  -host string
        Hostname or IP address, local or remote. (default "127.0.0.1")
  -portlist string
        A list of specific ports delimited by ','. Can be used w/ or w/o port range. (default ",")
  -portrange string
        A port range, delimited by '-'. 65535 (default "0-3000")
  -protocol string
        Specify the protocol for the scanned ports. (default "TCP")
  -timeout int
        Specify the timeout to wait on a port on the server. (default 5000)
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## License

Distributed under the MIT License. See LICENSE for more information.
