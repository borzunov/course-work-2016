import sys

n = int(sys.argv[1])  # n >= 2

result = []
for m in range(n + 1, 1, -1):
    for k in range(m + 1, 1, -1):
        for j in range(k + 1, 1, -1):
            for i in range(j + 1, 1, -1):
                result.append('10' * i + '0')
print(''.join(result))

# n = 2 * k, L(S_n) = 8 * k + 5, Z(S_n) = 7 * k + 2
# L(S_n) / Z(S_n) -> 8/7
