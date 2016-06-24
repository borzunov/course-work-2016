import sys

n = int(sys.argv[1])  # n >= 2

result = []
for m in range(n, 0, -1):
    for k in range(m, 0, -1):
        for j in range(k, 0, -1):
            for i in range(j, 0, -1):
                    result.append('10' * i + '0')
print(''.join(result))

# n = 2 * k, L(S_n) = 8 * k - 4, Z(S_n) = 7 * k - 6
# L(S_n) / Z(S_n) -> 8/7
