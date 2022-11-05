# CyberSec-subDomainEnumeration

The project is used to scan the subdomains of the given websites and enumerate all the existing
used or unused subdomains. The host can secure all the subdomains that are insecure. Passive
reconnaissance is an essential part of cyber security, and subdomain enumeration plays a major
role in web penetration testing. Therefore, in this project, we aim to build a tool that does both
subdomain enumeration as well as directory enumeration efficiently and in a fast manner. For
this we can use open source repositories like ‘crt.sh’ and brute-force. We plan to build a CLI tool
using Golang.

```
Compilation

go build -o ./bin/enumerator

```


```
Usage of enumerator:

    -s value
        Status code for success, default 200
    -u value
        Target URL to be bruteforced [required]
    -w string
        Path to wordlist [required]
    -c value
        Number of cores, default 1 (1 to 8)
    -t value
        Type of enumeration, default 0 (0 -> sub-domain, 1 -> directory)
```

```
Example:

    ./bin/enumerator -u google.com -s 200 -c 8 -t 1 -w ./wordlist/common.txt
```