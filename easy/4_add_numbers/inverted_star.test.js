const invertedStar = require("./inverted_star");

test('Check if pattern is correct', () => {
    const logSpy = jest.spyOn(console._stdout, 'write')
    invertedStar(5);
    const sanitizedOutput = logSpy.mock.calls.toString().replace(/,/g, '').trim();
    expect(sanitizedOutput).toBe( "*\n   **\n  ***\n ****\n*****");
  });