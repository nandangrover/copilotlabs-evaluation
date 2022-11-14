function CalculateValue(values) {
	value = Math.random() * 10;
	console.log("Calculated Random Value: ", value);
	values.push(value);
}

function main() {
	console.log("Go Channel Tutorial");

	values = [];

	CalculateValue(values);

	value = values.pop();
	console.log(value);
}

main();