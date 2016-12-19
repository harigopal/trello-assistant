# Tassi?

**T**rello **Assi**stant. It automates some of the repetitive tasks I find myself doing every week, checking the burn-down of our Trello board at [SV.CO](https://www.sv.co), and such.

## Environment variables

1. `TRELLO_KEY` - Trello developer key
2. `TRELLO_TOKEN` - Authenticated Trello token for above key.
3. `TRELLO_BOARD_ID` - ID of board that owner of credentials above has access to.

## Example run

```
$ tassi

Loading board... Done. Loaded Engineering
Loading all lists... Done. Loaded 11 lists.
Loading cards stats........... Done. Evaluated 134 cards.
+-------------------------------+-----------------+-------------+----------+-------------+----------+
|             LIST              | NUMBER OF CARDS | TOTAL SCORE | EASY (1) | AVERAGE (2) | HARD (3) |
+-------------------------------+-----------------+-------------+----------+-------------+----------+
| Ideas                         |              26 |           1 |        1 |           0 |        0 |
| Needs Discussion              |               9 |           0 |        0 |           0 |        0 |
| Accepted Ideas                |              20 |          45 |        2 |          11 |        7 |
| Ready For Roadmap             |               9 |          13 |        1 |           6 |        0 |
| Optimization, Internal & Bugs |              22 |          26 |        3 |          10 |        1 |
| Rollbar                       |              18 |           0 |        0 |           0 |        0 |
| Code Quality                  |               7 |           4 |        4 |           0 |        0 |
| Up Next                       |               6 |           6 |        1 |           1 |        1 |
| Doing                         |               3 |           8 |        0 |           1 |        2 |
| Review                        |               2 |           3 |        1 |           1 |        0 |
| Done                          |              12 |          18 |        7 |           4 |        1 |
+-------------------------------+-----------------+-------------+----------+-------------+----------+
```

## Execution

Build and use the binary.

```
$ go build src/github.com/harigopal/trello-assistant/tassi.go
```

## Why?

Because [developers are lazy](http://threevirtues.com/).

Plus, I'm a Golang noob, and this is my first use of Go for something _real_.
