function alphabetTriangle() {
    for (let i = 'A'.charCodeAt(0); i <= 'G'.charCodeAt(0); i++) {
        for (let j = 'A'.charCodeAt(0); j <= i; j++) {
            process.stdout.write(String.fromCharCode(j) + ' ');
        }
        process.stdout.write('\n');
    }
}

alphabetTriangle();

module.exports = alphabetTriangle;