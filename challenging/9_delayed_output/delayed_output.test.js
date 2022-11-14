const getFiles = require('./count_files');

test('Check if correct file count', async () => {
    const files = await getFiles('./');
    // files.then((files) => console.log(files.length));
    // const sanitizedOutput = logSpy.mock.calls.toString().replace(/,/g, '').trim();
    expect(files.length).toBe(323);
  });