import sys

n = int(sys.argv[1])  # n >= 2

result = []
for i in range(n + 1, 1, -1):
    result.append('10' * i + '0')
result.append('10')
print(''.join(result))
