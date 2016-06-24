import sys

n = int(sys.argv[1])  # n >= 2

result = []
for k in range(n, 0, -1):
    for j in range(k, 0, -1):
        for i in range(j, 0, -1):
            result.append('10' * i + '0')
print(''.join(result))

# n = 2 * k, L(S_n) = 6 * k - 2, Z(S_n) = 5 * k - 3
# L(S_n) / Z(S_n) -> 6/5
