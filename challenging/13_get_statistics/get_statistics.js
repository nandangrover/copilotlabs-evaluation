function get_statistics(input_list) {
    let sorted_input = input_list.sort();
    let input_length = sorted_input.length;

    let mean = sorted_input.reduce((a, b) => a + b) / input_length;

    let middle_idx = (input_length - 1) / 2;
    let median = sorted_input[middle_idx];

    if (input_length % 2 === 0) {
        let middle_number_1 = sorted_input[middle_idx];
        let middle_number_2 = sorted_input[middle_idx + 1];
        median = (middle_number_1 + middle_number_2) / 2;
    }

    let number_counts = {};
    for (let x of new Set(sorted_input)) {
        number_counts[x] = sorted_input.filter((number) => number === x).length;
    }
    let mode = Object.keys(number_counts).reduce((unique_number, key) => number_counts[key] > number_counts[unique_number] ? key : unique_number);

    let sample_variance = sorted_input.reduce((accumulator, number) => accumulator + (number - mean) ** 2 / (input_length - 1), 0);

    let sample_standard_deviation = sample_variance ** 0.5;

    let mean_standard_error = sample_standard_deviation / input_length ** 0.5;
    let z_score_standard_error = 1.96 * mean_standard_error;
    let mean_confidence_interval = [mean - z_score_standard_error, mean + z_score_standard_error];

    return {
        "mean": mean,
        "median": median,
        "mode": mode,
        "sample_variance": sample_variance,
        "sample_standard_deviation": sample_standard_deviation,
        "mean_confidence_interval": mean_confidence_interval,
    };
}

console.log(get_statistics([2, 1, 3, 4, 4, 5, 6, 7]));

module.exports = get_statistics;