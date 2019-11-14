RepoSAM
=======

A self-configuring Debian repository for I2P based on
[cowyo](https://github.com/geek1011/repogen) and sam-forwarder. Works as an
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