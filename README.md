# ShowCFDI
Displays information from CFDI(xml) inside the folder, shows the date, EmisorRFC, Receiver RFC, subtotal, tax, and Total.

## why?
I don't like to open each time the pdf to know the expenses for that declaration.
To compare the local CFDIs with the web CFDIs (located on the Sat website), it's easier to check the listed RFC and total for each CFDI with the remote CFDIs table.

## Install
1. Clone the repo.
 `git clone https://github.com/osadsanu/checkCFDI.git`
2. build it
 go build main.go -o bin/cfdi
3. install
 we have two options if your GOBIN path is correct just type: `go install` otherwise just move the binary file to your system path osx: `mv bin/cfdi /usr/local/bin`.
 Linux: `mv bin/cfdi /bin`.
4. Optional run the bash file to automatically clear the binaries, compile the program, inside the folder bin:
 `bash install .sh`