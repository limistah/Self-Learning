def euclidean(a, b):
    if a == 0 | b == 0:
        return 0
    remainder = a % b
    print((a, b, remainder))
    if remainder == 0:
        return b
    return euclidean(b, remainder)


print(euclidean(270, 92))
