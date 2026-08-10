grid_size = 100
steps = 100
grid = []
def count_neighbors(grid, row, col, target_char='#'):
    rows = len(grid)
    cols = len(grid[0])
    
    OFFSETS = [
        (-1, -1), (-1, 0), (-1, 1),
        ( 0, -1),          ( 0, 1),
        ( 1, -1), ( 1, 0), ( 1, 1)
    ]
    
    count = 0
    for dr, dc in OFFSETS:
        r, c = row + dr, col + dc
        if 0 <= r < rows and 0 <= c < cols:
            if grid[r][c] == target_char:
                count += 1
                
    return count

def partOne(grid, steps):
    temp_grid = [row[:] for row in grid]

    for _ in range(steps):
        for i in range(grid_size):
            for j in range(grid_size):
                count = count_neighbors(grid, i, j)
                if grid[i][j] == '#' and (count < 2 or count > 3):
                    temp_grid[i][j] = '.'
                elif grid[i][j] == '.' and count == 3:
                    temp_grid[i][j] = '#'
        grid = [row[:] for row in temp_grid]

    print(sum(row.count('#') for row in grid))

def partTwo(grid, steps):
    grid[0][0] = '#'
    grid[0][grid_size - 1] = '#'
    grid[grid_size - 1][0] = '#'
    grid[grid_size - 1][grid_size - 1] = '#'
    temp_grid = [row[:] for row in grid]

    for _ in range(steps):
        for i in range(grid_size):
            for j in range(grid_size):
                if (i == 0 and j == 0) or (i == 0 and j == grid_size - 1) or (i == grid_size - 1 and j == 0) or (i == grid_size - 1 and j == grid_size - 1):
                    continue
                count = count_neighbors(grid, i, j)
                if grid[i][j] == '#' and (count < 2 or count > 3):
                    temp_grid[i][j] = '.'
                elif grid[i][j] == '.' and count == 3:
                    temp_grid[i][j] = '#'
        grid = [row[:] for row in temp_grid]

    #print(grid)
    print(sum(row.count('#') for row in grid))

with open("input.txt") as f:
    grid = [[x for x in line.strip()] for line in f.readlines()]
    #partOne(grid, steps)
    partTwo(grid, steps)



