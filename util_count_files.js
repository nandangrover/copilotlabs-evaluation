/**
 * Utility function to count the number of files in the directory
 */
const { resolve } = require('path');
const { readdir } = require('fs').promises;

async function getFiles(dir) {
  const dirents = await readdir(dir, { withFileTypes: true });
  const files = await Promise.all(dirents.map((dirent) => {
    const res = resolve(dir, dirent.name);
    return dirent.isDirectory() && !res.match(/node_modules/) ? getFiles(res) : res;
  }));
  return files.flat().filter((file) => file.endsWith('.js') || file.endsWith('.test.js') || file.endsWith('.go') || file.endsWith('.py') || file.endsWith('.txt'));
}
getFiles('../copilotlabs-evaluation').then((files) => console.log(files.length));
