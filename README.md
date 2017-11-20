# hoc
Higher Order Calculator implemented in go language

## Installation

```
go get -u github.com/surabhig412/hoc
go get -u golang.org/x/tools/cmd/goyacc
```

## Run the code

```bash
$ make clean
$ make
$ ./hoc
```

## Operations

#### Mathematical Operations

```
$ 2 + 3
$ 2 * 3
$ 15 / 3
$ 20 % 3
$ -2 - 3
$ 2 ^ 3
$ sin(90)
$ cos(90)
$ atan(45)
$ log(10)
$ logten(10)
$ exp(10)
$ sqrt(4)
$ abs(-90)
```

#### Relational Operations

```
>, >=, <, <=, ==, !=
Example: 5 > 7
Program returns 0(false) or 1(true)
```

#### Loops

###### While loop

```
while (3<5) 4+5
```

###### If loop

```
if (3<5) 4+5 else 6+7
```
