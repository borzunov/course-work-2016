import sys

n = int(sys.argv[1])  # n >= 2

result = []
for i in range(n + 1, 1, -1):  # For i = n + 1 .. 2 (inclusive)
    result.append('10' * i + '0')
result.append('10')
print(''.join(result))

# n = 2 * k, L(S_n) = 2 * k + 3, Z(S_n) = k + 3
# L(S_n) / Z(S_n) -> 2
