![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)
![](https://travis-ci.org/goccmack/gogll.svg?branch=master)

Copyright 2019 Marius Ackerman. 

# Gogll
Gogll generates a matching lexer and parser from one grammar in Go. 

The lexer is a linear-time finite state automaton FSA [[Grune et al 2012](#Grune-et-al-2012)].

The parser is a clustered nonterminal parser (CNP) following [[Scott et al 2019](#Scott-et-al-2019)]. CNP is a version of generalised LL parsing (GLL) [[Scott & Johnstone 2016](#Scott-et-al-2016)]. GLL parsers can parse all context free (CF) languages.

GLL has worst-case cubic time and space complexity but linear complexity for all LL(1) productions [[Scott-et-al-2016](#Scott-et-al-2016)].

# News
## 2020-04-23
1. GoGLL now generates a linear-time lexer matching the CNP parser.
1. This version of GoGLL is fast - v3.0.3 parses its own grammar in 2.99 ms and the total time to compile itself is 27.7 ms. GoGLL v3.0.3 is **faster than Gocc**.

# Walking the parse forest
Gogll produces a binary subtree representation [(BSR)](#Scott-et-al-2019) set of the parse forest of a successful parse. See [Walking the BSR Parse Forest](doc/bsr/bsr.md)

# Benefits and disadvantages
The following table compares GLL parsers with LL-k/LR-k parsers and [PEGs](#Ford-2004)

|| GoGLL v3 | GoGLL V2 | LL-k/LR-k  | PEG
|---|---|---|---|---|
General CF grammars | Yes | Yes | No | No
Composable CF grammars | Yes | Yes | No | No
Handle ambiguity | Yes | Yes | No | No
Indirect left recursion | No problem | No problem | Bad | Bad
Speed (time to compile `gogll.md`) | 0.028 s| 0.244 s | 0.040 s | -

* General CF grammars allow the parser developer to write grammars that match the language most naturally.
* Composability allows pre-existing grammar modules to be imported.
* GLL produces a forest of all valid parses of a string. This provides a more systematic basis for disambiguation than k>1 lookahead and solves the problem of PEGs that hide ambiguity by selecting the first valid parse.
* Operator precedence can be implemented very easily by disambiguating the parse forest [[Afroozeh et al 2013](#Afroozeh-et-al-2013), [Basten & Vinju 2012](#Basten-2012)].

But

* Most non-trivial context free grammars will generate ambiguous parsers, requiring explicit disambiguation.
* GLL parsers are worst-case cubic in time and space complexity. The LL-1 parts of the grammar have linear complexity.

# Motivation to generate a separate lexer
The following observations were made while using GoGLLv2 on a couple of projects.

* Most of the ambiguity in grammars were generated by the lexical rules.
* Handling token separation explicitly produces messy, hard to maintain grammars.
* Most of a grammar input file is whitespace, which together with the additional 
ambiguity introduced by the lexical rules, causes most of the parse time in a 
scannerless parser.
* Writing good markdown with the grammar produced slow compilations.

# Input Symbols, Markdown Files
Gogll accepts UTF-8 input strings. 
A gogll parser has two parse functions: 
* `Parse(I []byte) []*ParseError`
* `func ParseFile(fname string) []*ParseError`   
If `fname` ends with `.md` the parser ignores all text outside the markdown code blocks delimited by triple backticks. See [gogll.md](gogll.md) for an example.

# Gogll Grammar
Gogll v3 has a BNF grammar. See [gogll.md](gogll.md)


# Status
* `gogll v0` was a bootstrap compiler implemented by a [gocc](https://github.com/goccmack/gocc) lexer and parser.
* `gogll v0` was used to compile `gogll v1`.
* `gogll v0` is currently used to compile a proprietary a query language.
* `gogll v1` compiles itself
* The query language mentioned above is being migrated to `gogll v1`.
* `gogll v1` is currently being used to implement a proprietary GUI definition language.

`gogll v1` is actively being developed.

# Features considered for future implementation
1. EBNF grammar support [Scott&Johnstone 2018](#Scott-et-al-2018).
1. Error reporting for normal people.
1. Better documentation, including how to traverse the binary subtree representation ([BSR](#Scott-et-al-2019)) of the parse forest.

# Documentation
At the moment this document and the [gogll grammar](gogll.md) are the only documentation. Have a look at 
`gogll/examples/ambiguous` for a simple example and also for simple disambiguation.

Alternatively look at `gogll.md` which is the input grammar and also the grammar
from which the `parser` for this version of `gogll` was generated. `gogll/da` disambiguates the parse forest for an input string.

# Changelog
[see](ChangeLog.md)

# Bibliography
<a name="Scott-et-al-2019"></a>
* Elizabeth Scott, Adrian Johnstone and L. Thomas van Binsbergen.  
Derivation representation using binary subtree sets.  
In: Science of Computer Programming (175) 2019

<a name="Scott-et-al-2018"></a>
* Elizabeth Scott and Adrian Johnstone.   
GLL Syntax Analysers For EBNF Grammars.   
In: [Science of Computer Programming
Volume 166, 15 November 2018](https://pure.royalholloway.ac.uk/portal/en/publications/gll-syntax-analysers-for-ebnf-grammars(58d1ec5e-28df-486a-879e-36d58a9f8abf).html)

<a name="Scott-et-al-2016"></a>
* Elizabeth Scott and Adrian Johnstone.   
Structuring the GLL parsing algorithm for performance.   
In: [Science of Computer Programming
Volume 125, 1 September 2016](https://pure.royalholloway.ac.uk/portal/en/publications/structuring-the-gll-parsing-algorithm-for-performance(a95fc020-9918-4f17-a87a-845e2aee12b8).html)

<a name="Afroozeh-et-al-2013"></a>
* Ali Afroozeh, Mark van den Brand, Adrian Johnstone, Elizabeth Scott, Jurgen Vinju.   
Safe Specification of Operator Precedence Rules.   
In: [Erwig M., Paige R.F., Van Wyk E. (eds) Software Language Engineering. SLE 2013. Lecture Notes in Computer Science, vol 8225. Springer, Cham](https://pure.royalholloway.ac.uk/portal/en/publications/safe-specification-of-operator-precedence-rules(0287d70e-92b8-4204-aafb-15a81de84968).html)

<a name="Grune-et-al-2012"></a>
* Dick Grune, Kees van Reeuwijk, Henri E. Bal, Ceriel J.H. Jacobs and Koen Langendoen.
Modern Compiler Design. Second Edition.
Springer 2012

<a name="Basten-2012"></a>
* Basten H.J.S., Vinju J.J. (2012) Parse Forest Diagnostics with Dr. Ambiguity. In: Sloane A., Aßmann U. (eds) Software Language Engineering. SLE 2011. [Lecture Notes in Computer Science, vol 6940. Springer, Berlin, Heidelberg](https://homepages.cwi.nl/~jurgenv/papers/SLE2011-2.pdf)

<a name="Ford-2004"></a>
* Bryan Ford. [Parsing Expression Grammars: A Recognition-Based Syntactic Foundation.](https://bford.info/pub/lang/peg.pdf)
