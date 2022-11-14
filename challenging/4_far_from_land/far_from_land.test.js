var tcs = [
  {
    grid: [
      [1, 1, 1],
      [1, 1, 1],
      [1, 1, 1],
    ],
    ans: -1,
  },

  {
    grid: [
      [1, 0, 1],
      [0, 0, 0],
      [1, 0, 1],
    ],
    ans: 2,
  },

  {
    grid: [
      [1, 0, 0],
      [0, 0, 0],
      [0, 0, 0],
    ],
    ans: 4,
  },
];

const assert = require("assert");
const maxDistance = require("./far_from_land.js");

function testMaxDistance() {
  tcs.forEach(function (item, idx) {
    describe("TestProgram", function () {
      it(`test_case_${idx}`, function () {
        expect(item.ans).toEqual(maxDistance(item.grid));
      });
    });
  });
}

console.time("Time cost");

testMaxDistance();

console.timeEnd("Time cost");
