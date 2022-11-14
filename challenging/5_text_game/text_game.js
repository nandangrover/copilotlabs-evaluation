const dx = [-1, 1, 0, 0];
const dy = [0, 0, -1, 1];

function maxDistance(grid) {
	const n = grid.length;
	const lands = [];
	for (let i = 0; i < n; i++) {
		for (let j = 0; j < n; j++) {
			if (grid[i][j] === 1) {
				lands.push([i, j]);
			}
		}
	}

	if (lands.length === n * n) {
		return -1;
	}

	let dist = -1;
	while (lands.length > 0) {
		dist++;
		const size = lands.length;
		for (let s = 0; s < size; s++) {
			const g = lands[s];
			const x = g[0];
			const y = g[1];
			for (let k = 0; k < 4; k++) {
				const i = x + dx[k];
				const j = y + dy[k];
				if (i < 0 || n <= i ||
					j < 0 || n <= j ||
					grid[i][j] === 1) {
					continue;
				}
				lands.push([i, j]);
				grid[i][j] = 1;
			}
		}
		lands.splice(0, size);
	}

	return dist;
}

function main() {
	const grid = [
		[1, 0, 1],
		[0, 0, 0],
		[1, 0, 1],
	];
	console.log(maxDistance(grid));
}

main();

module.exports = maxDistance;