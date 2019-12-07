RepoSAM
=======

A self-configuring Debian repository for I2P based on
[repogen](https://github.com/geek1011/repogen) and sam-forwarder. Works as an
alternative launcher for a repogen apt server that automatically forwards it to
I2P.

Compile/Installation
====================

[repogen](https://github.com/geek1011/repogen) is an apt repository management
tool with a web interface.

With your GOPATH set:

        git clone https://github.com/eydeekay/reposam $GOPATH/src/github.com/eyedeekay/reposam
        cd $GOPATH/src/github.com/eyedeekay/reposam
        make
        sudo make install


        Usage of ./reposam/reposam:
         -a string
              hostname to serve on (default "127.0.0.1")
         -c	Use an encrypted leaseset(true or false)
         -de string
              sets the description field used in the Release file (default "generated by reposam")
         -f string
              Use an ini file for configuration (default "none")
         -g	Uze gzip(true or false) (default true)
         -gc
              generates the Contents index
         -gw
              generate a web interface for browsing the packages
         -i	save i2p keys(and thus destinations) across reboots (default true)
         -ib int
              Set inbound tunnel backup quantity(0 to 5) (default 1)
         -id string
              directory to look for packages (default "not_a_dir")
         -il int
              Set inbound tunnel length(0 to 7) (default 3)
         -iq int
              Set inbound tunnel quantity(0 to 15) (default 2)
         -iv int
              Set inbound tunnel length variance(-7 to 7)
         -key string
              private key for signing repository packages (default "invalid_sk.asc")
         -l string
              Type of access list to use, can be "whitelist" "blacklist" or "none". (default "none")
         -m string
              Certificate name to use (default "cert")
         -mo string
              overrides the maintainer of all packages
         -n string
              name to give the tunnel(default reposam) (default "reposam")
         -ob int
              Set outbound tunnel backup quantity(0 to 5) (default 1)
         -od string
              directory to output the repository (default "not_a_dir")
         -ol int
              Set outbound tunnel length(0 to 7) (default 3)
         -oq int
              Set outbound tunnel quantity(0 to 15) (default 2)
         -or string
              sets the origin field used in the Release file (default "reposam")
         -ov int
              Set outbound tunnel length variance(-7 to 7)
         -p string
              port to serve locally on (default "7880")
         -r	Reduce tunnel quantity when idle(true or false)
         -rc int
              Reduce idle tunnel quantity to X (0 to 5) (default 3)
         -rt int
              Reduce tunnel quantity after X (milliseconds) (default 600000)
         -sh string
              sam host to connect to (default "127.0.0.1")
         -sp string
              sam port to connect to (default "7656")
         -sv
              show the version
         -t	Generate or use an existing TLS certificate
         -wa
              watch the input directory for new packages
         -wi int
              the interval to check for new packages (default 1)
         -z	Allow zero-hop, non-anonymous tunnels(true or false)
