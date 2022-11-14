from typing import List
from collections import deque


def max_distance(grid: List[List[int]]) -> int:
    n = len(grid)
    lands = []
    for i in range(n):
        for j in range(n):
            if grid[i][j] == 1:
                lands.append([i, j])

    if len(lands) == n * n:
        return -1

    dist = -1
    # BFS
    lands = deque(lands)
    while lands:
        dist += 1
        size = len(lands)
        for _ in range(size):
            x, y = lands.popleft()
            for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1)):
                i, j = x + dx, y + dy
                if 0 <= i < n and 0 <= j < n and grid[i][j] == 0:
                    lands.append([i, j])
                    grid[i][j] = 1

    return dist


def main():
    grid = [
        [1, 0, 1],
        [0, 0, 0],
        [1, 0, 1],
    ]
    print(max_distance(grid))


if __name__ == "__main__":
    main()