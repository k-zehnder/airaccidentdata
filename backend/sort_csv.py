import pandas as pd

def sort_csv():
    df = pd.read_csv("./downloaded_file.csv")
    print(df.head())
    sorted_df = df.sort_values(by="ENTRY_DATE", ascending=False)
    sorted_df.to_csv("./sorted_downloaded_file.csv", index=False)

if __name__ == "__main__":
    sort_csv()
