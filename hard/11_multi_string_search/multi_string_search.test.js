const  multiStringSearch = require('./multi_string_search');

test("Test Case #1", () => {
  expect(multiStringSearch('this is a big string', ['this', 'yo', 'is', 'a', 'bigger', 'string', 'kappa'])).toEqual([true, false, true, true, false, true, false]);
});
