// Number Pyramid Pattern

function pyramidPattern(n) {
    for (let i = 1; i <= n; i++) {
        for (let j = 0; j < n-i; j++) {
            process.stdout.write(" ")
        }
        for (let k = 0; k < i; k++) {
            process.stdout.write(i.toString())
        }
        process.stdout.write("\n")
    }
}

pyramidPattern(5);

module.exports = pyramidPattern;
