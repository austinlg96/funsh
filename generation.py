import random


print("Level 5")
for i in range(0,10):
    x = int(f'{random.randint(0,9)}{random.randint(0,5)}')
    y = int(f'{random.randint(0,5)}{random.randint(0,5)}')
    
    output  = "{" + f'"{x}*{y}","{x*y}", 12345' + "},"
    print(output)

