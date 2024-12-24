from collections import deque

# Read the grid from "Input.txt" file
with open("Input.txt", "r") as file:
    grid = [list(line.strip()) for line in file]

rows = len(grid)
cols = len(grid[0])

regions = []
seen = set()

for r in range(rows):
    for c in range(cols):
        if (r, c) in seen: continue
        seen.add((r, c))
        region = {(r, c)}
        q = deque([(r, c)])
        crop = grid[r][c]
        while q:
            cr, cc = q.popleft()
            for nr, nc in [(cr - 1, cc), (cr + 1, cc), (cr, cc - 1), (cr, cc + 1)]:
                if nr < 0 or nc < 0 or nr >= rows or nc >= cols: continue
                if grid[nr][nc] != crop: continue
                if (nr, nc) in region: continue
                region.add((nr, nc))
                q.append((nr, nc))
        seen |= region
        regions.append(region)

def sides(region):
    corner_candidates = set()
    for r, c in region:
        for cr, cc in [(r - 0.5, c - 0.5), (r + 0.5, c - 0.5), (r + 0.5, c + 0.5), (r - 0.5, c + 0.5)]:
            corner_candidates.add((cr, cc))
    corners = 0
    for cr, cc in corner_candidates:
        config = [(sr, sc) in region for sr, sc in [(cr - 0.5, cc - 0.5), (cr + 0.5, cc - 0.5), (cr + 0.5, cc + 0.5), (cr - 0.5, cc + 0.5)]]
        number = sum(config)
        if number == 1:
            corners += 1
        elif number == 2:
            if config == [True, False, True, False] or config == [False, True, False, True]:
                corners += 2
        elif number == 3:
            corners += 1
    return corners

# Calculate and print the result
total = 0
for region in regions:
    region_len = len(region)
    region_sides = sides(region)
    total += region_len * region_sides
    print(f"Region size: {region_len}, Sides: {region_sides}")  # Debugging output

print(f"Total: {total}")
