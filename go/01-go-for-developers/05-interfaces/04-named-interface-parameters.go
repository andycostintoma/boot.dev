package main

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
