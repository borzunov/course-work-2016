import sys

n = int(sys.argv[1])  # n >= 2

result = []
for k in range(n + 1, 1, -1):
    for j in range(k + 1, 1, -1):
        for i in range(j + 1, 1, -1):
            result.append('10' * i + '0')
print(''.join(result))
