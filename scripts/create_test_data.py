#!/usr/bin/env python3
"""
Creates test data against the OGCR OBP API:
  - ensures operator exists (creates if not)
  - ensures activity exists (creates if not, linked to operator)
  - COUNT parcels
  - COUNT parcel_owner_verification (status_code = in_progress)
  - COUNT activity_parcel_verification (linked to activity)
"""

import os
import sys
import json
import uuid
import random
import requests
from dotenv import load_dotenv

load_dotenv()  # reads .env without shell interpretation; existing env vars win

BASE_URL     = os.environ["OBP_URL"]
USERNAME     = os.environ["OBP_USERNAME"]
PASSWORD     = os.environ["OBP_PASSWORD"]
CONSUMER_KEY = os.environ["OBP_CONSUMER_KEY"]
ACTIVITY_ID          = os.environ.get("ACTIVITY_ID",  "a_05OFWI037022")
OPERATOR_ID          = os.environ.get("OPERATOR_ID",  "op_TESOBE01")
OGCR_WALLET_ADDRESS  = os.environ.get("OGCR_WALLET_ADDRESS")  # required only if operator must be created
COUNT                = int(os.environ.get("COUNT", "20"))

# ─── Auth ────────────────────────────────────────────────────────────────────

def get_token():
    resp = requests.post(
        f"{BASE_URL}/my/logins/direct",
        headers={
            "Authorization": f'DirectLogin username="{USERNAME}",password="{PASSWORD}",consumer_key="{CONSUMER_KEY}"',
            "Content-Type": "application/json",
        },
    )
    resp.raise_for_status()
    token = resp.json().get("token")
    if not token:
        sys.exit(f"No token in response: {resp.text}")
    print(f"Authenticated. Token: {token[:12]}...")
    return token


def auth_headers(token):
    return {
        "Authorization": f"DirectLogin token={token}",
        "Content-Type": "application/json",
    }


# ─── Helpers ─────────────────────────────────────────────────────────────────

def get(token, path):
    """GET a single dynamic entity by path; returns parsed JSON or None on 404."""
    url = f"{BASE_URL}/obp/dynamic-entity/{path}"
    resp = requests.get(url, headers=auth_headers(token))
    if resp.status_code == 404:
        return None
    if not resp.ok:
        print(f"  GET ERROR {resp.status_code}: {resp.text[:200]}")
        return None
    return resp.json()


def post(token, path, body):
    url = f"{BASE_URL}/obp/dynamic-entity/{path}"
    resp = requests.post(url, headers=auth_headers(token), json=body)
    if not resp.ok:
        print(f"  POST ERROR {resp.status_code}: {resp.text[:200]}")
        return None
    return resp.json()


def make_parcel_id():
    return "parcel_" + uuid.uuid4().hex[:10].upper()


def make_polygon(base_lon=14.38, base_lat=52.52, size=0.01):
    lon = base_lon + random.uniform(-0.5, 0.5)
    lat = base_lat + random.uniform(-0.5, 0.5)
    return {
        "type": "MultiPolygon",
        "coordinates": [[[[
            [lon,        lat],
            [lon + size, lat],
            [lon + size, lat + size],
            [lon,        lat + size],
            [lon,        lat],
        ]]]],
    }


# ─── Ensure operator & activity exist ────────────────────────────────────────

def ensure_operator(token, operator_id, ogcr_wallet_address):
    print(f"\n── Checking operator {operator_id} ──")
    existing = get(token, f"operator/{operator_id}")
    if existing:
        print(f"  Operator already exists.")
        return

    if not ogcr_wallet_address:
        sys.exit("OGCR_WALLET_ADDRESS is required to create a new operator — aborting.")

    body = {
        "operator_id": operator_id,
        "legal_name": "TESOBE GmbH (test)",
        "email": "test@tesobe.com",
        "phone": "+49000000000",
        "address_line_1": "Potsdamer Str. 1",
        "address_line_2": "",
        "postcode": "10785",
        "country_code": "DE",
        "ogcr_wallet_address": ogcr_wallet_address,
    }
    result = post(token, "operator", body)
    if result:
        print(f"  Operator created: {operator_id}")
    else:
        sys.exit(f"Failed to create operator {operator_id} — aborting.")


def ensure_activity(token, activity_id, operator_id):
    print(f"\n── Checking activity {activity_id} ──")
    existing = get(token, f"activity/{activity_id}")
    if existing:
        print(f"  Activity already exists.")
        return

    body = {
        "activity_id": activity_id,
        "operator_id": operator_id,
        "name": "Test Carbon Farming Activity",
        "summary": "Created by create_test_data.py",
        "description": "Automated test activity for tokenizer development.",
        "type": "carbon_farming",
        "city": "Berlin",
        "country_code": "DE",
        "multipolygon_coordinates": make_polygon(),
        "website": "https://ogcr.eu",
        "image": "",
        "media_links": "",
        "technologies_practices_processes": "regenerative agriculture",
        "start_date": "2024-01-01",
        "end_date": "2034-01-01",
        "term_commitment": 10,
        "cobenefits": "biodiversity",
        "methodologies": "SALUS v1.0",
        "monitoring_period_years": 1,
        "monitoring_period_start_date": "2024-01-01",
        "monitoring_period_end_date": "2025-01-01",
    }
    result = post(token, "activity", body)
    if result:
        print(f"  Activity created: {activity_id}")
    else:
        sys.exit(f"Failed to create activity {activity_id} — aborting.")


# ─── Creation functions ───────────────────────────────────────────────────────

def create_parcels(token, count):
    print(f"\n── Creating {count} parcels ──")
    parcel_ids = []
    for i in range(count):
        pid = make_parcel_id()
        body = {
            "parcel_id": pid,
            "multipolygon_coordinates_geometry": make_polygon(),
            "crs": "EPSG:4326",
        }
        result = post(token, "parcel", body)
        if result:
            print(f"  [{i+1:02d}] parcel created: {pid}")
            parcel_ids.append(pid)
        else:
            print(f"  [{i+1:02d}] failed for {pid}")
    return parcel_ids


def create_parcel_owner_verifications(token, parcel_ids):
    print(f"\n── Creating {len(parcel_ids)} parcel_owner_verifications (in_progress) ──")
    for i, pid in enumerate(parcel_ids):
        body = {
            "parcel_id": pid,
            "status_code": "in_progress",
            "status_message": "Ownership verification pending",
            "authority": "OGCR",
        }
        result = post(token, "parcel_owner_verification", body)
        if result:
            print(f"  [{i+1:02d}] parcel_owner_verification created for {pid}")
        else:
            print(f"  [{i+1:02d}] failed for {pid}")


def create_activity_parcel_verifications(token, parcel_ids, activity_id):
    print(f"\n── Creating {len(parcel_ids)} activity_parcel_verifications (activity={activity_id}) ──")
    for i, pid in enumerate(parcel_ids):
        body = {
            "parcel_id": pid,
            "activity_id": activity_id,
            "status_code": "in_progress",
            "status_message": "Parcel verification pending",
            "amount": random.randint(1, 100),
        }
        result = post(token, "activity_parcel_verification", body)
        if result:
            print(f"  [{i+1:02d}] activity_parcel_verification created for {pid}")
        else:
            print(f"  [{i+1:02d}] failed for {pid}")


# ─── Main ─────────────────────────────────────────────────────────────────────

if __name__ == "__main__":
    token = get_token()

    ensure_operator(token, OPERATOR_ID, OGCR_WALLET_ADDRESS)
    ensure_activity(token, ACTIVITY_ID, OPERATOR_ID)

    parcel_ids = create_parcels(token, COUNT)
    if not parcel_ids:
        sys.exit("No parcels created — aborting.")

    create_parcel_owner_verifications(token, parcel_ids)
    create_activity_parcel_verifications(token, parcel_ids, ACTIVITY_ID)

    print(f"\nDone. Created {len(parcel_ids)} parcels and associated records.")
    print("Parcel IDs:", json.dumps(parcel_ids, indent=2))
