function knuthMorrisPrattAlgorithm(string, substring) {
    const pattern = buildPattern(substring);
    return doesMatch(string, substring, pattern);
}


function buildPattern(substring) {
    const pattern = new Array(substring.length).fill(-1);
    let j = 0;
    let i = 1;
    while (i < substring.length) {
        if (substring[i] === substring[j]) {
            pattern[i] = j;
            i += 1;
            j += 1;
        } else if (j > 0) {
            j = pattern[j - 1] + 1;
        } else {
            i += 1;
        }
    }
    return pattern;
}


function doesMatch(string, substring, pattern) {
    let i = 0;
    let j = 0;
    while (i + substring.length - j <= string.length) {
        if (string[i] === substring[j]) {
            if (j === substring.length - 1) {
                return true;
            }
            i += 1;
            j += 1;
        } else if (j > 0) {
            j = pattern[j - 1] + 1;
        } else {
            i += 1;
        }
    }
    return false;
}

console.log(knuthMorrisPrattAlgorithm('aefoaefcdaefcdaed', 'aefcdaed'));
console.log(knuthMorrisPrattAlgorithm('aefoaefceaefcdaet', 'aefcdaed'));

module.exports = knuthMorrisPrattAlgorithm;