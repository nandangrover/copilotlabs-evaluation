const knuthMorrisPrattAlgorithm =  require('./knuth_morris_pratt_algorithm.js');

describe('nonAttackingQueens', () => {
  it('should return the number of non attacking queens', () => {
    expect(knuthMorrisPrattAlgorithm('aefoaefcdaefcdaed', 'aefcdaed')).toEqual(true);
  });
});