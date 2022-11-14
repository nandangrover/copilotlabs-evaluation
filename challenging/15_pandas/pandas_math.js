function pandas_operation(ds1, ds2) {
    let add = ds1 + ds2;

    let sub = ds1 - ds2;

    let mul = ds1 * ds2;

    let div = ds1 / ds2;

    return [add, sub, mul, div];
}

let ds1 = [2, 4, 6, 8, 10];
let ds2 = [1, 3, 5, 7, 9];
console.log(pandas_operation(ds1, ds2));

module.exports = pandas_operation;