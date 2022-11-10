def pyramidPattern(n):
    for i in range(1, n + 1):
        for j in range(0, n - i):
            print(" ", end="")
        for k in range(0, i):
            print(i, end="")
        print()


pyramidPattern(5)