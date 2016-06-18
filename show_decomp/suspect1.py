import sys

n = int(sys.argv[1])  # n >= 2

print(''.join('10' * i + '0' for i in range(n + 1, 1, -1)) + '10')
