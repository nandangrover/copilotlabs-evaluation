from typing import List


def smallest_difference(array1: List[int], array2: List[int]) -> List[int]:
    array1.sort()
    array2.sort()
    idx_one = 0
    idx_two = 0
    smallest = float("inf")
    current = float("inf")
    smallest_pair = []
    while idx_one < len(array1) and idx_two < len(array2):
        first = array1[idx_one]
        second = array2[idx_two]
        if first < second:
            current = second - first
            idx_one += 1
        elif second < first:
            current = first - second
            idx_two += 1
        else:
            return [first, second]
        if smallest > current:
            smallest = current
            smallest_pair = [first, second]
    return smallest_pair


if __name__ == "__main__":
    array1 = [-1, 5, 10, 20, 28, 3]
    array2 = [26, 134, 135, 15, 17]
    print(smallest_difference(array1, array2))