function invertedStar(n) {
	for (let i = 0; i <= n; i++) {
		for (let j = 0; j < n-i; j++) {
			process.stdout.write(" ")
		}
		for (let k = 0; k < i; k++) {
			process.stdout.write("*")
		}
		process.stdout.write("\n")
	}
}

invertedStar(5)

module.exports = invertedStar;