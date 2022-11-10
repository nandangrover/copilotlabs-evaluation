const assert = require('assert');
const pyramidPattern = require('./number_pyramid_pattern.js');

test('Check if pattern is correct', () => {
    const logSpy = jest.spyOn(console._stdout, 'write')
    pyramidPattern(5);
    const sanitizedOutput = logSpy.mock.calls.toString().replace(/,/g, '').trim();
    assert.equal(sanitizedOutput, "1\n   22\n  333\n 4444\n55555");
  });