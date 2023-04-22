#!/usr/bin/env bats

@test "wtz --timezones" {
    run wtz --timezones UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --timezones UTC,DoesNotExist
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 1 ]
}

@test "wtz --date --timezones" {
    run wtz --date 2020-01-01 --timezones UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --date 2021-01-01  --timezones UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]

    run wtz --date invalidDate  --timezones UTC,Australia/ACT
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 1 ]
}

@test "wtz --localtime" {
    ln -sf /usr/share/zoneinfo/Australia/Melbourne ${BATS_TMPDIR}/lt.$$
    run wtz --localtime ${BATS_TMPDIR}/lt.$$ --timezones UTC
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --date --localtime" {
    ln -sf /usr/share/zoneinfo/Australia/Melbourne ${BATS_TMPDIR}/lt.$$
    run wtz --localtime ${BATS_TMPDIR}/lt.$$ --date 2020-01-01 --timezones UTC,Australia/Melbourne
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}

@test "wtz --ignore-local-timezone" {
    run wtz --timezones UTC,Australia/Canberra --include-local-timezone=false
    printf '%s\n' 'output: ' "${output}" >&2
    [ "${status}" -eq 0 ]
}
