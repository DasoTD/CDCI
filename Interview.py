# from collections import Counter
def one(self):
    pass

def two(schedule: list[str], time):
    from datetime import datetime
    s1 = '12:30'
    s2 = '14:00' # for example
    format = '%H:%M'
    timesDiff = []
    times = datetime.strptime(s2, format) - datetime.strptime(s1, format)
    print(times)
    scheduleTime = datetime.strptime(schedule[0], format) - datetime.strptime(time, format)
    print(scheduleTime)
    scheduledTime = datetime.strptime(schedule[0], format)
    currentTime = datetime.strptime(time, format)
    if scheduledTime > currentTime:
        print(" e pass")
    else:
        print("e no pass")
    for i in range(len(schedule)):
        data = datetime.strptime(schedule[i], format) - datetime.strptime(time, format)
        timesDiff.append(data)
        # print(data, currentTime)
    print(timesDiff)
    # seconds = seconds % (24 * 3600)
    # if schedule[0] == time:
    #     print(-1)
    # passschedule[0] == time:
    #     print(-1)
    # pass

def three(n,m, figures):
    pass

def four(numbers: list[int], x):
    number = [str(x) for x in numbers]
    print(number)
    count = 0
    x = str(x)
    count =0
    for i in range(len(numbers)):
        for j in range( len(numbers)):
            if i!=j and str(numbers[i]) + str(numbers[j]) == x:
                count +=1
    return count
    # for i in range(len(number)):
    #     for j in range(i+1, len(number)):
    #         if i!=j and number[i] + number[j] == x:
    #             count +=1
    #         if number[j]+ number[i] == x:
    #             count += 1
    # return count
    

data = [1,212, 12, 12]
x = 1212
print(four(data, x))

# two(["12:00", "14:00", "19:55"], "14:30")