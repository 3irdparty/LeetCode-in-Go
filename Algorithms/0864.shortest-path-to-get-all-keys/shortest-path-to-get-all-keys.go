package problem0864

var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])

	// 按照题目给出的条件，在 30*30 的矩阵中，最多 6 把 key，
	// 使用 6 bit 的整数，就可以记录全部 2^6 = 64 种拥有钥匙的状态
	// 所以，30×30×64 的数组，就可以记录所有的状态。
	// 像这样知道范围的矩阵，尽量使用数组，可以快很多
	hasSeen := [30][30][64]bool{}
	queue := make([][3]int, 1, m*n*4)

	// 获取起点位置，和所有 key 的信息
	allKeys := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := grid[i][j]
			if b == '@' {
				hasSeen[i][j][0] = true
				queue[0] = [3]int{i, j, 0}
			} else if b >= 'a' {
				allKeys |= 1 << uint(b-'a')
			}
		}
	}

	steps := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			x, y, keys := queue[i][0], queue[i][1], queue[i][2]

			if keys == allKeys {
				return steps
			}

			for j := 0; j < 4; j++ {
				nx, ny := x+dx[j], y+dy[j]

				inRange := 0 <= nx && nx < m && 0 <= ny && ny < n
				if !inRange {
					continue
				}

				b := grid[nx][ny]
				if b == '#' ||
					'A' <= b && b <= 'F' && keys&(1<<uint(b-'A')) == 0 {
					// 遇见墙了，或者，没有钥匙开锁
					continue
				}

				nKeys := keys
				if b >= 'a' {
					// 记录刚刚遇到的钥匙
					nKeys |= 1 << uint(b-'a')
				}

				if hasSeen[nx][ny][nKeys] {
					continue
				}

				queue = append(queue, [3]int{nx, ny, nKeys})
				hasSeen[nx][ny][nKeys] = true
			}
		}

		queue = queue[size:]
		steps++
	}

	return -1
}
