#!/usr/bin/env python3
"""Verify an exact Telnyx Go tag and public module version are available."""
from __future__ import annotations

import argparse
import json
import os
import re
import subprocess
import sys
import time
import urllib.error
import urllib.request
from typing import Mapping, Optional, Sequence

VERSION_RE = re.compile(r"^[0-9]+\.[0-9]+\.[0-9]+$")
SHA_RE = re.compile(r"^[0-9a-f]{40}$")
MODULE = "github.com/team-telnyx/telnyx-go/v4"


class AvailabilityError(RuntimeError):
    pass


def validate_tag(payload: Mapping[str, object], tag: str, release_sha: str) -> None:
    obj = payload.get("object")
    if payload.get("ref") != "refs/tags/" + tag or not isinstance(obj, Mapping):
        raise AvailabilityError("tag reference is malformed")
    if obj.get("type") == "tag":
        raise AvailabilityError("annotated tag must be dereferenced before validation")
    if obj.get("type") != "commit" or obj.get("sha") != release_sha:
        raise AvailabilityError("tag does not resolve to the exact release commit")


def validate_proxy_info(payload: Mapping[str, object], tag: str) -> None:
    if payload.get("Version") != tag:
        raise AvailabilityError("Go proxy version does not match exact release")
    timestamp = payload.get("Time")
    if not isinstance(timestamp, str) or not timestamp.endswith("Z"):
        raise AvailabilityError("Go proxy timestamp is missing")


def validate_module_file(text: str, module: str) -> None:
    first = text.splitlines()[0].strip() if text.splitlines() else ""
    if first != "module " + module:
        raise AvailabilityError("Go module path does not match expected module")


def validate_go_list(payload: Mapping[str, object], module: str, tag: str) -> None:
    if payload.get("Path") != module:
        raise AvailabilityError("go list module path does not match expected module")
    if payload.get("Version") != tag:
        raise AvailabilityError("go list version does not match exact release")


def fetch_json(url: str, token: str = "") -> Mapping[str, object]:
    headers = {"User-Agent": "telnyx-go-release-readiness/1"}
    if token:
        headers["Authorization"] = "Bearer " + token
        headers["Accept"] = "application/vnd.github+json"
    with urllib.request.urlopen(urllib.request.Request(url, headers=headers), timeout=30) as response:
        payload = json.load(response)
    if not isinstance(payload, Mapping):
        raise AvailabilityError("invalid JSON response")
    return payload


def resolve_tag(repository: str, tag: str, token: str) -> Mapping[str, object]:
    payload = fetch_json("https://api.github.com/repos/%s/git/ref/tags/%s" % (repository, tag), token)
    for _ in range(5):
        obj = payload.get("object")
        if isinstance(obj, Mapping) and obj.get("type") == "commit":
            return payload
        if not isinstance(obj, Mapping) or obj.get("type") != "tag" or not SHA_RE.fullmatch(str(obj.get("sha", ""))):
            raise AvailabilityError("tag cannot be dereferenced")
        tag_object = fetch_json("https://api.github.com/repos/%s/git/tags/%s" % (repository, obj["sha"]), token)
        payload = {"ref": "refs/tags/" + tag, "object": tag_object.get("object")}
    raise AvailabilityError("tag dereference depth exceeded")


def fetch_text(url: str) -> str:
    request = urllib.request.Request(url, headers={"User-Agent": "telnyx-go-release-readiness/1"})
    with urllib.request.urlopen(request, timeout=30) as response:
        return response.read().decode("utf-8")


def module_resolution(module: str, tag: str) -> Mapping[str, object]:
    completed = subprocess.run(
        ["go", "list", "-m", "-json", module + "@" + tag],
        text=True, capture_output=True,
        env={**os.environ, "GOPROXY": "https://proxy.golang.org", "GONOSUMDB": ""},
    )
    if completed.returncode != 0:
        raise AvailabilityError("go list could not resolve exact public module version")
    payload = json.loads(completed.stdout)
    if not isinstance(payload, Mapping):
        raise AvailabilityError("go list returned invalid metadata")
    return payload


def verify(repository: str, version: str, release_sha: str) -> None:
    if not VERSION_RE.fullmatch(version) or not SHA_RE.fullmatch(release_sha):
        raise AvailabilityError("invalid exact release coordinates")
    tag = "v" + version
    token = os.environ.get("GH_TOKEN") or os.environ.get("GITHUB_TOKEN") or ""
    if not token:
        raise AvailabilityError("GitHub token is required")
    validate_tag(resolve_tag(repository, tag, token), tag, release_sha)
    escaped = MODULE.replace("!", "!!")
    base = "https://proxy.golang.org/%s/@v/%s" % (escaped, tag)
    validate_proxy_info(fetch_json(base + ".info"), tag)
    validate_module_file(fetch_text(base + ".mod"), MODULE)
    validate_go_list(module_resolution(MODULE, tag), MODULE, tag)
    # Fetching the immutable module archive proves the distributable payload exists.
    with urllib.request.urlopen(urllib.request.Request(base + ".zip", method="HEAD"), timeout=30) as response:
        if response.status != 200:
            raise AvailabilityError("Go module archive is unavailable")


def main(argv: Optional[Sequence[str]] = None) -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repository", default="team-telnyx/telnyx-go")
    parser.add_argument("--version", required=True)
    parser.add_argument("--release-sha", required=True)
    parser.add_argument("--attempts", type=int, default=20)
    parser.add_argument("--delay", type=int, default=15)
    args = parser.parse_args(argv)
    if args.attempts < 1 or args.delay < 0:
        return 2
    last = "not found"
    for attempt in range(args.attempts):
        try:
            verify(args.repository, args.version, args.release_sha)
            print("verified Go module %s@v%s at %s" % (MODULE, args.version, args.release_sha))
            return 0
        except (AvailabilityError, urllib.error.URLError, TimeoutError, ValueError, json.JSONDecodeError) as exc:
            last = str(exc)
            if attempt + 1 < args.attempts:
                time.sleep(args.delay)
    print("Go release availability failed: %s" % last, file=sys.stderr)
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
