
# Elixir In Action

Download and install elixir from [here](https://elixir-lang.org/install.html)

## Building Blocks

### The interactive shell

Use this to experiment the new language feature.

To start a shell, by running `iex` from the command line

This starts the BEAM and then starts an interactive Elixir shell insite of it.

Everything in Elixir is an expression that has a return value. Function calls and if conditionals as well.

Elixir uses line breaks to indicate the end of a complete expression.

Use System.halt to gracefully end the shell, or use Ctrl+-C twice to brutally exist it.

To get help, Do:
```
$ iex
iex(1) > h
```

### Working with variabales

As a dynamic programming language, Eixir does not support type, but uses the current data type of a variable to determine its type.

Elixir uses binding instead of assignment, which means, when a variable is initialized witha variable, the variable is bound to that value.

Every expressin in Elixir has a result. 

After an expression is evaluated the shell prints the result of the expression to the screen.

```
iex(1)> total_attendance = 1234
1234
iex(2)> total_attendance
1234
```

variable names can have the below format:

`variable_name`

`variableName`

`VariableName`

A variable can be rebound to, which does not mutate the existing memory location.

Rebound reserves new memory and reassigns the name to the new location.
```
iex(1)> total_attendance = 1234
1234
iex(2)> total_attendance = 12345
12345
iex(3)> total_attendance
12345
```

Code organization
