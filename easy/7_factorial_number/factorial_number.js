// Find factorial of a number

function factorial(num) {
    if (num === 0) {
        return 1;
    }
    return num * factorial(num - 1);
}

console.log(factorial(5));
console.log(factorial(10));
console.log(factorial(15));

module.exports = factorial;