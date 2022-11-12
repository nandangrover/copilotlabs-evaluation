const  SuffixTrie = require('./suffix_trie_construction');

test("Test Case #1", () => {
  const trie = new SuffixTrie('babc');
  const expected = {
    c: {'*': true},
    b: {
      c: {'*': true},
      a: {b: {c: {'*': true}}},
    },
    a: {b: {c: {'*': true}}},
  };
  expect(trie.root).toEqual(expected);
  expect(trie.contains('abc')).toEqual(true);
});
