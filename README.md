# Stresser

This is a very basic stress testing tool written in go. <br>

Currently, it hits the requested amount of get requests in your domain / path with provided number of goroutines.

# Installation

If you have go installed, you can simple run the main.go file or the whole package.

```bash
git clone https://github.com/regmicmahesh/stresser
cd stresser
go build .
#If you want to install
go install .
```

Now you're ready to run the stresser.

```bash
#Local Build
./stresser
#Installation
stresser
```

Or you can use the docker image provided.

```bash
docker run regmicmahesh/stresser -h
```


# Usage

You can check the help file by executing the given command.

```bash
❯ ./stresser -h
Usage of ./stresser:
  -c int
        Number of hits. (default 100)
  -t int
        Number of threads. (default 4)
  -u string
        URL (default "https://google.com")
```

The help command is self explanatory.

# Contribution

Every minor/major contribution is welcomed and highly appreciated.