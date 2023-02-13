# domains

![](https://i.imgur.com/AxOdYRA.png)

`domains` is a command-line tool for extracting domains from a list of input strings and outputting the unique domains found. It can read input either from a file or from standard input.

## Installation

To install `domains`, you must have Go installed on your system. Then, run the following command:

```
go install github.com/Ademking/domains@latest
```

This will download the package and its dependencies and build the domains binary, which you can run from the command line.

Alternatively, you can download the source code from this repository and build it manually:

```
git clone https://github.com/Ademking/domains
cd domains
go build
```

## Usage

To use domains, you can pipe input to it from another command:

```
cat input.txt | domains
```

This will read the contents of `input.txt`, extract any domains found, and output a list of unique domains to standard output.

By default, domains does not add any prefix or suffix to the extracted domains. However, you can use the `--prefix` and `--suffix` flags to add a string to the beginning or end of each domain:

```
cat input.txt | domains --prefix "https://" --suffix "/api"
```

This will add `https://` to the beginning of each domain and `/api` to the end of each domain.

You can also specify an input file using the --input flag:

```
domains --input input.txt
```

## Contributing

If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please open a pull request.

## License

`domains` is licensed under the MIT License. See the LICENSE file for more details.