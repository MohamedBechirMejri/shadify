<a href="https://goreportcard.com/report/github.com/cheatsnake/shadify"><img src="https://goreportcard.com/badge/github.com/cheatsnake/shadify" alt="Go report"/></a>
<a href="https://github.com/cheatsnake/shadify/releases"><img src="https://img.shields.io/github/v/release/cheatsnake/shadify.svg" alt="GitHub repo size"/></a>
<a href="https://img.shields.io/github/repo-size/cheatsnake/shadify?color=blue"><img src="https://img.shields.io/github/repo-size/cheatsnake/shadify?color=blue" alt="GitHub repo size"/></a>
<a href="https://img.shields.io/github/license/cheatsnake/shadify?color=orange"><img src="https://img.shields.io/github/license/cheatsnake/shadify?color=orange" alt="GitHub repo license"/></a>
<a href="https://github.com/cheatsnake/shadify/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="GitHub repo issues"/></a>

# Shadify

Focus on creating incredible and beautiful applications without worrying about complicated logic.

## 🔬 Quick overview

Shadify is a powerful API for generating data and executing various logic to create different applications and games. It uses HTTP requests to interact with the service and returns convenient JSON responses. Shadify can be very useful for developers who do Frontend-development. With the modules already available, you can focus entirely on creating beautiful UIs for applications that will enrich your portfolio.

The codebase is written in the Go programming language, which is great for creating server-side applications, as well as providing high computational power.

## 📃 Documentation
<div align="center"><a href="https://github.com/cheatsnake/shadify/blob/master/README_RU.md#%D0%B4%D0%BE%D0%BA%D1%83%D0%BC%D0%B5%D0%BD%D1%82%D0%B0%D1%86%D0%B8%D1%8F">Доступно на русском</a></div>
</br>

The service is divided into independent modules. Each module starts with a brief description of what the module is oriented at (be it a game, a puzzle, a task, etc.). This is followed by a detailed description of each HTTP interface, with descriptions of the possible parameters and return responses.

-   [Sudoku](#sudoku)
-   [Takuzu](#takuzu)
-   [Set](#set)
-   [Math](#math)
-   [Schulte](#schulte)
-   [Minesweeper](#minesweeper)
-   [Wordsearch](#wordsearch)
-   [Anagram](#anagram)
-	[Countries](#countries)
-	[Camp](#camp)
-	[Kuromasu](#kuromasu)

### Sudoku

<p align="center"><img src="./images/sudoku.png" alt="Sudoku"/></p>

[Sudoku](https://en.wikipedia.org/wiki/Sudoku) is a popular number puzzle game. The essence of the game is to fill free cells with numbers from 1 to 9 so that in each row, each column and each small 3×3 square each digit occurs only once.

-   Generating random Sudoku

```rust
GET https://shadify.dev/api/sudoku/generator
```

| Parameter | Description                                                                                                                      |
| --------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `fill`    | _Optional_ <br> A number from 0 to 100 that corresponds to the field fill level (as a percentage). <br> The default value is 30. |

Returned response:

```json
{
    "grid": [
        [9, 8, 7, 3, 2, 1, 5, 6, 4],
        [6, 5, 4, 9, 8, 7, 2, 3, 1],
        [3, 2, 1, 6, 5, 4, 8, 9, 7],
        [7, 6, 5, 1, 9, 8, 3, 4, 2],
        [1, 9, 8, 4, 3, 2, 6, 7, 5],
        [4, 3, 2, 7, 6, 5, 9, 1, 8],
        [8, 7, 6, 2, 1, 9, 4, 5, 3],
        [5, 4, 3, 8, 7, 6, 1, 2, 9],
        [2, 1, 9, 5, 4, 3, 7, 8, 6]
    ],
    "task": [
        [0, 8, 7, 0, 0, 0, 0, 0, 4],
        [0, 0, 0, 9, 0, 0, 2, 0, 0],
        [0, 2, 1, 6, 5, 4, 0, 0, 7],
        [0, 0, 0, 1, 0, 8, 0, 4, 2],
        [0, 9, 8, 4, 3, 0, 0, 0, 0],
        [4, 0, 2, 0, 0, 0, 0, 0, 0],
        [8, 0, 6, 2, 0, 0, 0, 0, 0],
        [5, 0, 3, 0, 0, 6, 0, 2, 9],
        [0, 0, 0, 0, 4, 0, 0, 0, 0]
    ]
}
```

> - `grid` - fully solved Sudoku. <br> 
> - `task` - the task, which must be solved by filling in the zeros with the correct numbers.

-   Checking Sudoku

To check Sudoku you can use GET-request with required query-parameter

```rust
GET https://shadify.dev/api/sudoku/verifier
```

| Parameter | Description                                                                                                                              |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `task`    | _Required_ <br> A string of the form `123-123-123`, which corresponds to the Sudoku field, where each line (row) is separated by a dash. |

An example of a GET-request:

```rust
https://shadify.dev/api/sudoku/verifier?task=123564897-456897231-789231564-234675918-567918342-891342675-345786129-678129453-912453786
```

Also, for convenience, you can use POST-request with body in JSON format

```rust
POST https://shadify.dev/api/sudoku/verifier
```

Example body for a POST request:

```json
{
    "task": [
        [9, 6, 3, 2, 8, 5, 7, 4, 1],
        [8, 5, 2, 1, 7, 4, 6, 3, 9],
        [1, 7, 4, 3, 9, 6, 8, 5, 2],
        [6, 3, 9, 8, 5, 2, 3, 1, 7],
        [5, 2, 8, 7, 4, 1, 3, 9, 6],
        [7, 4, 1, 9, 6, 3, 5, 2, 8],
        [4, 1, 7, 6, 3, 9, 2, 8, 5],
        [3, 9, 6, 5, 2, 8, 1, 7, 4],
        [2, 8, 5, 4, 1, 7, 9, 6, 3]
    ]
}
```

Returned response:

```json
{
    "isError": true,
    "position": "row-4"
}
```

### Takuzu

<p align="center"><img src="./images/takuzu.png" alt="Takuzu"/></p>

[Takuzu](https://en.wikipedia.org/wiki/Takuzu) (a.k.a. Binairo) is an entertaining puzzle game with simple rules. All you have to do is to fill a square field of a certain size with two digits (or colors) while following three simple rules:

1. Each column and each row must be unique
2. Each row and each column must have an equal number of tiles of each digit
3. no more than two tiles with the same digit in a row (110001 is wrong, 110010 is right).

For an example, you can see my app [Binario](https://binario.vercel.app)

-   Generating Takuzu

```rust
GET https://shadify.dev/api/takuzu/generator
```

| Parameter | Description                                                                                                                      |
| --------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `size`    | _Optional_ <br> An even number from 4 to 12, which specifies the size of the field. <br> The default value is 8.                 |
| `fill`    | _Optional_ <br> A number from 0 to 100 that corresponds to the field fill level (as a percentage). <br> The default value is 33. |

Returned response:

```json
{
    "size": 8,
    "field": [
        ["1", "1", "0", "0", "1", "1", "0", "0"],
        ["1", "1", "0", "1", "0", "0", "1", "0"],
        ["0", "0", "1", "0", "1", "1", "0", "1"],
        ["0", "1", "1", "0", "0", "1", "1", "0"],
        ["1", "0", "0", "1", "1", "0", "0", "1"],
        ["0", "0", "1", "1", "0", "0", "1", "1"],
        ["0", "1", "1", "0", "1", "1", "0", "0"],
        ["1", "0", "0", "1", "0", "0", "1", "1"]
    ],
    "task": [
        ["1", "1", "0", "0", "x", "x", "0", "0"],
        ["x", "1", "0", "x", "0", "x", "x", "x"],
        ["x", "x", "1", "x", "x", "x", "x", "x"],
        ["x", "1", "1", "x", "x", "x", "1", "0"],
        ["x", "x", "x", "x", "x", "x", "x", "1"],
        ["0", "x", "x", "x", "x", "x", "x", "1"],
        ["x", "x", "x", "x", "x", "1", "0", "x"],
        ["x", "x", "x", "x", "x", "x", "1", "1"]
    ]
}
```

-   Checking Takuzu

To check Takuzu, you can use a GET-request with the required query parameter:

```rust
GET https://shadify.dev/api/takuzu/verifier
```

| Parameter | Description                                                                                                                                                                                                                                           |
| --------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `task`    | _Required_ <br> A string of the form `1010-1100-0011-0101`, which corresponds to the Takuzu field, where each line (row) is separated by a dash. It can only contain 0 and 1, the number of rows must be equal to the number of elements in each row. |

An example of a GET-request:

```rust
https://shadify.dev/api/takuzu/verifier?task=1010-1100-0101-0101
```

Also, for convenience, you can use POST-request with body in JSON format

```rust
POST https://shadify.dev/api/takuzu/verifier
```

Example body for a POST-request:

```json
{
    "task": [
        ["1", "0", "1", "0"],
        ["1", "1", "0", "0"],
        ["0", "0", "1", "1"],
        ["0", "1", "0", "1"]
    ]
}
```

Returned response:

```json
{
    "isError": true,
    "message": "duplication",
    "position": ["row-3", "row-4"]
}
```

### Set

<p align="center"><img src="./images/set.png" alt="Set"/></p>

[Set (card game)](<https://en.wikipedia.org/wiki/Set_(card_game)>) – a fascinating card game. The game deck consists of 81 cards, each with one, two, or three of the same symbol (rhombus, oval, or wave) of the same color (red, green, or purple) and texture (shaded, shaded, or outline only). The essence of the game is to find a set - a set of three cards that meets certain conditions.

To understand the rules, read Wikipedia or watch [this interesting video](https://youtu.be/NzXDfSFQ1c0).

-   Get all cards

```rust
GET https://shadify.dev/api/set/start
```

Always returns the same array of 81 objects. Each object corresponds to one of the cards. An example of a card:

```json
{
    "_id": 0,
    "number": 1,
    "shape": "diamond",
    "color": "red",
    "shading": "solid"
}
```

> - `_id` - a unique identifier for each card <br> 
> - `number` – number of figures: _1_ / _2_ / _3_ <br> 
> - `shape` – body shape: _diamond_ / _squiggle_ / _oval_ <br> 
> - `color` – figure color: _red_ / _green_ / _purple_ <br> 
> - `shading` – figure shading: _solid_ / _striped_ / _open_

-   Generating a new game

```rust
GET https://shadify.dev/api/set/start
```

| Parameter      | Description                                                                                                                                                                                                                                                                       |
| -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `possibleSets` | _Optional_ <br> A true/false string that enables/disables the search for possible sets in the current `layout`. The list of possible sets is not necessary for the game and acts only as a hint and evidence that sets exist in the current `layout`. <br> Default value is true. |

Returned response:

```json
{
	"freeCards": [<69 cards>],
	"layout": [<12 cards>],
	"possibleSets": [[<3 cards>]],
	"wonCards": [],
	"state": "20-4-11-10-12-46-70-62-41-23-3-8"
}
```

> - `freeCards` – an array of objects corresponding to free cards that have not yet been used in the game. <br> 
> - `layout` – an array of objects corresponding to the card that are available to play, i.e. to search for sets. <br> 
> - `possibleSets` – array containing arrays that include exactly 3 objects each. Each 3 object corresponds to a combination of three cards forming a set, which can be assembled from the cards available on the current `layout`. <br> 
> - `wonCards` – an array of objects corresponding to the won cards, which will no longer participate in the game. <br> 
> - `state` – a unique identifier for the current game state. It is needed to perform actions on deleting sets, adding additional cards or just to load the current game state. It has the following form _1-2-3@4-5-6_, where the numbers to the left of the sign _@_ correspond to the unique identifiers of those cards that are in `layout`, and the numbers on the right are in `wonCards`.

-   Load state

```rust
GET https://shadify.dev/api/set/<:state>
```

| Parameter      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| -------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `possibleSets` | _Optional_ <br> A _true_/_false_ string that enables/disables the search for possible sets in the current `layout`. <br> Default value is _true_.                                                                                                                                                                                                                                                                                                                                   |
| `action`       | _Optional_ <br> The _add_/_remove_ string, which allows you to perform the appropriate action with the current game state. <br> The _add_ string adds 3 random cards from the current `freeCards` array to the current `layout` (available only if the `layout` size does not exceed 20 cards). <br> The _remove_ string removes the specified combination of three cards from the current `layout`. To do this you must use the `cards` parameter. <br> There is no default value. |
| `cards`        | _Required for `action=remove`_ <br> A string of the form _1-2-3_, where each number corresponds to the unique identifier of one of the cards that make up the set. <br> There is no default value.                                                                                                                                                                                                                                                                                  |

Examples of requests with state loading:

```rust
https://shadify.dev/api/set/0-27-53-10-46-15-16-64-32-23-29-6?possibleSets=false
```

```rust
https://shadify.dev/api/set/41-47-7-53-13-46-25-36-72-60-15-80?action=add
```

```rust
https://shadify.dev/api/set/0-27-53-10-46-15-16-64-32-23-29-6?action=remove&cards=0-16-23
```

### Math

<p align="center"><img src="./images/math-expressions.png" alt="Math-expressions"/></p>

A module for generating random mathematical expressions. Great for creating various simulators, where you'll have to wiggle your brain. You can see [my application](https://cheatsnake.github.io/MindMath/) as an example.

-   Generating an expression with addition


```rust
GET https://shadify.dev/api/math/add
```

-   Generating an expression with subtraction

```rust
GET https://shadify.dev/api/math/sub
```

-   Generating an expression with multiplication

```rust
GET https://shadify.dev/api/math/mul
```

-   Generating an expression with division

```rust
GET https://shadify.dev/api/math/div
```

The following parameters are available for the above endpoints:

| Parameter                | Description                                                                                                                                                                                                                                                                                                                                                       |
| ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `minFisrt`, `maxFirst`   | _Optional_ <br> Parameters that change the range of generated numbers only for the first number present in the expression. <br> The default value for `minFirst` is 1, for `maxFirst` it is 99.                                                                                                                                                                   |
| `minSecond`, `maxSecond` | _Optional_ <br> Parameters that change the range of generated numbers only for the second number present in the expression. These parameters do not affect expressions with division, because the second number is chosen randomly from the list of possible divisors of the first number. <br> The default value for `minSecond` is 1, for `maxSecond` it is 99. |
| `negative`               | _Optional_ <br> The number 0/1 which enables/disables the possibility of generating a result with a negative value for expressions with subtraction. Initially, when generating an expression with subtraction, the first number is always greater than the second number. To change this behavior, use this parameter. <br> Default value is 0.                  |

Returned response:

```json
{
    "first": 17,
    "second": 56,
    "operation": "+",
    "expression": "17 + 56",
    "answer": 73
}
```

-   Generating a quadratic equation

<p align="center"><img src="./images/quadratic-equation.png" alt="Quadratic-equation"/></p>

[Quadratic equation](https://en.wikipedia.org/wiki/Quadratic_equation) is an equation of the form ax2 + bx + c = 0, where the coefficients a, b and c are arbitrary numbers.

```rust
GET https://shadify.dev/api/math/quad
```

| Parameter      | Description                                                                                                                   |
| -------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `minA`, `maxA` | _Optional_ <br> Minimum and maximum values of the coefficient A. <br> The default value for `minA` is 1, for `maxA` it is 20. |
| `minB`, `maxB` | _Optional_ <br> Minimum and maximum values of the coefficient B. <br> The default value for `minB` is 1, for `maxB` it is 40. |
| `minC`, `maxC` | _Optional_ <br> Minimum and maximum values of the coefficient C. <br> The default value for `minC` is 1, for `maxC` it is 20. |

Returned response:

```json
{
    "equation": "1x^2 + 14x + 24 = 0",
    "a": 1,
    "b": 14,
    "c": 24,
    "discriminant": 100,
    "x1": "-12.000",
    "x2": "-2.000"
}
```

### Schulte

<p align="center"><img src="./images/schulte.gif" alt="Schulte"/></p>

[Schulte Tables](https://youtu.be/4R1kycHqsqM) – tables with randomly arranged objects (usually numbers or letters) used to test and develop quickness of finding these objects in a certain order (usually in ascending order for numbers and alphabetical order for letters). Exercises with tables can improve peripheral visual perception, which will have a positive impact on the skill of speed reading.

-   Generating a Random Table

```rust
GET https://shadify.dev/api/schulte/generator
```

| Parameter | Description                                                                                                                                                                                                               |
| --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `size`    | _Optional_ <br> The size of the generated table. The available range is from 1 to 15. <br> The default value is 5.                                                                                                        |
| `mode`    | _Optional_ <br> The _number_ / _alphabet_ string, which defines the generated values for the table: numbers or letters. If _alphabet_ is selected, the `size` parameter will always be 5. <br> Default value is _number_. |

Returned response:

```json
{
    "grid": [
        [9, 18, 6, 1, 5],
        [11, 24, 25, 15, 19],
        [13, 7, 10, 21, 23],
        [20, 22, 12, 17, 16],
        [2, 8, 4, 14, 3]
    ]
}
```

### Minesweeper

<p align="center"><img src="./images/minesweeper.png" alt="Minesweeper"/></p>

[Minesweeper](<https://en.wikipedia.org/wiki/Minesweeper_(video_game)>) – a computer puzzle game in which the playing field is divided into adjacent cells, some of which are mined. The number of mined cells is known. The goal of the game is to open all the cells that do not contain mines. This game has become quite popular among Windows users, since it was pre-installed by default on older versions of that OS.

-   Generating a random field

```rust
GET https://shadify.dev/api/minesweeper/generator
```

| Parameter | Description                                                                                                                                                                                                                                                                                                                                                   |
| --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `start`   | _Required_ <br> A string of the form _1-2_, which sets the starting position of the player. There will never be mines in and around this position. The first number is the X coordinate (from 1 to the value of the `widht` parameter), the second number is the Y coordinate (from 1 to the value of the `height` parameter) <br> There is no default value. |
| `width`   | _Optional_ <br> The number that sets the width of the generated field. The total number of cells in the field must not exceed 1000. <br> The default value is 9.                                                                                                                                                                                              |
| `height`  | _Optional_ <br> The number that sets the height of the generated field. The total number of cells in the field must not exceed 1000. <br> The default value is 9.                                                                                                                                                                                             |
| `mines`   | _Optional_ <br> The number that sets the number of mines on the field. The number of mines must not exceed 25% of the total number of cells on the field. <br> The default value is 12.                                                                                                                                                                       |

Returned response:

```json
{
    "start": "3-5",
    "width": 9,
    "height": 9,
    "board": [
        ["o", "o", "1", "x", "3", "x", "2", "x", "2"],
        ["1", "1", "1", "2", "x", "2", "2", "2", "x"],
        ["x", "1", "o", "1", "1", "1", "o", "1", "1"],
        ["1", "1", "o", "o", "o", "o", "o", "o", "o"],
        ["o", "o", "o", "1", "1", "1", "o", "o", "o"],
        ["2", "2", "1", "1", "x", "1", "o", "o", "o"],
        ["x", "x", "2", "2", "1", "1", "o", "1", "1"],
        ["x", "4", "x", "1", "o", "o", "o", "1", "x"],
        ["1", "2", "1", "1", "o", "o", "o", "1", "1"]
    ],
    "mines": 12
}
```

### Wordsearch

<p align="center"><img src="./images/wordsearch.png" alt="Wordsearch"/></p>

[Word search](https://en.wikipedia.org/wiki/Word_search) is a puzzle consisting of letters of words placed in a square or rectangular grid. The aim of the puzzle is to find and mark all the words hidden in the grid. The words can be placed horizontally, vertically or diagonally.

-   Generating a random grid

```rust
GET https://shadify.dev/api/wordsearch/generator
```

| Parameter | Description                                                                                                                                                          |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `width`   | _Optional_ <br> A number from 5 to 20 that specifies the width of the grid. The total number of cells in the field must not exceed 256. <br> The default value is 9. |
| `height`  | _Optional_ <br> A number from 5 to 20 that specifies the width of the grid. The total number of cells in the field must not exceed 256. <br> The default value is 9. |

Returned response:

```json
{
    "width": 9,
    "height": 9,
    "wordsCount": 10,
    "grid": [
        ["e", "n", "r", "w", "r", "o", "l", "o", "c"],
        ["t", "o", "o", "o", "a", "l", "n", "b", "p"],
        ["e", "o", "o", "l", "z", "a", "w", "r", "a"],
        ["l", "n", "s", "p", "r", "e", "o", "u", "n"],
        ["h", "r", "t", "w", "e", "m", "t", "g", "t"],
        ["t", "e", "e", "o", "l", "t", "n", "b", "s"],
        ["a", "t", "r", "n", "l", "a", "w", "y", "t"],
        ["v", "f", "s", "s", "e", "o", "o", "f", "j"],
        ["i", "a", "b", "d", "t", "z", "d", "s", "n"]
    ],
    "words": [
        { "word": "color", "position": { "start": [9, 1], "end": [5, 1] } },
        { "word": "downtown", "position": { "start": [7, 9], "end": [7, 2] } },
        { "word": "teller", "position": { "start": [5, 9], "end": [5, 4] } },
        { "word": "pants", "position": { "start": [9, 2], "end": [9, 6] } },
        { "word": "athlete", "position": { "start": [1, 7], "end": [1, 1] } },
        { "word": "afternoon", "position": { "start": [2, 9], "end": [2, 1] } },
        { "word": "snowplow", "position": { "start": [4, 8], "end": [4, 1] } },
        { "word": "rooster", "position": { "start": [3, 1], "end": [3, 7] } },
        { "word": "rugby", "position": { "start": [8, 3], "end": [8, 7] } },
        { "word": "oatmeal", "position": { "start": [6, 8], "end": [6, 2] } }
    ]
}
```

### Anagram

<p align="center"><img src="./images/anagram.png" alt="Anagram"/></p>

[Anagrams](https://en.wikipedia.org/wiki/Anagrams_(game)) is a whole type of puzzles associated with composing all possible words from a given set of letters. This module implements the simplest variation of anagrams: a random word is given, the letters of which should be used to compose as many other words as possible. 

Only English nouns are used for composition, a list of which can be found in this repository at the path `./data/nouts.txt`.

```rust
GET https://shadify.dev/api/anagram/generator
```

Returned response:

```json
{
    "task":"possibility",
    "words":[
        "bit","bolt","boy","lip","list","loss","oil",
        "pilot","plot","toy","pot","slip","soil","soy",
        "spot","spy","stop","tip","top"
    ]
}
```

> - `task` - a word from which you need to make other words. <br> 
> - `words` - an array of all possible words that are compiled from `task`. 

### Countries

Countries module allows you to generate quizes such as guess the capital or guess the country from an image of the flag. It is a simple and useful module for creating applications to test and practice knowledge of all countries.

-   Generating capital quiz

<p align="center"><img src="./images/guess-capital.png" alt="Guess capital"/></p>

```rust
GET https://shadify.dev/api/countries/capital-quiz
```

| Parameter | Description                                                                                                                      |
| --------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `variants` | _Optional_ <br> A number from 2 to 6 corresponding to the number of different options from which you have to choose the correct capital of the given country. <br> The default value is 4. |
| `amount`   | _Optional_ <br> A number from 1 to 20 which is responsible for the number of quizzes returned. The use of this parameter ensures that among all quizzes received, all will be unique. <br> The default value is 1. |

Returned response:

```json
{
	"country": "Cyprus",
	"flag": "https://flagcdn.com/w320/cy.png",
	"variants": [
		"Nicosia",
		"Juba",
		"Oslo",
		"Jamestown"
	],
	"answer": "Nicosia"
}
```

> - `country` - the country for which you must guess the capital. <br> 
> - `flag` - country's flag image (it powered by [Flagpedia API](https://flagpedia.net/download/api), so you can customize it). <br> 
> - `variants` - possible options for answering. <br> 
> - `answer` - correct answer.

-   Generating country quiz

<p align="center"><img src="./images/guess-country.png" alt="Guess country"/></p>

```rust
GET https://shadify.dev/api/countries/country-quiz
```

| Parameter  | Description                                                                                                                      |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `variants` | _Optional_ <br> A number from 2 to 6 corresponding to the number of different options from which you have to choose the correct country of the given flag image. <br> The default value is 4. |
| `amount`   | _Optional_ <br> A number from 1 to 20 which is responsible for the number of quizzes returned. The use of this parameter ensures that among all quizzes received, all will be unique. <br> The default value is 1. |

Returned response:

```json
{
	"flag": "https://flagcdn.com/w320/dk.png",
	"variants": [
		"Yemen",
		"Denmark",
		"Norfolk Island",
		"Vietnam"
	],
	"answer": "Denmark"
}
```

> `flag` - the flag of the country you have to guess. <br> `variants` - possible options for answering. <br> `answer` - correct answer.

### Camp

<p align="center"><img src="./images/camp.png" alt="Camp"/></p>

[Camp (aka Tents and trees, Tents)](https://www.chiark.greenend.org.uk/~sgtatham/puzzles/js/tents.html) is a logic puzzle with simple rules and challenging solutions. The rules of Camp are simple:

1. Pair each tree with a tent adjacent horizontally or vertically. This should be a 1 to 1 relation.
2. Tents never touch each other even diagonally
3. The clues outside the grid indicate the number of tents on that row/column.

-   Generating a random task

```rust
GET https://shadify.dev/api/camp/generator
```

| Parameter  | Description                                                                                                                      |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `width`    | _Optional_ <br> A number from 5 to 15 which corresponds to the width of the generated field. <br> The default value is 7.        |
| `height`   | _Optional_ <br> A number from 5 to 15 which corresponds to the height of the generated field. <br> The default value is 7.       |
| `solution` | _Optional_ <br> The true/false value that specifies whether the solution should be sent with the task or not. <br> The default value is true.       |

Returned response:

```json
{
    "width": 7,
    "height": 7,
    "trees": 9,
    "rowTents": [1, 2, 2, 1, 0, 2, 1],
    "columnTents": [1, 1, 0, 2, 0, 2, 3],
    "task": [
        [0, 0, 0, 1, 0, 0, 1],
        [0, 0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 0, 0, 1],
        [0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 0]
    ],
    "solution": [
        [0, 0, 0, 1, 0, 2, 1],
        [0, 0, 0, 2, 0, 1, 2],
        [0, 2, 0, 0, 0, 2, 0],
        [0, 1, 0, 0, 0, 1, 2],
        [0, 0, 0, 0, 0, 0, 1],
        [2, 0, 0, 2, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 2]
    ]
}
```

> - `width` – the width of the generated field. <br> 
> - `heigth` – the height of the generated field. 
> - `trees` – the total number of trees on the field (and therefore the total number of tents that need to be set up). <br> 
> - `rowTents` – an array of values that corresponds to the number of tents that should be placed in the corresponding row from 1 to the `height` value. <br> 
> - `columnTents` – an array of values that corresponds to the number of tents that should be placed in the corresponding column from 1 to the `width` value. <br> 
> - `task` – the task that needs to be solved. Values of ones are trees around which you must place values of twos, which correspond to tents. <br> 
> - `solution` – a completely solved task, where the values of ones are trees, and the values of twos are tents.

- Verifying a completed task

```rust
POST https://shadify.dev/api/camp/verifier
```

Example of JSON body: 

```json
{
    "rowTents": [1, 2, 2, 1, 0, 2, 1],
    "columnTents": [1, 1, 0, 2, 0, 2, 3],
    "solution": [
        [0, 0, 2, 1, 0, 2, 1],
        [0, 0, 0, 0, 0, 1, 2],
        [0, 2, 0, 0, 0, 2, 0],
        [0, 1, 0, 0, 0, 1, 2],
        [0, 0, 0, 0, 0, 0, 1],
        [2, 0, 0, 2, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 2]
    ]
}
```

Returned response:

```json
{
	"isError": true,
	"position": "row-1",
	"message": "the number of tents in each row must match the assignment"
}
```

### Kuromasu

<p align="center"><img src="./images/kuromasu.png" alt="Kuromasu"/></p>

[Kuromasu](https://en.wikipedia.org/wiki/Kuromasu) is a puzzle played on a rectangular grid. Some of these cells have numbers in them. Each cell may be either black or white. The object is to determine what type each cell is. The following rules determine which cells are which:

1. Each number on the board represents the number of white cells that can be seen from that cell, including itself. A cell can be seen from another cell if both cells are within the same row or column, and there are no black cells between them in that row or column.
2. Numbered cells must not be black.
3. No two black cells must be horizontally or vertically adjacent.
4. All the white cells must be connected horizontally or vertically.

-   Generating a random task

```rust
GET https://shadify.dev/api/kuromasu/generator
```

| Parameter  | Description                                                                                                                      |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `width`    | _Optional_ <br> A number from 4 to 15 which corresponds to the width of the generated field. <br> The default value is 5.        |
| `height`   | _Optional_ <br> A number from 4 to 15 which corresponds to the height of the generated field. <br> The default value is 5.       |
| `fill`     | _Optional_ <br> A number from 10 to 50 which corresponds to the % level of filling the task with ready cells. <br> The default value is 30.       |

Returned response:

```json
{
  "width": 5,
  "height": 5,
  "task": [
    ["o", "o", "o", "o", "6"],
    ["o", "o", "8", "o", "9"],
    ["o", "o", "o", "7", "o"],
    ["5", "o", "o", "o", "o"],
    ["9", "o", "8", "5", "o"]
  ],
  "solution": [
    ["6", "4", "x", "4", "6"],
    ["9", "7", "8", "7", "9"],
    ["9", "7", "8", "7", "9"],
    ["5", "x", "4", "x", "5"],
    ["9", "5", "8", "5", "9"]
  ]
}
```

> - `width` – the width of the generated grid. <br> 
> - `heigth` – the height of the generated grid. <br> 
> - `task` – the task that needs to be solved by arranging the boxes (values of `x`) in the right cells. <br> 
> - `solution` – a completely solved task, where cells with `x` values are boxes and numbers is just view counts of each particular cell.

- Verifying a completed task

```rust
POST https://shadify.dev/api/kuromasu/verifier
```

Example of JSON body: 

```json
{
  "solution": [
    ["6", "4", "x", "4", "x"],
    ["9", "7", "8", "7", "9"],
    ["9", "7", "8", "7", "9"],
    ["5", "x", "4", "x", "5"],
    ["9", "5", "8", "5", "9"]
  ]
}
```

Returned response:

```json
{
	"isValid": false,
	"position": {
		"row": 1,
		"column": 4
	}
}
```

## 🚀 Server startup

1. Clone this repository onto your computer:

```sh
git clone https://github.com/cheatsnake/shadify.git
```

2. Inside the project, run this command to install the necessary packages:

```sh
go mod download
```

> Make sure you have already [installed Go](https://go.dev) on your computer.

3. Start the server by running the last command:

```sh
go run cmd/server/main.go
```

> The server will start at the address http://localhost:5000

For easy testing, use a ready-made collection for [Insomnia](https://insomnia.rest). Open `Settings` > `Data` > `Import Data` > `From URL` and paste link to the [insomnia.shadify.json](./insomnia.shadify.json) file. Enjoy!

## 🐳 Docker container startup

- Running a Docker container for development (after any file changes, server will be restarts):

```sh
docker compose up
```

- Running a Docker container for production:

```sh
docker build -t shadify . --target prod
```

```sh
docker run --rm -p 5000:5000 --name shadify shadify
```

<div align="center">Made with &#9829;</div>
<div align="center"><a href="https://github.com/cheatsnake/shadify/blob/master/LICENSE">LICENSE</a> 2022-2023</div>
