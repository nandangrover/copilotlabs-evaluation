def LongestBalancedSubstring(str):
    return max(
        getLongestBalancedInDirection(str, True),
        getLongestBalancedInDirection(str, False),
    )


def getLongestBalancedInDirection(str, leftToRight):
    openingParens = '('
    startIdx = 0
    step = 1
    if not leftToRight:
        openingParens = ')'
        startIdx = len(str) - 1
        step = -1

    maxLength = 0

    openingCount = 0
    closingCount = 0

    idx = startIdx
    while idx >= 0 and idx < len(str):
        char = str[idx]

        if char == openingParens:
            openingCount += 1
        else:
            closingCount += 1
        if openingCount == closingCount:
            maxLength = max(maxLength, closingCount * 2)
        elif closingCount > openingCount:
            openingCount = 0
            closingCount = 0

        idx += step

    return maxLength


def max(a, b):
    if a > b:
        return a
    return b


if __name__ == "__main__":
    input = "(()))("
    print(LongestBalancedSubstring(input))  # 4