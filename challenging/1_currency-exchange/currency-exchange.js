function clean(input_file, output_file) {
    var output = {"date": "", "base_currency": "EUR",
                  "target_currency": "", "exchange_rate": 0.0};

    var currency_data = json.load(f);

    for (var date in currency_data) {
        for (var key in currency_data[date].keys()) {
            output['date'] = date;
            output['target_currency'] = key;
            output['exchange_rate'] = currency_data[date][key];
            json.dump(output, f);
            f.write('\n');
        }
    }
}


function curate(input_file, output_file) {
    var currency_data = {};

    var raw_input_data = list(f);

    for (var json_str in raw_input_data) {
        var data = json.loads(json_str);

        if (data['target_currency'] not in currency_data.keys()) {
            currency_data[data['target_currency']] = [];
        }
        currency_data[data['target_currency']].append(data);
    }

    for (var currency in currency_data.keys()) {
        currency_data[currency].sort(
            key=lambda json: datetime.datetime.fromisoformat(json['date']));
        var last_exchange_rate = currency_data[currency][0]['exchange_rate'];
        var last_date = currency_data[currency][0]['date'];

        for (var index in range(1, len(currency_data[currency]))) {
            var data = currency_data[currency][index];

            if (next_date_exist(last_date, currency_data[currency])) {
                last_exchange_rate = data['exchange_rate'];
                last_date = data['date'];
            } else {
                var output = {"date": "", "base_currency": "EUR",
                              "target_currency": "", "exchange_rate": 0.0};
                output['date'] = (datetime.datetime.fromisoformat(
                    last_date) + datetime.timedelta(days=1)).strftime('%Y-%m-%d');
                output['target_currency'] = currency;
                output['exchange_rate'] = last_exchange_rate;

                currency_data[currency].append(output);

                last_date = output['date'];
            }
        }
    }

    for (var currency in currency_data) {
        for (var data in currency_data[currency]) {
            json.dump(data, f);
            f.write('\n');
        }
    }
}

function next_date_exist(date, data) {
    var next_date = (datetime.datetime.fromisoformat(date) +
                     datetime.timedelta(days=1)).strftime('%Y-%m-%d');
    if (next((item for item in data if item["date"] == next_date), false)) {
        return true;
    } else {
        return false;
    }
}