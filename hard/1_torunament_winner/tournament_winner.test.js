const  TournamentWinner = require('./tournament_winner');

describe('tournamentWinner', () => {
  it('should work', () => {
    const competitions = [
      ['HTML', 'C#'],
      ['C#', 'Python'],
      ['Python', 'HTML'],
    ];
    const results = [0, 0, 1];
    const expected = 'Python';
    const actual = TournamentWinner(competitions, results);
    expect(expected).toEqual(actual);
  });
});