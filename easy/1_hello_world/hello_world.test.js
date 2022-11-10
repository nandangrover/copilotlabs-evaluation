const helloworld = require("./hello_world");

test('console.log the text "Hello World"', () => {
    console.log = jest.fn();
    helloworld();
    expect(console.log.mock.calls[0][0]).toBe('Hello World');
  });