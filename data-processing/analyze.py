"""
Takes a hidden data-collection dir and analyzes it for stuff in DB schema.

PIP Requirements: pandas
"""

import pandas as pd
import os

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

def main():
    baseDir = "data-collection"
    df = pd.concat([pd.read_csv(f"{baseDir}/{f}") for f in os.listdir("data-collection")])

    for _, row in df.iterrows():
        if row['Location'] == " " or row['Location'] == "":
            row['Location'] = "nan" 

        if quickCheck(row['Instructor']):
            instructor = " ".join(reversed(row["Instructor"].split(",")))
            row['Instructor'] = instructor[1:]
        
        if quickCheck(row['Times']):
            times = row['Times'].split("-")
            row['Times'] = f"{convert24(times[0][:-1])}-{convert24(times[1][1:])}"

        # if quickCheck(row['Location']):
        #     print(row['Location'])
        print(f"{row['Course']},{row['Title']},{row['Days']},{row['Times']},{row['Location']},{row['Instructor']}")
    
    # print(df.to_string()) # complete dataframe thing

if __name__ == "__main__":
    main()