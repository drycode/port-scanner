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

```
Usage of ./main:
  -hosts string
        A list DNS names or IP addresses (local or remote) delimited by ','. Additionally, for IP addresses
        the user can provide a valid CIDR notation block, and the range of IPs defined in that block will
        be scanned.

        Ex. "127.0.0.1, www.google.com, 192.0.0.0/24, 100.0.0.0-100.0.1.0"

        WARNING: A large range of IP addresses compounds exponentially against the list of ports to scan. 10 hosts @ 10k ports == 100k total scans

        (default "127.0.0.1")

  -ports string
        A list of specific ports delimited by ','. Optionally: A range of ports can be provided in addition to to comma delimited specific ports.

        Ex. "80, 443, 100-200, 6543"

  -protocol string
        Specify the protocol for the scanned ports. (default "TCP")

  -timeout int
        Specify the timeout to wait on a port on the server. (default 5000)
```

#### Example Usage

```
➜  go run main.go --ports="80,4423,100-105,40-45" --hosts='127.0.0.1,google.com,facebook.com'
Starting Golang GoScan v0.1.0 ( github.com/drypycode/portscanner/v0.1.0 ) at Wed, 28 Apr 2021 19:35:45 EDT

Scanning ports on  [127.0.0.1 facebook.com google.com]
[#####################################################################################################]  100%
GoScan done: 12 ports scanned in 10 seconds.

Open Ports
facebook.com:80
google.com:80
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
