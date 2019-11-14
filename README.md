CowYoSam
========

A zero-configuration wiki for I2P based on
[cowyo](https://github.com/schollz/cowyo) and sam-forwarder. Works as an
alternative launcher for a cowyo wiki that automatically forwards it to I2P.

Or something I should probably have not done between the hours of noon and 1PM
on Wednesday, November 6th.

Compile/Installation
====================

[cowyo](https://github.com/schollz/cowyo) is sort of like what you'd get if you crossed a wiki and a pastebin.
Anyone can edit, anonymously, in the default configuration. For now, I have
disabled file uploads and some some other features are non-configurable. I will
add them back in soon.

With your GOPATH set:

        git clone https://github.com/eydeekay/cowyosam $GOPATH/src/github.com/eyedeekay/cowyosam
        cd $GOPATH/src/github.com/eyedeekay/cowyosam
        make
        sudo make install