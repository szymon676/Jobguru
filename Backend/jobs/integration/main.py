import requests
import time

API_URL = "http://localhost:5001/jobs"


def send_request(method, url, payload=None):
    try:
        resp = requests.request(method, url, json=payload)

        if resp.ok:
            print(f"{method} request successful")
        else:
            print(f"wrong status, expected {resp.status_code} got {resp.status_code}")

    except requests.exceptions.RequestException as e:
        print(f"something went wrong: {e}")


def main():
    payload = {
        "user_id": 1,
        "title": "Cobol developer for bank",
        "company": "bank sa",
        "skills": ["cobol"],
        "salary": 30000,
        "description": "cobol developer to develop perfoment systems",
        "currency": "USD",
        "date": "2022-12-12",
        "location": "manchaster",
    }
    send_request("POST", API_URL, payload)

    send_request("GET", API_URL)

    payload = {
        "user_id": 1,
        "title": "Typescript developer",
        "company": "ASDAS!@#!@#@!",
        "skills": ["typescript", "python"],
        "salary": 20000,
        "description": "cobol developer to develop perfoment systems",
        "currency": "zł",
        "date": "2023-12-12",
        "location": "Warsaw",
    }
    send_request("PUT", f"{API_URL}/8", payload)

    send_request("DELETE", f"{API_URL}/9")

if __name__ == "__main__":
    main()
