# Langton's Ant in Go

Langton's Ant is a simple algorithm created by Chris Langton in 1986. It is famous for its unusual chaotic behavior. The ant moves on a two-dimensional board and follows several rules:

* If the current cell is white, the ant changes it to black, turns right, and moves one step forward.
* If the current cell is black, the ant changes it to white, turns left, and moves one step forward.

## In my implementation, you can find:

### 1. `Ant`

* `x` - the horizontal coordinate;
* `y` - the vertical coordinate;
* `direction` - the direction in which the ant is facing.

### 2. `Board`

* `height` - the vertical size of the board;
* `width` - the horizontal size of the board;
* `field` - a two-dimensional collection of cells.

### 3. `Cell Colors`

* `○` — black cell;
* `●` — white cell;
* `M` — the current position of the ant.

## Running the program

```bash
go run . \
  -width=20 \
  -height=20 \
  -iterations=500 \
  -delay=100ms
```

Available flags:

* `-width` — board width;
* `-height` — board height;
* `-iterations` — number of simulation steps;
* `-delay` — delay between frames, for example `100ms` or `1s`.

To display help:

```bash
go run . -help
```
