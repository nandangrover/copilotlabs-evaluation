import pandas as pd

def pandas_operation(ds1, ds2):
    add = ds1 + ds2

    sub = ds1 - ds2

    mul = ds1 * ds2

    div = ds1 / ds2

    return list(map(lambda x: x.tolist(), [add, sub, mul, div]))
    
ds1 = pd.Series([2, 4, 6, 8, 10])
ds2 = pd.Series([1, 3, 5, 7, 9])
print(pandas_operation(ds1, ds2))