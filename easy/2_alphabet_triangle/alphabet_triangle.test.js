const alphabetTriangle = require("./alphabet_triangle");

test('Check if pattern is correct', () => {
    const logSpy = jest.spyOn(console._stdout, 'write')
    alphabetTriangle();
    const sanitizedOutput = logSpy.mock.calls.toString().replace(/,/g, '').trim();
    expect(sanitizedOutput).toBe("A \nA B \nA B C \nA B C D \nA B C D E \nA B C D E F \nA B C D E F G");
  });