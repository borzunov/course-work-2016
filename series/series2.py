import sys

n = int(sys.argv[1])  # n >= 2

result = []
for k in range(n + 1, 1, -1):
    for i in range(k + 1, 1, -1):
        result.append('10' * i + '0')
print(''.join(result))

# n = 2 * k, L(S_n) = 4 * k + 2, Z(S_n) = 3 * k + 2
# L(S_n) / Z(S_n) -> 4/3
