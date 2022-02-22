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


def print_dc():
    baseDir = "data-collection"
    df = pd.concat([pd.read_csv(f"{baseDir}/{f}") for f in os.listdir("data-collection")])

    for _, row in df.iterrows():
        row["Location"] = checkNan(row["Location"])

        if quickCheck(row['Instructor']):
            instructor = " ".join(reversed(row["Instructor"].split(",")))
            row['Instructor'] = instructor[1:]
        
        if quickCheck(row['Times']):
            times = row['Times'].split("-")
            row['Times'] = f"{convert24(times[0][:-1])}-{convert24(times[1][1:])}"

        print(f"{row['Course']},{row['Title']},{row['Days']},{row['Times']},{row['Location']},{row['Instructor']}")

def dc():
    baseDir = "data-collection"

    d = {}
    df = pd.concat([pd.read_csv(f"{baseDir}/{f}") for f in os.listdir("data-collection")])

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


        d[f"{row['Course']}:{row['Section']}"] = {
            "Title": row["Title"],
            "Days": row["Days"],
            "Times": row["Times"],
            "Location": row["Location"],
            "Instructor": row["Instructor"]            
        }

    return d

def main():
    d = dc()
    print(json.dumps(d, sort_keys=True))


if __name__ == "__main__":
    main()