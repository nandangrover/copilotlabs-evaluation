import json
import datetime


def clean(input_file, output_file):
    output = {"date": "", "base_currency": "EUR",
              "target_currency": "", "exchange_rate": 0.0}

    with open(input_file) as f:
        currency_data = json.load(f)

    with open(output_file, 'w') as f:
        for date in currency_data:
            for key in currency_data[date].keys():
                output['date'] = date
                output['target_currency'] = key
                output['exchange_rate'] = currency_data[date][key]
                json.dump(output, f)
                f.write('\n')


def curate(input_file, output_file):
    currency_data = {}

    with open(input_file) as f:
        raw_input_data = list(f)

    for json_str in raw_input_data:
        data = json.loads(json_str)

        if data['target_currency'] not in currency_data.keys():
            currency_data[data['target_currency']] = []
        currency_data[data['target_currency']].append(data)

    for currency in currency_data.keys():
        currency_data[currency].sort(
            key=lambda json: datetime.datetime.fromisoformat(json['date']))
        last_exchange_rate = currency_data[currency][0]['exchange_rate']
        last_date = currency_data[currency][0]['date']

        for index in range(1, len(currency_data[currency])):
            data = currency_data[currency][index]

            if next_date_exist(last_date, currency_data[currency]):
                last_exchange_rate = data['exchange_rate']
                last_date = data['date']
            else:
                output = {"date": "", "base_currency": "EUR",
                          "target_currency": "", "exchange_rate": 0.0}
                output['date'] = (datetime.datetime.fromisoformat(
                    last_date) + datetime.timedelta(days=1)).strftime('%Y-%m-%d')
                output['target_currency'] = currency
                output['exchange_rate'] = last_exchange_rate

                currency_data[currency].append(output)

                last_date = output['date']

    with open(output_file, 'w') as f:
        for currency in currency_data:
            for data in currency_data[currency]:
                json.dump(data, f)
                f.write('\n')

def next_date_exist(date, data):
    next_date = (datetime.datetime.fromisoformat(date) +
                 datetime.timedelta(days=1)).strftime('%Y-%m-%d')
    if next((item for item in data if item["date"] == next_date), False):
        return True
    else:
        return False
