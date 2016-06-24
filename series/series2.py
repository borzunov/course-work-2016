import sys

n = int(sys.argv[1])  # n >= 2

result = []
for k in range(n, 0, -1):
    for i in range(k, 0, -1):
        result.append('10' * i + '0')
print(''.join(result))

# n = 2 * k, L(S_n) = 4 * k, Z(S_n) = 3 * k
# L(S_n) / Z(S_n) -> 4/3
