Introduction
============

The gousb package is an attempt at wrapping the libusb library into a Go-like binding.

Supported platforms include:

- ~~linux~~
- ~~darwin~~
- windows


Installation
============
Installing gousb package is really easy:
---------------------------------------------------

    go get -v github.com/google/gousb


Notes for installation on Windows
---------------------------------

You'll need:

- Gcc - tested on [Win-Builds](http://win-builds.org/) and MSYS/MINGW
- libusb-1.0 - included as static library (windows only)

