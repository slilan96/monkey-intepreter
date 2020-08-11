# The first part - Lexical Analysis

We need to convert the source code into something that the machine can
understand. However as the amount of code grows large, it becomes cumbersome
to keep track of everything that has happened. Therefore, it's imperative that
we find an **easier** way to keep track of this stuff.

Therefore the source code has to undergo two transformations:

  1. From source code to tokens.(lexing / lexical analysis).
  2. From tokens to Abstract Syntax Trees(more on this later).

Step 1 is done by a *lexer* or *scanner* or *tokenizer*. The last two
definitions can be used interchangeably, but there are some **nuances** so
watch out for this one.

## Tokens

What are Tokens? They are small and easily categorizable data structures.

## Defining Tokens


