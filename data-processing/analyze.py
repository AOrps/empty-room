"""
Takes a hidden data-collection dir and analyzes it for stuff in DB schema.

PIP Requirements: pandas
"""

import pandas as pd
import os


def main():
    baseDir = "data-collection"
    # df = pd.read_csv(f"{baseDir}/ACCT.csv")
    # pd.read_csv(f"{baseDir}/AD.csv")
    # print([f"{baseDir}/{f}" for f in os.listdir("data-collection")])
    df = pd.concat([pd.read_csv(f"{baseDir}/{f}") for f in os.listdir("data-collection")])

    for _, row in df.iterrows():
        if row['Location'] == " " or row['Location'] == "":
            row['Location'] = "nan" 

        if row['Instructor'] != "nan":
            instructor = "".join(row["Instructor"].split(","))
            row['Instructor'] = instructor
        print(f"{row['Course']},{row['Title']},{row['Days']},{row['Times']},{row['Location']},{row['Instructor']}")

    # print(df.to_string())

if __name__ == "__main__":
    main()