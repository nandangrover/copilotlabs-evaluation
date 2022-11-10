// Check leap year

function is_leap_year(year) {
    if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
        return true;
    }
    return false;
}

console.log(is_leap_year(2000));  // true
console.log(is_leap_year(2001));  // false
console.log(is_leap_year(2004));  // true
console.log(is_leap_year(2100));  // false

module.exports = is_leap_year;