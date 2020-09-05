#!/usr/bin/env bats

@test "wtz" {
    run wtz
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --tz" {
    run wtz --tz UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --tz UTC,DoesNotExist
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 1 ]
}

@test "wtz --date" {
    run wtz --date 2020-01-01
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --date 2021-01-01
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --date invalidDate
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 1 ]
}

@test "wtz --date --tz" {
    run wtz --date 2020-01-01 --tz UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --localtime" {
    ln -sf /usr/share/zoneinfo/Australia/Melbourne ${BATS_TMPDIR}/lt.$$
    run wtz --localtime ${BATS_TMPDIR}/lt.$$
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --date --localtime" {
    ln -sf /usr/share/zoneinfo/Australia/Melbourne ${BATS_TMPDIR}/lt.$$
    run wtz --localtime ${BATS_TMPDIR}/lt.$$ --date 2020-01-01
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --date --localtime --tz" {
    ln -sf /usr/share/zoneinfo/Australia/Melbourne ${BATS_TMPDIR}/lt.$$
    run wtz --localtime ${BATS_TMPDIR}/lt.$$ --date 2020-01-01 --tz UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

