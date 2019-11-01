# ShowCFDI
Displays information from CFDI(xml) inside the folder, shows the date, EmisorRFC, Receiver RFC, subtotal, tax and Total.

## Why?
I don't like to open each time the pdf in order to know the expenses for that declaration.
In order to compare the local CFDIs with the web CFDIs (located in the Sat web site), its easier to check the listed RFC and total for each CFDI with the remote CFDIs table.

## Install
1. Clone the repo.
    `git clone 
2. build it
    go buld main.go -o bin/cfdi
3. install
    we have two options if your GOBIN path is correct just type: `go install` other wise just move the binary file to your system path 
    osx: `mv bin/cfdi /usr/local/bin`.
    linux: `mv bin/cfdi /bin`.
    4. optional run the bash