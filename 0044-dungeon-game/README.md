# Dungeon Game

## Description

The demons had captured the princess and imprisoned her in the bottom-right corner of a `m x n` dungeon. The dungeon consists of `m x n` rooms laid out in a 2D grid. Our valiant knight was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that increase the knight's health (represented by positive integers).

To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

Return the knight's minimum initial health so that he can rescue the princess.

Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

## Complexity
- Time Complexity: O(m*n)
- Space Complexity: O(m*n)

## Usage
```bash
make run n=0044-dungeon-game
```

## Testing
```bash
make test n=0044-dungeon-game
```
