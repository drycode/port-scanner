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
        A list DNS names or IP addresses (local or remote) delimited by ','. Additionally, for IP addresses the user can provide a valid 
        CIDR notation block, and the range of IPs defined in that block will be scanned. 
        Ex. '127.0.0.1, www.google.com, 192.0.0.0/24, 100.0.0.0-100.0.1.0' 
    
        WARNING: A large range of IP addresses compounds exponentially against the list of ports to scan. 
        10 hosts @ 10k ports == 100k total scans
         (default "127.0.0.1")
  -jump
        [Optional] Allows you to build and run the portscanner on a remote machine.
         Currently supports OS: Linux | Architecture: ARM64 
  -output string
        [Optional] Output filepath to which open ports will be written. Included filepath will determine output type.
        Supported file types: .json, .txt
  -ports string
        A list of specific ports delimited by ','. Optionally: A range of ports can be provided in addition to to comma delimited 
        specific ports.
        Ex. '80, 443, 100-200, 6543'
  -protocol string
        Specify the protocol for the scanned ports. (default "TCP")
  -remote-host string
        [Optional] Remote host IP address or DNS to use as a jump box. Useful for assessing the open ports 
        secured behind a firewall. (requires --jump)
  -remote-user string
        [Optional] Login username for the remote machine. (requires --jump)
  -ssh-key string
        [Optional] Path to the ssh key used to connect to the remote container (requires --jump)
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

##### With Jump
```
➜  go run main.go --ports=80,4423,100-105,40-45,7000-9000 --hosts='10.0.0.0' --protocol=TCP --remote-host=ec2-machine.compute-1.amazonaws.com --remote-user='ec2-user' --ssh-key='/Users/ssh/key.pem' --jump

Starting Golang GoScan v0.1.0 ( github.com/drypycode/portscanner/v0.1.0 ) at Mon, 10 May 2021 22:16:56 UTC

ec2-machine.compute-1.amazonaws.com is scanning ports on  [target-ec2-dns.compute-1.amazonaws.com]
[####################################################################################################]  100%
GoScan done: 2012 ports scanned in 10.04 seconds. 

Open Ports
...
```

##### With Output Flag
```
➜  go run main.go --ports="80,4423,100-105,40-45, 1000-20000" --hosts='127.0.0.1,localhost,google.com' --output="/tmp/dat2.json"
Starting Golang GoScan v0.1.0 ( github.com/drypycode/portscanner/v0.1.0 ) at Mon, 10 May 2021 18:19:55 EDT

Scanning ports on  [127.0.0.1 google.com localhost]
[####################################################################################################]  100%
GoScan done: 30036 ports scanned in 12.02 seconds. 

➜  cat /tmp/dat2.json 
[
  {
    "Host": "127.0.0.1",
    "Ports": [
      "8000",
      "5000"
    ]
  },
  {
    "Host": "google.com",
    "Ports": [
      "80"
    ]
  },
  {
    "Host": "localhost",
    "Ports": [
      "8000",
      "5000"
    ]
  }
]
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
