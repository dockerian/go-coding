#!/usr/bin/env bash

TESTLOG="$1"
VERBOSE="${VERBOSE:-${TEST_VERBOSE}}"

if [[ -f "./cover.out" ]]; then
	echo 'to see coverage report, `open cover.html` after `go tool cover`'
	echo '    go tool cover -html=cover.out -o cover.html '
	echo '    open cover.html '
	echo ''
fi

if [[ -f "${TESTLOG}" ]]; then
	COUNT_FAIL=`(cat "${TESTLOG}" | grep "\--- FAIL:" | wc -l)`
	COUNT_PASS=`(cat "${TESTLOG}" | grep "\--- PASS:" | wc -l)`
	COUNT_SKIP=`(cat "${TESTLOG}" | grep "\--- SKIP:" | wc -l)`

  if [[ "${COUNT_FAIL}${COUNT_PASS}${COUNT_SKIP}" != "000" ]]; then
		echo -e "\n==================== TEST SUMMARY ===================="

		if [[ ${COUNT_PASS} != 0 ]] || [[ "${VERBOSE}" != "" ]]; then
			printf "\n*** Passed tests  : %2d ***\n" ${COUNT_PASS}
			(cat "${TESTLOG}" | grep "\--- PASS:" | cut -d':' -f2 | sort)
		fi
		if [[ ${COUNT_SKIP} != 0 ]] || [[ "${VERBOSE}" != "" ]]; then
			printf "\n*** Skipped tests : %2d ***\n" ${COUNT_SKIP}
			(cat "${TESTLOG}" | grep "\--- SKIP:" | cut -d':' -f2 | sort)
		fi
		if [[ ${COUNT_FAIL} != 0 ]] || [[ "${VERBOSE}" != "" ]]; then
			printf "\n*** Failed tests  : %2d ***\n" ${COUNT_FAIL}
			(cat "${TESTLOG}" | grep "\--- FAIL:"| cut -d':' -f2 | sort)
		fi

		echo -e "\n=====================================================\n"
	elif [[ "${VERBOSE}" == "" ]]; then
		echo "No failed test (TEST_VERBOSE is unset)"
	fi

	# The exit code is 0 if there are no test failures.
	echo "exit code: "`cat "${TESTLOG}" | grep "\--- FAIL:" | wc -l`
	exit `cat "${TESTLOG}" | grep "\--- FAIL:" | wc -l`
fi
