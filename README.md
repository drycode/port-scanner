[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# Port Scanner

> This program is intended to be a more efficient version of other popular port scanners like [nmap](https://nmap.org/).

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [About the Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgements](#acknowledgements)

<!-- ABOUT THE PROJECT -->

## About The Project

### Built With

- [Go](https://golang.org/)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- [Go](https://golang.org/dl/)

### Installation

```sh
go get -u github.com/drypycode/port-scanner
```

<!-- USAGE EXAMPLES -->

## Usage

```sh
Usage
  -host string
        Hostname or IP address, local or remote. (default "127.0.0.1")
  -portlist string
        A list of specific ports delimited by ','. Can be used w/ or w/o port range.
  -portrange string
        A port range, delimited by '-'. 65535 (default "0-3000")
  -protocol string
        Specify the protocol for the scanned ports. (default "TCP")
  -timeout int
        Specify the timeout to wait on a port on the server. (default 5000)
```

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Your Name - daniel.richard.young@gmail.com

Project Link: [https://github.com/drypycode/port-scanner](https://github.com/drypycode/port-scanner)

<!-- ACKNOWLEDGEMENTS -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/drypycode/repo.svg?style=flat-square
[contributors-url]: https://github.com/drypycode/repo/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/drypycode/repo.svg?style=flat-square
[forks-url]: https://github.com/drypycode/repo/network/members
[stars-shield]: https://img.shields.io/github/stars/drypycode/repo.svg?style=flat-square
[stars-url]: https://github.com/drypycode/repo/stargazers
[issues-shield]: https://img.shields.io/github/issues/drypycode/repo.svg?style=flat-square
[issues-url]: https://github.com/drypycode/repo/issues
[license-shield]: https://img.shields.io/github/license/drypycode/repo.svg?style=flat-square
[license-url]: https://github.com/drypycode/repo/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/youngdanielr
[product-screenshot]: images/screenshot.png
