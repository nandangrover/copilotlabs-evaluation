const HOME_TEAM_WON = 1;

// O(n) time | O(k) space - where n is the number
// of competitions and k is the number of teams
function TournamentWinner(competitions, results) {
	let currentBestTeam = "";
	const scores = { [currentBestTeam]: 0 };

	for (let idx = 0; idx < competitions.length; idx++) {
		const result = results[idx];
		const [homeTeam, awayTeam] = competitions[idx];

		let winningTeam = awayTeam;
		if (result === HOME_TEAM_WON) {
			winningTeam = homeTeam;
		}

		updateScores(winningTeam, 3, scores);

		if (scores[winningTeam] > scores[currentBestTeam]) {
			currentBestTeam = winningTeam;
		}
	}

	return currentBestTeam;
}

function updateScores(team, points, scores) {
	if (!(team in scores)) scores[team] = 0;
	scores[team] += points;
}

const competitions = [
	["HTML", "C#"],
	["C#", "Python"],
	["Python", "HTML"],
];
const results = [0, 0, 1];
console.log(TournamentWinner(competitions, results));
module.exports = TournamentWinner;
