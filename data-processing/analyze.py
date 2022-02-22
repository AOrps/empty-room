"""
Takes a hidden data-collection dir and analyzes it for stuff in DB schema.

PIP Requirements: pandas
"""

import pandas as pd
import os
import json

def convert24(time: str) -> str:
    setOffset = False
    
    meridian = time[-2:]
    if meridian == "PM":
        setOffset = True
    
    time_breakdown = time[:-2].split(":")
    hour, minute = int(time_breakdown[0]), int(time_breakdown[1])

    if setOffset and hour != 12:
        hour += 12
    
    if minute == 0:
        minute = "00"
    return f"{hour}:{minute}"

quickCheck = lambda val: str(val).lower() != "nan"
checkNan = lambda x: "nan" if x == " " or x == "" or str(x).lower() == "nan" else x

def iter1(df):
    d = {}
    i = 1
    for _, row in df.iterrows():
        row["Days"] = checkNan(row["Days"])
        row["Instructor"] = checkNan(row["Instructor"])
        row["Location"] = checkNan(row["Location"])
        row["Times"] = checkNan(row["Times"]) 

        if quickCheck(row['Instructor']):
            instructor = " ".join(reversed(row["Instructor"].split(",")))
            row['Instructor'] = instructor[1:]
        
        if quickCheck(row['Times']):
            times = row['Times'].split("-")
            row['Times'] = f"{convert24(times[0][:-1])}-{convert24(times[1][1:])}"

        d[f"{i}"] = {
            "id": f"{i}",
            "title": row["Title"],
            "days": row["Days"],
            "times": row["Times"],
            "location": row["Location"],
            "instructor": row["Instructor"],
            "course": row["Course"],
            "section": row["Section"]
        }
        i += 1
    return d 

def iter2(df):
    d = {}
    for _, row in df.iterrows():
        row["Days"] = checkNan(row["Days"])
        row["Instructor"] = checkNan(row["Instructor"])
        row["Location"] = checkNan(row["Location"])
        row["Times"] = checkNan(row["Times"]) 

        loc = row['Location']
        if loc != "nan":
            base_loc = loc.split(" ")
            if quickCheck(row['Instructor']):
                instructor = " ".join(reversed(row["Instructor"].split(",")))
                row['Instructor'] = instructor[1:]
            
            if quickCheck(row['Times']):
                times = row['Times'].split("-")
                row['Times'] = f"{convert24(times[0][:-1])}-{convert24(times[1][1:])}"

            key, title, days, time = f"{row['Course']}:{row['Section']}", row["Title"], row["Days"], row["Times"]
            instructor = row['Instructor']
            building, room = base_loc[0], "".join(base_loc[1:])
            
            if building not in d:
                d[building] = [(room, key, title, days, time, instructor)]
            else:
                d[building].append((room, key, title, days, time, instructor)) 

    return d

def dc():
    baseDir = "data-collection"

    d = {}
    df = pd.concat([pd.read_csv(f"{baseDir}/{f}") for f in os.listdir("data-collection")])

    d = iter2(df)

    return d


def main():
    d = dc()
    print(json.dumps(d, sort_keys=False))


if __name__ == "__main__":
    main()