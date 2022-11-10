# Check leap year

def is_leap_year(year):
    if year % 4 == 0 and year % 100 != 0 or year % 400 == 0:
        return True
    return False

print(is_leap_year(2000))  # True
print(is_leap_year(2001))  # False
print(is_leap_year(2004))  # True
print(is_leap_year(2100))  # False
