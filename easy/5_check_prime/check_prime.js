function is_prime(n) {
    if (n == 1) {
        return false;
    }
    for (let i = 2; i < n; i++) {
        if (n % i == 0) {
            return false;
        }
    }
    return true;
}

console.log(is_prime(1));  // false
console.log(is_prime(17));  // true

module.exports = is_prime;