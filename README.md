# Tomb Robber, Ludum Dare 40

![7tombr](https://user-images.githubusercontent.com/12766039/113908375-250c5b00-97df-11eb-96b6-db65025505c9.png)

Made in 72 hours for the Ludum Dare game jam Dec. 2017 by Roninkoi.

You're a grave robber in an ancient egyptian tomb. Your task is to steal as many valuables as possible and get out before you anger the ancient spirits too much. The more you've stolen, the angrier the enemies get. 

The ankhmeter decreases every time you pick up an item. This causes you to slow down, and the enemies make more damage. The meter also replenishes with time. However, levels have a time limit. If the ankhmeter or timer goes to 0, you die.

Written in Go from scratch using WebGL. Compiled to js using gopherjs.

https://roninkoi.itch.io/tomb-robber

### Controls

| Key | Action |
| --- | ------ |
| WASD | Move |
| SPACE | Attack |
| RIGHT/LEFT ARROWS | Rotate camera |

### Building

This requires [GopherJS](https://github.com/gopherjs/gopherjs) and the correct version of golang (in this case go1.12.16). To download dependencies and compile to js:
```
go1.12.16 get
gopherjs build
```

