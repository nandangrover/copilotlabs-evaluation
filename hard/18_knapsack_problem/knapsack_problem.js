function knapsackProblem(items, capacity) {
  let knapsackValues = new Array(items.length + 1).fill().map(() => new Array(capacity + 1).fill(0));
  for (let i = 1; i < items.length + 1; i++) {
      let currentWeight = items[i - 1][1];
      let currentValue = items[i - 1][0];
      for (let c = 0; c < capacity + 1; c++) {
          if (currentWeight > c) {
              knapsackValues[i][c] = knapsackValues[i - 1][c];
          } else {
              knapsackValues[i][c] = Math.max(
                  knapsackValues[i - 1][c], knapsackValues[i - 1][c - currentWeight] + currentValue
              );
          }
      }
  }
  return [knapsackValues[items.length][capacity], getKnapsackItems(knapsackValues, items)];
}

function getKnapsackItems(knapsackValues, items) {
  let sequence = [];
  let i = knapsackValues.length - 1;
  let c = knapsackValues[0].length - 1;
  while (i > 0) {
      if (knapsackValues[i][c] === knapsackValues[i - 1][c]) {
          i -= 1;
      } else {
          sequence.push(i - 1);
          c -= items[i - 1][1];
          i -= 1;
      }
      if (c === 0) {
          break;
      }
  }
  return sequence.reverse();
}

console.log(knapsackProblem([[1, 2], [4, 3], [5, 6], [6, 7]], 10));

module.exports = knapsackProblem;
