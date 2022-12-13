# maze-generator

## Usage

```bash
go run .
```

## Game screen

Enter the height and width of the maze. The height and width must be an odd number greater than or equal to 5.

```
Height?
Width?
```

Mazes are randomly generated. Enter the command and go to the goal.

```
#######
#.#...#
#.#@#.#
#...#.#
#.###.#
#.#G..#
#######
move (hjkl), quit (q):
```

| Character | Description |
| --------- | ----------- |
| #         | Wall        |
| .         | Path        |
| @         | Player      |
| G         | Goal        |

## Command

| Key | Action          |
| --- | --------------- |
| h   | Move left       |
| j   | Move down       |
| k   | Move up         |
| l   | Move right      |
| H   | Fast move left  |
| J   | Fast move down  |
| K   | Fast move up    |
| L   | Fast move right |
| q   | Quit            |
