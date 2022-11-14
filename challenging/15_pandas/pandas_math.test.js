const pandas_operation = require("./pandas_math");
const assert = require('assert');

describe("TestProgram", function () {
    it("test_case_1", function () {
        let ds1 = [2, 4, 6, 8, 10];
        let ds2 = [1, 3, 5, 7, 9];
        var actual = pandas_operation(ds1, ds2);
        assert.deepEqual([[3, 7, 11, 15, 19], [1, 1, 1, 1, 1], [2, 12, 30, 56, 90], [2.0, 1.3333333333333333, 1.2, 1.1428571428571428, 1.1111111111111112]], actual);
    });
});